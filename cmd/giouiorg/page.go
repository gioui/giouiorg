// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gomarkdown/markdown"
	"gopkg.in/yaml.v2"
)

type frontMatter struct {
	Title string `yaml:"title"`
}

type page struct {
	Front   frontMatter
	Content []byte
}

var (
	docTmpl *template.Template
	pages   = make(map[string][]byte)
)

func init() {
	docTmpl = template.Must(template.ParseFiles(
		filepath.Join("template", "page.tmpl"),
		filepath.Join("template", "root.tmpl"),
	))
	if err := loadDocs(filepath.Join("content")); err != nil {
		log.Fatal(err)
	}
}

func loadDocs(root string) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || filepath.Ext(path) != ".md" {
			return nil
		}
		name := path[len(root):]
		name = name[:len(name)-len(".md")]
		content, err := loadMarkdown(path)
		if err != nil {
			return err
		}
		pages[name] = content
		return nil
	})
}

func loadMarkdown(f string) ([]byte, error) {
	content, err := ioutil.ReadFile(f)
	if err != nil {
		return nil, err
	}
	page, err := loadPage(content)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to parse front matter: %v", f, err)
	}
	if page.Front.Title == "" {
		page.Front.Title = "Gio - immediate mode GUI in Go"
	}
	html := markdown.ToHTML(page.Content, nil, nil)
	args := struct {
		Front   frontMatter
		Content template.HTML
	}{page.Front, template.HTML(html)}
	var buf bytes.Buffer
	if err := docTmpl.ExecuteTemplate(&buf, "root", args); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func loadPage(content []byte) (page, error) {
	front := frontMatter{}
	split := bytes.SplitN(content, []byte("---"), 3)
	if len(split) < 3 || len(split[0]) > 0 {
		// No front matter available.
		return page{front, content}, nil
	}
	header, md := split[1], split[2]
	if err := yaml.Unmarshal(header, &front); err != nil {
		return page{}, err
	}
	return page{front, md}, nil
}

// pageHandler serves a page from the content directory, or
// falls back to a fallback handler if none were found.
func pageHandler(fallback http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if strings.HasSuffix(path, "/") {
			path = path + "index"
		}
		p, ok := pages[path]
		if !ok {
			fallback.ServeHTTP(w, r)
			return
		}
		w.Write(p)
	})
}
