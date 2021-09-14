// SPDX-License-Identifier: Unlicense OR MIT

package page

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"golang.org/x/tools/godoc/static"
	"gopkg.in/yaml.v2"
)

type Site struct {
	defaultTitle string
}

type frontMatter struct {
	Title    string `yaml:"title"`
	Subtitle string `yaml:"subtitle"`
}

type page struct {
	Front   frontMatter
	Content []byte

	TableOfContents []internalLink
}

type internalLink struct {
	Title     string
	HeadingID string
}

var (
	docTmpl   *template.Template
	pages     = make(map[string][]byte)
	errNoPage = errors.New("no such page")
)

type menuEntry struct {
	Title string
	Link  string
}

var (
	topMenu = []menuEntry{
		{"Home", "/"},
		{"Install", "/doc/install"},
		{"Mobile", "/doc/mobile"},
		{"Integrate", "/doc/integrate"},
		{"Architecture", "/doc/architecture"},
		{"Contribute", "/doc/contribute"},
		{"FAQ", "/doc/faq"},
	}
)

const (
	contentRoot = "content"
	includeRoot = "include"
)

func init() {
	docTmpl = template.Must(
		template.New("").Funcs(template.FuncMap{
			"IsSubPage": func(url, link string) bool {
				if url == "/index" {
					return link == "/"
				}
				if link == "/" {
					return false
				}

				return strings.HasPrefix(url, link)
			},
		}).ParseFiles(
			filepath.Join("template", "nav.tmpl"),
			filepath.Join("template", "page.tmpl"),
			filepath.Join("template", "root.tmpl"),
		))
}

func NewSite(defaultTitle string) (*Site, error) {
	s := &Site{
		defaultTitle: defaultTitle,
	}
	err := s.loadDocs(filepath.Join(contentRoot))
	return s, err
}

func ScriptsHandler(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	for _, script := range []string{"jquery.js", "playground.js"} {
		buf.WriteString(static.Files[script])
	}
	for _, script := range []string{"site.js"} {
		path := filepath.Join("files", script)
		content, err := ioutil.ReadFile(path)
		if err != nil {
			log.Printf("scriptsHandler: failed to find %q", path)
			http.Error(w, "scriptsHandler failed", http.StatusInternalServerError)
		}
		buf.Write(content)
	}
	w.Header().Set("Content-type", "application/javascript")
	w.Write(buf.Bytes())
}

func (s *Site) loadDocs(root string) error {
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || filepath.Ext(path) != ".md" {
			return nil
		}
		name := path[len(root):]
		name = name[:len(name)-len(".md")]
		content, err := s.loadMarkdown(name)
		if err != nil {
			return err
		}
		pages[name] = content
		return nil
	})
}

func (s *Site) servePage(w io.Writer, path string) error {
	var page []byte
	if os.Getenv("GAE_APPLICATION") != "" {
		p, ok := pages[path]
		if !ok {
			return errNoPage
		}
		page = p
	} else {
		p, err := s.loadMarkdown(path)
		switch {
		case os.IsNotExist(err):
			return errNoPage
		case err != nil:
			return err
		}
		page = p
	}
	_, err := w.Write(page)
	return err
}

func (s *Site) loadMarkdown(url string) ([]byte, error) {
	path := filepath.Join(contentRoot, url+".md")
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	page, err := loadPage(content)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to parse front matter: %v", path, err)
	}
	if page.Front.Title == "" {
		page.Front.Title = s.defaultTitle
	}

	mdp := parser.NewWithExtensions(parser.CommonExtensions | parser.Includes | parser.Attributes | parser.Footnotes | parser.HeadingIDs | parser.AutoHeadingIDs)
	mdp.Opts.ReadIncludeFn = func(from, path string, addr []byte) []byte {
		content, err := includeExample(path, string(addr))
		if err != nil {
			content = []byte(err.Error())
		}
		return content
	}

	doc := markdown.Parse(page.Content, mdp)
	for _, node := range doc.GetChildren() {
		if h, ok := node.(*ast.Heading); ok {
			page.TableOfContents = append(page.TableOfContents,
				internalLink{
					Title:     renderCaption(h),
					HeadingID: h.HeadingID,
				})
		}
	}

	html := markdown.Render(doc, html.NewRenderer(html.RendererOptions{Flags: html.CommonFlags}))
	args := map[string]interface{}{
		"Menu":            topMenu,
		"URL":             url,
		"Front":           page.Front,
		"Content":         template.HTML(html),
		"TableOfContents": page.TableOfContents,
	}
	var buf bytes.Buffer
	if err := docTmpl.ExecuteTemplate(&buf, "root", args); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func renderCaption(h *ast.Heading) string {
	var s string
	for _, c := range h.Children {
		if t, ok := c.(*ast.Text); ok {
			s += string(t.Literal)
		}
	}
	return s
}

func includeExample(path string, addr string) ([]byte, error) {
	path = filepath.Join(includeRoot, path)
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
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
	return content, err
}

// undent removes the number of leading tab characters in the first
// line from all lines.
func undent(text []byte) []byte {
	first := true
	ntabs := 0
	var buf bytes.Buffer
	scanner := bufio.NewScanner(bytes.NewReader(text))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasSuffix(line, "OMIT") {
			continue
		}
		if first {
			for ntabs < len(line) && line[ntabs] == '\t' {
				ntabs++
			}
			first = false
		}
		i := 0
		for i < ntabs && len(line) > 0 && line[0] == '\t' {
			i++
			line = line[1:]
		}
		buf.WriteString(line)
		buf.WriteByte('\n')
	}
	return buf.Bytes()
}

func regexFor(r string) (*regexp.Regexp, error) {
	if len(r) < 2 || r[0] != '/' || r[len(r)-1] != '/' {
		return nil, errors.New("missing / separators")
	}
	r = r[1 : len(r)-1]
	return regexp.Compile(r)
}

func loadPage(content []byte) (page, error) {
	front := frontMatter{}
	split := bytes.SplitN(content, []byte("---"), 3)
	if len(split) < 3 || len(split[0]) > 0 {
		// No front matter available.
		return page{Front: front, Content: content}, nil
	}
	header, md := split[1], split[2]
	if err := yaml.Unmarshal(header, &front); err != nil {
		return page{}, err
	}
	return page{Front: front, Content: md}, nil
}

// Handler returns a http handler that serves a page from the content
// directory, or falls back to a fallback handler if none were found.
func (s *Site) Handler(fallback http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if strings.HasSuffix(path, "/") {
			path = path + "index"
		}
		if err := s.servePage(w, path); err != nil {
			if err == errNoPage {
				fallback.ServeHTTP(w, r)
			} else {
				log.Printf("%s: %v", path, err)
				http.Error(w, "failed to serve page", http.StatusInternalServerError)
			}
		}
	})
}
