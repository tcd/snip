package test

import (
	"testing"

	snip "github.com/tcd/snip/snippet"
)

func TestParseUltisnipsFile(t *testing.T) {
	want := snip.Snippet{
		Name:        "",
		Trigger:     "lorem",
		Description: "Lorem Ipsum - 50 Words",
		Rules:       "b",
		Body: []string{
			"Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod",
			"tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At",
			"vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren,",
			"no sea takimata sanctus est Lorem ipsum dolor sit amet.",
		},
		Rules:       "b",
		Description: "Lorem Ipsum - 50 Words",
	}

	have, err := snip.ParseUltisnipsFile("./testdata/vim/ultisnips/single/all.lorem.snippets")
	if err != nil {
		t.Errorf("error: %s", err)
	}

	err = snip.CompareSnippets(have[0], want)
	if err != nil {
		t.Errorf("error: %s", err)
	}
}
