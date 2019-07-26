package test

import (
	"testing"

	snip "github.com/tcd/snip/snippet"
)

func TestSublimeSnippet(t *testing.T) {
	want := snip.Snippet{
		Name:        "lorem",
		Trigger:     "lorem",
		Description: "HTML - Lorem Ipsum",
		Rules:       "",
		Body:        []string{"Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."},
	}

	have, err := snip.ParseSublimeFile("./testdata/sublime/lorem.sublime-snippet")
	if err != nil {
		t.Errorf("error: %s", err)
	}

	err = snip.CompareSnippets(have, want)
	if err != nil {
		t.Errorf("error: %s", err)
	}
}
