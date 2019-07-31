// Package snippet provides snippet models.
package snippet

import "strings"

// Snippet models a platform independent snippet.
type Snippet struct {
	Name        string
	Trigger     string
	Description string
	Scope       string
	Body        []string
}

// String() method for use by Printf, Println, etc.
func (s Snippet) String() string {
	var sb strings.Builder

	if s.Name != "" {
		sb.WriteString("Name:        " + s.Name + "\n")
	}
	if s.Trigger != "" {
		sb.WriteString("Trigger:     " + s.Trigger + "\n")
	}
	if s.Description != "" {
		sb.WriteString("Description: " + s.Description + "\n")
	}
	if s.Scope != "" {
		sb.WriteString("Scope:       " + s.Scope + "\n")
	}
	sb.WriteString("Body:\n")
	sb.WriteString(strings.Join(s.Body, "\n"))

	return sb.String()
}
