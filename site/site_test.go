// SPDX-License-Identifier: Unlicense OR MIT

package site

import (
	"bytes"
	"flag"
	"os"
	"path/filepath"
	"testing"
)

var updateRender = flag.Bool("update", false, "update render content")

func TestParse(t *testing.T) {
	config := Config{
		BaseURL:   "https://gioui.org",
		Content:   os.DirFS("./testdata/content"),
		Templates: os.DirFS("./testdata/templates"),
		Includes:  os.DirFS("./testdata/includes"),
	}

	site, err := config.Parse()
	if err != nil {
		t.Fatal(err)
	}

	for _, page := range site.Pages {
		path := filepath.FromSlash("testdata/ref/"+page.Slug) + ".html"
		rsspath := filepath.Join(filepath.FromSlash("testdata/ref/"+page.Slug), "rss.xml")

		if *updateRender {
			_ = os.MkdirAll(filepath.Dir(path), 0o600)
			if err := os.WriteFile(path, page.Rendered, 0o600); err != nil {
				t.Errorf("failed to write %q: %v", path, err)
			}
			if page.RSS {
				if err := os.WriteFile(rsspath, page.RenderedRSS, 0o600); err != nil {
					t.Errorf("failed to write %q: %v", path, err)
				}
			}
			continue
		}

		data, err := os.ReadFile(path)
		if err != nil {
			t.Errorf("failed to read %q: %v", path, err)
		}
		if !bytes.Equal(data, page.Rendered) {
			t.Errorf("different content %q", path)
		}

		if page.RSS {
			rssdata, err := os.ReadFile(rsspath)
			if err != nil {
				t.Errorf("failed to read %q: %v", path, err)
			}
			if !bytes.Equal(rssdata, page.RenderedRSS) {
				t.Errorf("different content %q", path)
			}
		}
	}
}
