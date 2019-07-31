package test

import (
	"testing"

	snip "github.com/tcd/snip/snippet"
)

func TestParseNeosnippetFile(t *testing.T) {
	want := []snip.Snippet{
		snip.Snippet{
			Name:    "set",
			Trigger: "set",
			Body:    []string{"set ${1:#:NAME}(${2:#:ARGS}) {", "${0:TARGET}", "}"},
		},
	}

	have, err := snip.ParseNeosnippetFile("./testdata/vim/neosnippet/multi/javascript.snip")
	if err != nil {
		t.Errorf("error: %s", err)
	}

	for i := range want {
		got, err := snip.FindSnippetByName(want[i].Name, have)
		if err != nil {
			t.Errorf("error: %s", err)
			continue
		}
		err = snip.CompareSnippets(want[i], got)
		if err != nil {
			t.Errorf("error: %s", err)
		}
	}
}
