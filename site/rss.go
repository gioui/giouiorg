// SPDX-License-Identifier: Unlicense OR MIT

package site

import (
	"encoding/xml"
	"fmt"
	"net/url"
)

// RenderRSS creates an rss.xml from the specified items.
func (site *Site) renderRSS(page PageData) ([]byte, error) {
	channelLink, err := url.JoinPath(site.BaseURL, page.URL())
	if err != nil {
		return nil, fmt.Errorf("unable to create channel link %q %q: %w", site.BaseURL, page.URL(), err)
	}

	channel := RSSChannel{
		Title:       page.Title,
		Link:        channelLink,
		Description: page.Subtitle,
	}

	for _, child := range page.Children {
		itemLink, err := url.JoinPath(site.BaseURL, child.URL())
		if err != nil {
			return nil, fmt.Errorf("unable to create item link %q %q: %w", site.BaseURL, child.URL(), err)
		}

		item := RSSItem{
			Title:       child.Title,
			Link:        itemLink,
			Description: child.Summary,
		}

		if child.Date != nil {
			item.PubDate = child.Date.Format("Mon, 02 Jan 2006")
		}

		channel.Items = append(channel.Items, item)
	}

	data, err := xml.MarshalIndent(RSSXML{
		Version:  "2.0",
		Channels: []RSSChannel{channel},
	}, "", "    ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal rss: %w", err)
	}

	return append([]byte(xml.Header), data...), nil
}

type RSSXML struct {
	XMLName  xml.Name     `xml:"rss"`
	Version  string       `xml:"version,attr"`
	Channels []RSSChannel `xml:"channel"`
}

type RSSChannel struct {
	XMLName     xml.Name  `xml:"channel"`
	Title       string    `xml:"title"`
	Link        string    `xml:"link"`
	Description string    `xml:"description"`
	Items       []RSSItem `xml:"item"`
}

type RSSItem struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	Description string   `xml:"description"`
	PubDate     string   `xml:"pubDate,omitempty"`
}
