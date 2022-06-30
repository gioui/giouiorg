// SPDX-License-Identifier: Unlicense OR MIT

package site

import (
	"bytes"
	"fmt"
	"path"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

// FrontMatter is used to define properties for the page.
type FrontMatter struct {
	Slug       string `yaml:"url"`
	Title      string `yaml:"title"`
	Subtitle   string `yaml:"subtitle"`
	ShortTitle string `yaml:"shorttitle"`

	After  string `yaml:"after"`
	Before string `yaml:"before"`

	HideChildren   bool     `yaml:"hidechildren"`
	ChildrenNoLink bool     `yaml:"childrennolink"`
	Children       []string `yaml:"children"`

	Images []Image        `yaml:"images"`
	Links  []ExternalLink `yaml:"links"`
}

type Image struct {
	Alt    string `yaml:"alt"`
	Source string `yaml:"source"`
}

type ExternalLink struct {
	Title string `yaml:"title"`
	URL   string `yaml:"url"`
}

func (fm FrontMatter) URL() string { return "/" + fm.Slug }

// parseFrontMatter parses a "---" delimited header from a file.
func parseFrontMatter(fpath string, content []byte) (FrontMatter, []byte, error) {
	front := FrontMatter{
		Slug: fpath[:len(fpath)-len(filepath.Ext(fpath))],
	}

	if path.Base(front.Slug) == "index" {
		front.Slug = path.Dir(fpath)
		if front.Slug == "." {
			front.Slug = ""
		}
	}

	split := bytes.SplitN(content, []byte("---"), 3)
	if len(split) < 3 || len(split[0]) > 0 {
		return front, content, fmt.Errorf("front matter missing")
	}

	header, md := split[1], split[2]
	if err := yaml.Unmarshal(header, &front); err != nil {
		return front, content, err
	}

	resolveRelativePath(&front.After, fpath)
	resolveRelativePath(&front.Before, fpath)
	for i := range front.Children {
		resolveRelativePath(&front.Children[i], fpath)
	}

	for i := range front.Images {
		resolveRelativeImageURL(&front.Images[i].Source, fpath)
	}

	return front, md, nil
}

func resolveRelativePath(target *string, workingPath string) {
	if !strings.HasPrefix(*target, "./") {
		return
	}

	*target = path.Clean(path.Join(path.Dir(workingPath), *target))
}

func resolveRelativeImageURL(target *string, workingPath string) {
	if !strings.HasPrefix(*target, "./") {
		return
	}

	// TODO: avoid difference between paths with and without front slash
	// use / prefixed slugs and paths throughout the codebase.
	*target = "/" + path.Clean(path.Join(path.Dir(workingPath), *target))
}
