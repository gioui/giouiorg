// SPDX-License-Identifier: Unlicense OR MIT

package site

import (
	"bytes"
	"fmt"
	"html/template"
	"io/fs"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

func ParseMarkdown(includes fs.FS, path string, content []byte) (*Page, error) {
	front, content, err := parseFrontMatter(path, content)
	if err != nil {
		return nil, fmt.Errorf("parseFrontMatter: %w", err)
	}

	mdp := parser.NewWithExtensions(parser.CommonExtensions | parser.Includes | parser.Attributes | parser.Footnotes | parser.HeadingIDs | parser.AutoHeadingIDs)
	mdp.Opts.ReadIncludeFn = func(from, path string, addr []byte) []byte {
		full, err := fs.ReadFile(includes, path)
		if err != nil {
			return []byte(fmt.Sprintf("error loading include: %v", err))
		}
		content, err := extractInclude(full, string(addr))
		if err != nil {
			return []byte(fmt.Sprintf("error extracting include %q: %v", string(addr), err))
		}
		return content
	}

	page := &Page{}
	page.FrontMatter = front

	doc := markdown.Parse(content, mdp)
	for _, node := range doc.GetChildren() {
		if h, ok := node.(*ast.Heading); ok {
			page.TOC = append(page.TOC,
				InternalLink{
					Level:     h.Level,
					Title:     renderAsText(h),
					HeadingID: h.HeadingID,
				},
			)
		}
	}

	page.Content = template.HTML(markdown.Render(doc, html.NewRenderer(html.RendererOptions{Flags: html.CommonFlags})))
	return page, nil
}

func extractInclude(content []byte, addr string) ([]byte, error) {
	if addr != "" {
		rng := strings.SplitN(addr, ",", 2)
		if len(rng) != 2 {
			return nil, fmt.Errorf("invalid address: %s", addr)
		}
		start, end := rng[0], rng[1]
		startR, err := regexFor(start)
		if err != nil {
			return nil, fmt.Errorf("%s: %v", start, err)
		}
		endR, err := regexFor(end)
		if err != nil {
			return nil, fmt.Errorf("%s: %v", end, err)
		}
		startIdx := startR.FindIndex(content)
		if startIdx != nil {
			content = content[startIdx[0]:]
		}
		endIdx := endR.FindIndex(content)
		if endIdx != nil {
			content = content[:endIdx[1]]
		}
		_ = endR
	}

	// clear any leading and trailing whitespace
	content = bytes.Trim(content, "\n\r")
	content = undent(content)

	// clear any leading and trailing whitespace left-over from omitting lines
	content = bytes.Trim(content, "\n\r")
	content = append(content, '\n')

	return content, nil
}

func renderAsText(h *ast.Heading) string {
	var s string
	for _, c := range h.Children {
		if t, ok := c.(*ast.Text); ok {
			s += string(t.Literal)
		}
	}
	return s
}
