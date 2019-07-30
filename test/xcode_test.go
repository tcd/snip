package test

import (
	"testing"

	snip "github.com/tcd/snip/snippet"
)

func TestParseXcodeFile(t *testing.T) {
	want := snip.Snippet{
		Trigger:     "swift-forin",
		Description: "Swift for-in loop with casting",
		Body: []string{
			"for case let <#object#> in <#collection#> as [<#Type#>] {",
			"",
			"}",
		},
	}

	have, err := snip.ParseXcodeFile("./testdata/xcode/swift-forin.codesnippet")
	if err != nil {
		t.Errorf("error: %s", err)
	}

	err = snip.CompareSnippets(have, want)
	if err != nil {
		t.Errorf("error: %s", err)
	}
}
