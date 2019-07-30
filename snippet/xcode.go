package snippet

import (
	"io/ioutil"
	"strings"

	"howett.net/plist"
)

type xcodeSnippet struct {
	Trigger          string   `plist:"IDECodeSnippetCompletionPrefix"`
	Title            string   `plist:"IDECodeSnippetTitle"`
	Scope            string   `plist:"IDECodeSnippetLanguage"`
	Summary          string   `plist:"IDECodeSnippetSummary"`
	Contents         string   `plist:"IDECodeSnippetContents"`
	CompletionScopes []string `plist:"IDECodeSnippetCompletionScopes"`
	UUID             string   `plist:"IDECodeSnippetIdentifier"`
	UserSnippet      bool     `plist:"IDECodeSnippetUserSnippet"`
	SnippetVersion   int      `plist:"IDECodeSnippetVersion"`
}

// ParseXcodeFile Parses a plist file containing an Xcode snippet.
func ParseXcodeFile(path string) (Snippet, error) {
	var s Snippet
	var xs xcodeSnippet

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return s, err
	}

	_, err = plist.Unmarshal(bytes, &xs)
	if err != nil {
		return s, err
	}

	return xcodeSnippetToSnippet(xs), nil
}

func xcodeSnippetToSnippet(s xcodeSnippet) Snippet {
	return Snippet{
		Trigger:     strings.TrimSpace(s.Trigger),
		Body:        strings.Split(s.Contents, "\n"),
		Description: strings.TrimSpace(s.Title),
	}
}
