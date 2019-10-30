// SPDX-License-Identifier: Unlicense OR MIT

package playground

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"
)

const baseURL = "https://play.golang.org"

func init() {
	for _, cmd := range []string{"/compile", "/share"} {
		http.HandleFunc(cmd, proxy)
	}
}

func proxy(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	proxyReq, err := http.NewRequest("POST", "https://play.golang.org"+r.URL.Path, r.Body)
	if err != nil {
		http.Error(w, "http.NewRequest failed", http.StatusInternalServerError)
	}
	ctx, cancel := context.WithTimeout(r.Context(), 60*time.Second)
	defer cancel()
	proxyReq = proxyReq.WithContext(ctx)
	proxyReq.Header.Set("User-Agent", "Gio_Playground")
	proxyReq.Header.Set("Content-Type", r.Header.Get("Content-Type"))
	resp, err := http.DefaultClient.Do(proxyReq)
	if err != nil {
		log.Printf("playground: %v", err)
		http.Error(w, "failed to proxy playground request", http.StatusInternalServerError)
	}
	defer resp.Body.Close()
	io.Copy(w, resp.Body)
}
