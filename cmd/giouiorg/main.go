// SPDX-License-Identifier: Unlicense OR MIT

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "gioui.org/website/internal/playground"
	"golang.org/x/tools/godoc/static"
)

func main() {
	subHandler("/static/", http.HandlerFunc(staticHandler))
	subHandler("/files/", http.FileServer(http.Dir("files")))
	subHandler("/issue/", http.HandlerFunc(issueHandler))
	subHandler("/commit/", http.HandlerFunc(commitHandler))
	subHandler("/patch/", http.HandlerFunc(patchesHandler))
	http.Handle("/", vanityHandler(
		pageHandler(http.HandlerFunc(godocHandler)),
	))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	content, ok := static.Files[r.URL.Path]
	if !ok {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte(content))
}

func subHandler(root string, handler http.Handler) {
	http.Handle(root, http.StripPrefix(root, handler))
}

func patchesHandler(w http.ResponseWriter, r *http.Request) {
	url := "https://lists.sr.ht/~eliasnaur/gio/patches/" + r.URL.Path
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
	godocURL := "https://godoc.org/gioui.org" + r.URL.Path
	resp, err := http.Head(godocURL)
	switch {
	case err != nil:
		http.Error(w, "failed to HEAD godoc.org", http.StatusInternalServerError)
	case resp.StatusCode == http.StatusOK:
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
			case "/website":
				root = "gioui.org/website"
				repo = "https://git.sr.ht/~eliasnaur/giouiorg"
			default:
				root = "gioui.org"
				repo = "https://git.sr.ht/~eliasnaur/gio"
			}
			fmt.Fprintf(w, `<html><head>
<meta name="go-import" content="%[1]s git %[2]s">
<meta name="go-source" content="%[1]s _ %[2]s/tree/master{/dir} %[2]s/tree/master{/dir}/{file}#L{line}">
</head></html>`, root, repo)
			return
		}
		fallback.ServeHTTP(w, r)
	})
}
