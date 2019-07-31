// Package snippet provides snippet models.
package snippet

import (
	"io/ioutil"
	"regexp"
	"strings"
)

// Ultisnips returns a Snippet in Ultisnips format.
func (s Snippet) Ultisnips() string {

	var sb strings.Builder

	sb.WriteString("snippet")
	sb.WriteString(" " + s.Trigger)
	if s.Description != "" {
		sb.WriteString(" \"" + s.Description + "\"")
	}
	sb.WriteString("\n")
	sb.WriteString(strings.Join(s.Body, "\n"))
	sb.WriteString("\n")
	sb.WriteString("endsnippet")

	return sb.String()
}

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
	header := lines[0]
	body := strings.Join(lines[1:len(lines)-1], "\n")

	_ = cleanUltisnipBody(body)

	// options := regexp.MustCompile(`(?:\s+)([biemrstwA]+$)`).FindAllString(header, -1)
	// if len(options) == 1 {
	// 	if strings.Contains(options[0], "r") {
	// 		// handle delims for trigger
	// 	}
	// }

	matches := getGroups(`^snippet\s+(?P<trigger>\S+)\s(?P<description>".+")(?P<options>\s[biemrstwA]+$)?`, header)

	return Snippet{
		Trigger:     trimQuotes(matches["trigger"]),
		Body:        strings.Split(cleanUltisnipBody(body), "\n"),
		Description: trimQuotes(matches["description"]),
	}
}

// Remove interpolated code from the body of a snippet.
func cleanUltisnipBody(s string) string {
	// Python Interpolation: `\x60\!p[^\x60]+\x60` "`\!p[^`]+`"
	// Any interpolationL `\x60[^\x60]+\x60
	iPattern := regexp.MustCompile("`[^`]+`")

	interpolations := iPattern.FindAllString(s, -1)
	if len(interpolations) > 0 {
		s = iPattern.ReplaceAllString(s, "")
	}
	return s
}
