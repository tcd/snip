// Package snippet provides snippet models.
package snippet

import "fmt"

// Snippet models a platform independent snippet.
type Snippet struct {
	Name        string
	Trigger     string
	Rules       string
	Body        []string
	Description string
}

func (s Snippet) String() string {
	return fmt.Sprint(
		"Name: ", s.Name, "\n",
		"Trigger: ", s.Trigger, "\n",
		"Rules: ", s.Rules, "\n",
		"Body: ", s.Body, "\n",
		"Description: ", s.Description,
	)
}
