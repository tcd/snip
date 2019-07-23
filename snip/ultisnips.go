// Package snippet provides snippet models.
package snippet

import (
	"io/ioutil"
	"log"
	"regexp"
)

// ParseFile returns an array of snippets in a file at the given path.
func ParseFile(path string) []string {
	bytes := getBytes(path)
	re := regexp.MustCompile(`(?msU:((^snippet.+$\n)(.+)(?:endsnippet)))`)
	matches := re.FindAllString(string(bytes), -1) // TODO: Find out if FindAll() is faster than FindAllString()
	return matches
}

// Gets a slice of bytes from a file; return value will be empty if an error occurs.
func getBytes(file string) []byte {
	bytes, err := ioutil.ReadFile(file) // TODO: use `func ReadAll(r io.Reader) ([]byte, error)``
	if err != nil {
		log.Fatalf("Demonic Invasion In Progress: %s", err.Error())
		return []byte{}
	}
	return bytes
}
