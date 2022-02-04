// SPDX-License-Identifier: Unlicense OR MIT

package site

import (
	"fmt"
	"html/template"
	"io/fs"
)

func ParseTemplate(includes fs.FS, path string, content []byte) (*Page, error) {
	front, content, err := parseFrontMatter(path, content)
	if err != nil {
		return nil, fmt.Errorf("parseFrontMatter: %w", err)
	}

	// early test for parsing the template
	_, err = template.New("").Parse(string(content))
	if err != nil {
		return nil, fmt.Errorf("parse template: %w", err)
	}

	page := &Page{}
	page.FrontMatter = front
	page.Template = string(content)

	return page, nil
}
