package snippet

import (
	"io/ioutil"
	"regexp"
	"strings"
)

type neosnippetSnippet struct {
	Snippet string   // name
	Abbr    string   // abbreviation
	Alias   string   // aliases
	Regexp  string   // pattern
	Options string   // options
	Body    []string // body
}

// ParseNeosnippetFile returns an array of Snippets from file at the given path.
func ParseNeosnippetFile(path string) ([]Snippet, error) {
	var snippets []Snippet

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return snippets, err
	}
	re := regexp.MustCompile(`(?msU:(^snippet(.+)(^$)))`)
	matches := re.FindAllString(string(bytes), -1)

	for _, s := range matches {
		snippets = append(snippets, parseNeosnippet(s))
	}

	return snippets, nil
}

func parseNeosnippet(s string) Snippet {
	var snip Snippet

	lines := strings.Split(s, "\n")

	for _, line := range lines {
		if len(line) >= 5 && line[:5] == "abbr " {
			snip.Trigger = strings.TrimSpace(line[4:])
		}
		if len(line) >= 8 && line[:8] == "snippet " {
			snip.Name = strings.TrimSpace(line[8:])
			if snip.Trigger == "" {
				snip.Trigger = snip.Name
			}
		}
	}

	re := regexp.MustCompile(`(?ms:(^\s+.+))`)
	matches := re.FindAllString(s, -1)
	if len(matches) > 0 {
		bodyLines := strings.Split(matches[0], "\n")
		for _, line := range bodyLines {
			snip.Body = append(snip.Body, strings.TrimSpace(line))
		}
	}

	return snip
}
