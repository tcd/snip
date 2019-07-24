package snippet

import (
	"io/ioutil"
	"log"
	"regexp"
)

// Gets a slice of bytes from a file.
func bytesFromFile(file string) []byte {
	bytes, err := ioutil.ReadFile(file) // TODO: use `func ReadAll(r io.Reader) ([]byte, error)``
	if err != nil {
		log.Fatalf("Error reading file '%s': %s", file, err.Error())
		return []byte{}
	}
	return bytes
}

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
