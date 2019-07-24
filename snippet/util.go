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

func compareSnippets(have Snippet, want Snippet) error {
	if have.Name != want.Name {
		return errors.New("Name mismatch:\nhave = " + have.Name + "\nwant = " + want.Name)
	}
	if have.Trigger != want.Trigger {
		return errors.New("Trigger mismatch:\nhave = " + have.Trigger + "\nwant = " + want.Trigger)
	}
	if have.Rules != want.Rules {
		return errors.New("Rules mismatch:\nhave = " + have.Rules + "\nwant = " + want.Rules)
	}
	if have.Description != want.Description {
		return errors.New("Description mismatch:\nhave = " + have.Description + "\nwant = " + want.Description)
	}
	for i := range have.Body {
		if have.Body[i] != want.Body[i] {
			return errors.New("Body mismatch:\nhave = " + have.Body[i] + "\nwant = " + want.Body[i])
		}
	}
	return nil
}
