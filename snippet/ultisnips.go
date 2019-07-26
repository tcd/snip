// Package snippet provides snippet models.
package snippet

import (
	"io/ioutil"
	"regexp"
	"strings"
)

// ParseUltisnipsFile returns an array of strings containing all snippets in a file at the given path.
func ParseUltisnipsFile(path string) ([]Snippet, error) {
	var snippets []Snippet

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return snippets, err
	}
	re := regexp.MustCompile(`(?msU:((^snippet.+$\n)(.+)(?:endsnippet)))`)
	matches := re.FindAllString(string(bytes), -1)

	for _, s := range matches {
		snippets = append(snippets, ParseUltisnipsSnippet(s))
	}
	return snippets, nil
}

// ParseUltisnipsSnippet parses a raw string UltiSnips snippet.
func ParseUltisnipsSnippet(s string) Snippet {

	lines := strings.Split(s, "\n")

	matches := getGroups(`^snippet\s+(?P<trigger>.?\w+.?)\s(?P<description>".+")\s(?P<options>[biemrstwA]+)?$`, lines[0])
	body := lines[:len(lines)-2]

	return Snippet{
		Name:        "",
		Trigger:     matches["trigger"],
		Rules:       matches["rules"],
		Body:        body,
		Description: matches["description"],
	}
}
