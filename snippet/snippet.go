// Package snippet provides snippet models.
package snippet

import (
	"fmt"
	"strings"
)

// Snippet models a platform independent snippet.
type Snippet struct {
	Name        string
	Trigger     string
	Description string
	Rules       string
	Body        []string
}

func (s Snippet) String() string {
	return fmt.Sprint(
		"Name:        ", s.Name, "\n",
		"Trigger:     ", s.Trigger, "\n",
		"Description: ", s.Description, "\n",
		"Rules:       ", s.Rules, "\n",
		"Body:        ", "\n", bodyLoop(s.Body),
	)
}

func bodyLoop(body []string) string {
	var sb strings.Builder
	for _, str := range body {
		sb.WriteString(str + "\n")
	}
	return sb.String()
}
