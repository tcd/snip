package snippet

import (
	"errors"
	"regexp"
)

// getGroups parses testString with the given regular expression
// and returns the group values defined in the expression.
//
// Thanks eluleci (https://stackoverflow.com/a/39635221/7687024)
//
func getGroups(regEx, testString string) (groupsMap map[string]string) {

	compiledRegEx := regexp.MustCompile(regEx)

	match := compiledRegEx.FindStringSubmatch(testString)

	groupsMap = make(map[string]string)

	for i, name := range compiledRegEx.SubexpNames() {
		if i > 0 && i <= len(match) {
			groupsMap[name] = match[i]
		}
	}

	return groupsMap
}

// trimQuotes Removes surrounding quotes from a string if they are present.
func trimQuotes(s string) string {
	if len(s) >= 2 {
		if s[0] == '"' && s[len(s)-1] == '"' {
			return s[1 : len(s)-1]
		}
	}
	return s
}

// CompareSnippets Compares two snip.Snippets and returns nil
// if their fields match or and error if they don't.
func CompareSnippets(have Snippet, want Snippet) error {
	if have.Name != want.Name {
		return errors.New("Name mismatch:\nhave = " + have.Name + "\nwant = " + want.Name)
	}
	if have.Trigger != want.Trigger {
		return errors.New("Trigger mismatch:\nhave = " + have.Trigger + "\nwant = " + want.Trigger)
	}
	if have.Description != want.Description {
		return errors.New("Description mismatch:\nhave = " + have.Description + "\nwant = " + want.Description)
	}
	if have.Rules != want.Rules {
		return errors.New("Rules mismatch:\nhave = " + have.Rules + "\nwant = " + want.Rules)
	}
	if have.Scope != want.Scope {
		return errors.New("Scope mismatch:\nhave = " + have.Scope + "\nwant = " + want.Scope)
	}
	for i := range have.Body {
		if have.Body[i] != want.Body[i] {
			return errors.New("Body mismatch:\nhave = " + have.Body[i] + "\nwant = " + want.Body[i])
		}
	}
	return nil
}

// FindSnippetByName finds a given snip.Snippet in an array of snip.Snippets
// by name, or returns an error if one is not found.
func FindSnippetByName(name string, snippets []Snippet) (Snippet, error) {
	for _, s := range snippets {
		if s.Name == name {
			return s, nil
		}
	}
	return Snippet{}, errors.New("No matching snippet found")

}
