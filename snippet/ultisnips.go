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
	if s.Rules != "" {
		sb.WriteString(" " + s.Rules)
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
	// body := strings.Join(lines[1:len(lines)-1], "\n")

	// options := regexp.MustCompile(`(?:\s+)([biemrstwA]+$)`).FindAllString(header, -1)
	// if len(options) == 1 {
	// 	if strings.Contains(options[0], "r") {
	// 		// handle delims for trigger
	// 	}
	// }

	// Python Interpolation: `\x60\!p[^\x60]+\x60`
	// Any interpolationL `\x60[^\x60]+\x60
	// interpolations := regexp.MustCompile("`[^`]+`").FindAllString(body, -1)
	// if len(interpolations) > 0 {
	// 	if strings.Contains(interpolations[0], "r") {
	// 		// handle delims for trigger
	// 		fmt.Println(interpolations)
	// 	}
	// }

	matches := getGroups(`^snippet\s+(?P<trigger>.?\w+.?)\s(?P<description>".+")\s(?P<options>[biemrstwA]+$)?`, header)

	return Snippet{
		Name:        "",
		Trigger:     trimQuotes(matches["trigger"]),
		Rules:       matches["options"],
		Body:        lines[1 : len(lines)-1],
		Description: trimQuotes(matches["description"]),
	}
}
