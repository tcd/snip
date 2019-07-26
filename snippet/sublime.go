package snippet

import (
	"encoding/xml"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type sublimeSnippet struct {
	Body        string `xml:"content"`
	Trigger     string `xml:"tabTrigger"`
	Description string `xml:"description"`
	Scope       string `xml:"scope"`
}

// ParseSublimeFile parses a Sublime Text ".sublime-snippet" file
// and returns a snip.Snippet or an error if something goes wrong.
func ParseSublimeFile(path string) (Snippet, error) {
	var s sublimeSnippet

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return Snippet{}, err
	}

	err = xml.Unmarshal(bytes, &s)
	if err != nil {
		return Snippet{}, err
	}

	return sublimeSnippetToSnippet(s, path), nil
}

func sublimeSnippetToSnippet(s sublimeSnippet, path string) Snippet {
	name := filepath.Base(path)
	// Not using `strings.TrimSuffix` here because we
	// don't want to remove any extra `.`'s in the filename.
	name = strings.Replace(name, ".sublime-snippet", "", -1)

	return Snippet{
		Name:        strings.TrimSpace(name),
		Trigger:     strings.TrimSpace(s.Trigger),
		Rules:       "",
		Body:        strings.Split(s.Body, "\n"),
		Description: strings.TrimSpace(s.Description),
	}
}
