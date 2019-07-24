// Package snippet provides snippet models.
package snippet

import (
	"regexp"
	"strings"
)

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

// ParseUltisnipsFile returns an array of strings containing all snippets in a file at the given path.
func ParseUltisnipsFile(path string) []string {
	bytes := bytesFromFile(path)
	re := regexp.MustCompile(`(?msU:((^snippet.+$\n)(.+)(?:endsnippet)))`)
	matches := re.FindAllString(string(bytes), -1)
	return matches
}
