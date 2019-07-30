package snippet

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

// Yasnippet returns a Snippet in YASnippet format.
func (s Snippet) Yasnippet() string {
	var sb strings.Builder

	sb.WriteString("# -*- mode: snippet -*-\n")
	if s.Name != "" {
		sb.WriteString("# name: " + s.Name + "\n")
	}
	if s.Trigger != "" {
		sb.WriteString("# key: " + s.Trigger + "\n")
	}
	sb.WriteString("# type: snippet\n")
	sb.WriteString("# --\n")
	sb.WriteString(strings.Join(s.Body, "\n"))

	return sb.String()
}

type yasnippetSnippet struct {
	Key         string   `raw:"# key:"`  // trigger
	Name        string   `raw:"# name:"` // name; filename by default
	Condition   string   `raw:"# condition:"`
	Group       string   `raw:"# group:"`
	ExpandEnv   string   `raw:"# expand-env:"` // a list of lists assigning values to variables
	Binding     string   `raw:"# binding:"`
	Type        string   `raw:"# type:"` // snippet or command; we can only work with a snippet
	UUID        string   `raw:"# uuid:"`
	Contributor string   `raw:"# contributor:"`
	Body        []string `raw:"# --"`
}

// ParseYasnippetFile parses a YASnippet file at the given path and
// returns a Snippet or an error if something goes wrong.
func ParseYasnippetFile(path string) (Snippet, error) {
	var s Snippet

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return s, err
	}

	s = parseYasnippet(string(bytes), path)

	return s, nil
}

// parseYasnippet only gets the minimal ammount of info needed
// for a Snippet, much is left on the table.
func parseYasnippet(snip string, path string) Snippet {
	var s Snippet

	lines := strings.Split(snip, "\n")

	for i, line := range lines {
		if len(line) == 0 {
			continue
		}
		if strings.Contains(line, "# name: ") {
			s.Name = strings.TrimSpace(line[7:])
		}
		if strings.Contains(line, "# key: ") {
			s.Trigger = strings.TrimSpace(line[6:])
		}
		if strings.Contains(line, "# --") {
			s.Body = lines[i+1:]
		}
	}

	if s.Name == "" {
		s.Name = filepath.Base(path)
	}

	return s
}

// func parseYasnippetFull() yasnippetSnippet {
// 	var s yasnippetSnippet
// 	return s
// }
