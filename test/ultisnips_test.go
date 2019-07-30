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

func TestUltisnipsSlashDelimiters(t *testing.T) {
	want := snip.Snippet{
		Name:        "",
		Trigger:     "/^main/",
		Description: "Main function",
		Rules:       "r",
		Body: []string{
			"func main() {",
			"	${0:${VISUAL}}",
			"}",
		},
	}

	have, err := snip.ParseUltisnipsFile("./testdata/vim/ultisnips/single/go.main.snippets")
	if err != nil {
		t.Errorf("error: %s", err)
	}

	err = snip.CompareSnippets(have[0], want)
	if err != nil {
		t.Errorf("error: %s", err)
	}
}

func TestUltisnipsPythonInterpolation(t *testing.T) {
	want := snip.Snippet{
		Name:        "",
		Trigger:     "",
		Description: "doc block (triple quotes)",
		Rules:       "",
		Body: []string{
			"`!p snip.rv = triple_quotes(snip)`",
			"${1:${VISUAL:doc}}",
			"`!p snip.rv = triple_quotes(snip)`",
		},
	}

	have, err := snip.ParseUltisnipsFile("./testdata/vim/ultisnips/single/python.doc.snippets")
	if err != nil {
		t.Errorf("error: %s", err)
	}

	err = snip.CompareSnippets(have[0], want)
	if err != nil {
		t.Errorf("error: %s", err)
	}
}
