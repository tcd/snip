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
			Description: "A multiline snippet",
		},
		snip.Snippet{
			Name:        "single line snippet",
			Trigger:     "single",
			Body:        []string{"single line, but it's pretty long"},
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

// func TestTopLevelScope(t *testing.T) {
// 	want := []snip.Snippet{
// 		snip.Snippet{
// 			Name:        "single import",
// 			Trigger:     "im",
// 			Body:        []string{`import \"${1:package}\"`},
// 			Description: "Snippet for import statement",
// 		},
// 		snip.Snippet{
// 			Name:        "multiple imports",
// 			Trigger:     "ims",
// 			Body:        []string{`import (\n\t\"${1:package}\"\n)`},
// 			Description: "Snippet for a import block",
// 		},
// 		snip.Snippet{
// 			Name:        "single constant",
// 			Trigger:     "co",
// 			Body:        []string{"const ${1:name} = ${2:value}"},
// 			Description: "Snippet for constant",
// 		},
// 		snip.Snippet{
// 			Name:        "multiple constants",
// 			Trigger:     "cos",
// 			Body:        []string{"const (\n\t${1:name} = ${2:value}\n)"},
// 			Description: "Snippet for a constant block",
// 		},
// 	}

// 	have, err := snip.ParseCodeFile("./testdata/vscode/test2.json")
// 	if err != nil {
// 		t.Errorf("error: %s", err)
// 	}

// 	for i := range have {
// 		got, err := snip.FindSnippetByName(have[i].Name, want)
// 		if err != nil {
// 			t.Errorf("error: %s", err)
// 			continue
// 		}
// 		err = snip.CompareSnippets(have[i], got)
// 		if err != nil {
// 			t.Errorf("error: %s", err)
// 		}
// 	}
// }
