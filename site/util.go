// SPDX-License-Identifier: Unlicense OR MIT

package site

import (
	"bufio"
	"bytes"
	"errors"
	"regexp"
	"strings"
)

func regexFor(r string) (*regexp.Regexp, error) {
	if len(r) < 2 || r[0] != '/' || r[len(r)-1] != '/' {
		return nil, errors.New("missing / separators")
	}
	r = r[1 : len(r)-1]
	return regexp.Compile(r)
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
