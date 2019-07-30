package snippet

import (
	"io/ioutil"
	"strings"

	"howett.net/plist"
)

type tmSnippet struct {
	Name       string `plist:"name"`
	TabTrigger string `plist:"tabTrigger"`
	Content    string `plist:"content"`
	Scope      string `plist:"scope"`
	UUID       string `plist:"uuid"`
}

var tmVariables = []string{
	"$TM_FILEPATH",
	"$TM_DIRECTORY",
	"$TM_CURRENT_LINE",
	"$TM_CURRENT_WORD",
	"$TM_LINE_NUMBER",
	"$TM_SELECTED_TEXT",
	"$TM_SOFT_TABS",
	"$TM_TAB_SIZE",
}

// ParseTextmateFile returns a tmSnippet object with the data
// parsed from a TextMate snippet (.plist) file.
func ParseTextmateFile(path string) (Snippet, error) {
	var s Snippet
	var ts tmSnippet

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return s, err
	}

	_, err = plist.Unmarshal(bytes, &ts)
	if err != nil {
		return s, err
	}

	return textMateSnippetToSnippet(ts), nil
}

func textMateSnippetToSnippet(s tmSnippet) Snippet {
	var snip Snippet

	snip.Name = strings.TrimSpace(s.Name)
	snip.Trigger = strings.TrimSpace(s.TabTrigger)
	snip.Scope = strings.TrimSpace(s.Scope)
	snip.Body = strings.Split(s.Content, "\n")

	return snip
}
