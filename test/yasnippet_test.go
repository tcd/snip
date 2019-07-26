package test

import (
	"testing"

	snip "github.com/tcd/snip/snippet"
)

func TestYasnippetSnippet(t *testing.T) {
	want := snip.Snippet{
		Name:        "console.log",
		Trigger:     "clg",
		Body:        []string{"", "console.log(${1:object})", ""},
		Rules:       "",
		Description: "",
	}

	have, err := snip.ParseYasnippetFile("./testdata/yasnippet/js/clg")
	if err != nil {
		t.Errorf("error: %s", err)
	}

	err = snip.CompareSnippets(have, want)
	if err != nil {
		t.Errorf("error: %s", err)
	}
}
