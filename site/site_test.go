// SPDX-License-Identifier: Unlicense OR MIT

package site

import (
	"bytes"
	"flag"
	"os"
	"path/filepath"
	"testing"
)

var (
	updateRender = flag.Bool("update", false, "update render content")
)

func TestParse(t *testing.T) {
	config := Config{
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
		if *updateRender {
			os.MkdirAll(filepath.Dir(path), 0600)
			if err := os.WriteFile(path, page.Rendered, 0600); err != nil {
				t.Errorf("failed to write %q: %v", path, err)
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
	}
}
