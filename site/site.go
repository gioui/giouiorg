// SPDX-License-Identifier: Unlicense OR MIT

package site

import (
	"bytes"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"path/filepath"
)

type Config struct {
	Content   fs.FS
	Templates fs.FS
	Includes  fs.FS
	Files     fs.FS
	Fallback  http.HandlerFunc
}

type Site struct {
	// Pages are indexed without prefixed "/".
	Pages        map[string]*Page
	Template     *template.Template
	BaseTemplate *template.Template
	Content      fs.FS
	Files        fs.FS
	Fallback     http.HandlerFunc
}

type Page struct {
	FrontMatter
	TOC []InternalLink

	Template string

	Content  template.HTML
	Rendered []byte

	Parent   *Page
	Before   *Page
	After    *Page
	Children []*Page
}

type InternalLink struct {
	Level     int
	Title     string
	HeadingID string
}

func (config Config) Parse() (*Site, error) {
	site := &Site{
		Pages:    map[string]*Page{},
		Content:  config.Content,
		Files:    config.Files,
		Fallback: config.Fallback,
	}

	if err := site.loadTemplates(config.Templates); err != nil {
		return nil, fmt.Errorf("loadTemplates: %w", err)
	}
	if err := site.loadContent(config.Content, config.Includes); err != nil {
		return nil, fmt.Errorf("loadContent: %w", err)
	}
	if err := site.linkPages(); err != nil {
		return nil, fmt.Errorf("linkPages: %w", err)
	}
	if err := site.renderPages(); err != nil {
		return nil, fmt.Errorf("renderPages: %w", err)
	}

	return site, nil
}

// loadTemplates loads all the templates from templates.
func (site *Site) loadTemplates(templates fs.FS) error {
	templ, err := template.New("site.tmpl").ParseFS(templates, "*.tmpl")
	if err != nil {
		return fmt.Errorf("parse failed: %w", err)
	}

	site.Template = templ
	site.BaseTemplate, _ = templ.Clone()

	return nil
}

// loadContent loads all the content from root and includes.
func (site *Site) loadContent(contentRoot, includes fs.FS) error {
	return fs.WalkDir(contentRoot, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}

		content, err := fs.ReadFile(contentRoot, path)
		if err != nil {
			return fmt.Errorf("failed to load page %q: %w", path, err)
		}

		ext := filepath.Ext(path)
		if imageExt[ext] {
			// ignore images
			return nil
		}

		switch ext {
		case ".tmpl":
			return site.loadTemplate(includes, path, content)
		case ".md":
			return site.loadMarkdown(includes, path, content)
		default:
			return fmt.Errorf("unknown page extension %q", path)
		}
	})
}

// loadTemplate loads a single go template file into the index.
func (site *Site) loadTemplate(includes fs.FS, path string, content []byte) error {
	page, err := ParseTemplate(includes, path, content)
	if err != nil {
		return fmt.Errorf("ParseTemplate: %w", err)
	}

	if _, exists := site.Pages[page.Slug]; exists {
		return fmt.Errorf("duplicate page %q", page.Slug)
	}

	site.Pages[page.Slug] = page

	return nil
}

// loadMarkdown loads a single markdown file into the index.
func (site *Site) loadMarkdown(includes fs.FS, path string, content []byte) error {
	page, err := ParseMarkdown(includes, path, content)
	if err != nil {
		return fmt.Errorf("ParseMarkdown: %w", err)
	}

	if _, exists := site.Pages[page.Slug]; exists {
		return fmt.Errorf("duplicate page %q", page.Slug)
	}

	site.Pages[page.Slug] = page

	return nil
}

// linkPages adds links between pages with children and children themselves.
func (site *Site) linkPages() error {
	for _, page := range site.Pages {
		if len(page.FrontMatter.Children) == 0 {
			continue
		}

		for _, slug := range page.FrontMatter.Children {
			child, ok := site.Pages[slug]
			if !ok {
				return fmt.Errorf("unable to find page %q", slug)
			}

			page.Children = append(page.Children, child)

			if child.Parent != nil {
				return fmt.Errorf("child %q already has parent %q; wanted to add %q", slug, child.Parent.Slug, page.Slug)
			}
		}

		if page.FrontMatter.After != "" {
			page.After = site.Pages[page.FrontMatter.After]
			if page.After == nil {
				return fmt.Errorf("after %q does not exist", page.FrontMatter.After)
			}
		}
		if page.FrontMatter.Before != "" {
			page.Before = site.Pages[page.FrontMatter.Before]
			if page.Before == nil {
				return fmt.Errorf("before %q does not exist", page.FrontMatter.Before)
			}
		}

		if page.FrontMatter.ChildrenNoLink {
			continue
		}

		for i, child := range page.Children {
			child.Parent = page
			if i > 0 {
				child.Before = page.Children[i-1]
			}
			if i < len(page.Children)-1 {
				child.After = page.Children[i+1]
			}
		}
	}

	return nil
}

// renderPages renders the final html for each page.
func (site *Site) renderPages() error {
	root := site.Pages[""]

	for _, page := range site.Pages {
		type renderData struct {
			Nav   Nav
			Front *FrontMatter
			*Page
		}

		nav := buildNav(root, page)
		nav.Active = root == page // override index.md active

		data := renderData{
			Nav:   nav,
			Front: &page.FrontMatter,
			Page:  page,
		}

		var buf bytes.Buffer
		if page.Template != "" {
			siteTempl, err := site.BaseTemplate.Clone()
			if err != nil {
				return fmt.Errorf("template clone failed: %w", err)
			}
			t, err := siteTempl.Parse(string(page.Template))
			if err != nil {
				return fmt.Errorf("template parse failed: %w", err)
			}
			if err := t.ExecuteTemplate(&buf, "root", data); err != nil {
				return fmt.Errorf("template failed: %w", err)
			}
		} else {
			if err := site.Template.ExecuteTemplate(&buf, "root", data); err != nil {
				return fmt.Errorf("template failed: %w", err)
			}
		}
		page.Rendered = buf.Bytes()
	}
	return nil
}

type Nav struct {
	Page     *Page
	Active   bool
	Children []Nav
}

func buildNav(page *Page, active *Page) Nav {
	nav := Nav{
		Page:   page,
		Active: active.hasParent(page) || active == page,
	}
	for _, child := range page.Children {
		nav.Children = append(nav.Children, buildNav(child, active))
	}
	return nav
}

func (page *Page) hasParent(target *Page) bool {
	for parent := page.Parent; parent != nil; parent = parent.Parent {
		if parent == target {
			return true
		}
	}
	return false
}
