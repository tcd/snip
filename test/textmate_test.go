package test

import (
	"testing"

	snip "github.com/tcd/snip/snippet"
)

func TestParseTextmateFile(t *testing.T) {
	want := snip.Snippet{
		Name:    "log",
		Trigger: "log",
		Scope:   "source.js",
		Body:    []string{"console.log(${1:$TM_SELECTED_TEXT});"},
	}

	have, err := snip.ParseTextmateFile("./testdata/textmate/js.log.tmSnippet")
	if err != nil {
		t.Errorf("error: %s", err)
	}

	err = snip.CompareSnippets(have, want)
	if err != nil {
		t.Errorf("error: %s", err)
	}
}
