package test

import (
	"testing"

	snip "github.com/tcd/snip/snippet"
)

func TestParseCodeFile(t *testing.T) {
	want := []snip.Snippet{
		snip.Snippet{
			Name:        "multiline snippet",
			Trigger:     "multi",
			Body:        []string{"line1", "line2", "line3"},
			Rules:       "",
			Description: "A multiline snippet",
		},
		snip.Snippet{
			Name:        "single line snippet",
			Trigger:     "single",
			Body:        []string{"single line, but it's pretty long"},
			Rules:       "",
			Description: "A single line snippet",
		},
	}

	have, err := snip.ParseCodeFile("./testdata/vscode/test1.json")
	if err != nil {
		t.Errorf("error: %s", err)
	}

	for i := range have {
		got, err := snip.FindSnippetByName(have[i].Name, want)
		if err != nil {
			t.Errorf("error: %s", err)
			continue
		}
		err = snip.CompareSnippets(have[i], got)
		if err != nil {
			t.Errorf("error: %s", err)
		}
	}
}
