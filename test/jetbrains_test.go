package test

import (
	"testing"

	snip "github.com/tcd/snip/snippet"
)

func TestJetBrainsFile(t *testing.T) {
	want := []snip.Snippet{
		snip.Snippet{
			Name:        "cll",
			Trigger:     "cll",
			Body:        []string{"console.log($END$);"},
			Description: "Console: console.log (text only)",
		},
	}

	have, err := snip.ParseJetBrainsFile("./testdata/jetbrains/multi/test1.xml")
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
