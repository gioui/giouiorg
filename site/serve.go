// SPDX-License-Identifier: Unlicense OR MIT

package site

import (
	"bytes"
	"io/fs"
	"log"
	"net/http"
	"strings"

	"golang.org/x/tools/godoc/static"
)

func (site *Site) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/files/") {
		http.StripPrefix("/files/",
			http.FileServer(http.FS(site.Files)),
		).ServeHTTP(w, r)
		return
	}

	if r.URL.Path == "/scripts.js" {
		site.handleScripts(w, r)
		return
	}

	slug := strings.TrimPrefix(r.URL.Path, "/")
	page, ok := site.Pages[slug]
	if !ok {
		site.Fallback(w, r)
		return
	}
	if page.Slug != slug {
		http.Redirect(w, r, "/"+page.Slug, http.StatusPermanentRedirect)
		return
	}

	_, _ = w.Write(page.Rendered)
}

func (site *Site) handleScripts(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	for _, script := range []string{"jquery.js", "playground.js"} {
		buf.WriteString(static.Files[script])
	}
	for _, script := range []string{"site.js"} {
		content, err := fs.ReadFile(site.Files, script)
		if err != nil {
			log.Printf("scriptsHandler: failed to find %q", script)
			http.Error(w, "scriptsHandler failed", http.StatusInternalServerError)
		}
		buf.Write(content)
	}
	w.Header().Set("Content-type", "application/javascript")
	w.Write(buf.Bytes())
}
