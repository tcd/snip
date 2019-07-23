package snippet

import (
	"io/ioutil"
	"log"
)

// File models a file and comes with some handy goodness.
type File struct {
	Path     string
	Contents []byte
}

// NewFile news up a new File.
func NewFile(path string) *File {
	file := File{
		Path:     path,
		Contents: getBytes(path),
	}
	return &file
}

// Gets a slice of bytes from a file.
func getBytes(file string) []byte {
	bytes, err := ioutil.ReadFile(file) // TODO: use `func ReadAll(r io.Reader) ([]byte, error)``
	if err != nil {
		log.Fatalf("Error reading file '%s': %s", file, err.Error())
		return []byte{}
	}
	return bytes
}

// Return a File's contents as a string.
func (f *File) String() string {
	return string(f.Contents)
}
