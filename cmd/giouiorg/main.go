// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "gioui.org/website/internal/playground"
	"gioui.org/website/page"
)

func main() {
	subHandler("/scripts.js", http.HandlerFunc(page.ScriptsHandler))
	subHandler("/files/", http.FileServer(http.Dir("files")))
	subHandler("/issue/", http.HandlerFunc(issueHandler))
	subHandler("/commit/", http.HandlerFunc(commitHandler))
	subHandler("/patch/", http.HandlerFunc(patchesHandler))
	subHandler("/include/files/", http.FileServer(http.Dir("include/files")))

	site, err := page.NewSite("Gio - immediate mode GUI in Go")
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", vanityHandler(
		site.Handler(http.HandlerFunc(godocHandler)),
	))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func subHandler(root string, handler http.Handler) {
	http.Handle(root, http.StripPrefix(root, handler))
}

func patchesHandler(w http.ResponseWriter, r *http.Request) {
	url := "https://lists.sr.ht/~eliasnaur/gio-patches/patches/" + r.URL.Path
	http.Redirect(w, r, url, http.StatusFound)
}

func issueHandler(w http.ResponseWriter, r *http.Request) {
	url := "https://todo.sr.ht/~eliasnaur/gio/" + r.URL.Path
	http.Redirect(w, r, url, http.StatusFound)
}

func commitHandler(w http.ResponseWriter, r *http.Request) {
	commit := r.URL.Path
	var url string
	if commit == "/" {
		url = "https://git.sr.ht/~eliasnaur/gio/log"
	} else {
		url = "https://git.sr.ht/~eliasnaur/gio/commit/" + commit
	}
	http.Redirect(w, r, url, http.StatusFound)
}

func godocHandler(w http.ResponseWriter, r *http.Request) {
	godocURL := "https://pkg.go.dev/gioui.org" + r.URL.Path
	resp, err := http.Head(godocURL)
	switch {
	case err != nil:
		http.Error(w, "failed to HEAD godoc.org", http.StatusInternalServerError)
	case resp.StatusCode == http.StatusOK:
		http.Redirect(w, r, godocURL, http.StatusFound)
	case resp.StatusCode == http.StatusMethodNotAllowed:
		// Because of https://github.com/golang/go/issues/43739, we can't HEAD
		// the pkg.go.dev site. Redirect blindly.
		http.Redirect(w, r, godocURL, http.StatusFound)
	default:
		http.NotFound(w, r)
	}
}

// vanityHandler serves git location meta headers for the go tool.
// If the go-get=1 query is not present it falls back to handler.
func vanityHandler(fallback http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("go-get") == "1" {
			var repo, root string
			switch r.URL.Path {
			case "/example":
				root = "gioui.org/example"
				repo = "https://git.sr.ht/~eliasnaur/gio-example"
			case "/website":
				root = "gioui.org/website"
				repo = "https://git.sr.ht/~eliasnaur/giouiorg"
			case "/x":
				root = "gioui.org/x"
				repo = "https://git.sr.ht/~whereswaldon/gio-x"
			default:
				root = "gioui.org"
				repo = "https://git.sr.ht/~eliasnaur/gio"
			}
			fmt.Fprintf(w, `<html><head>
<meta name="go-import" content="%[1]s git %[2]s">
<meta name="go-source" content="%[1]s _ %[2]s/tree/main{/dir} %[2]s/tree/main{/dir}/{file}#L{line}">
</head></html>`, root, repo)
			return
		}
		fallback.ServeHTTP(w, r)
	})
}
