// Package snippet provides snippet models.
package snippet

import "strings"

// Snippet models a platform independent snippet.
type Snippet struct {
	Name        string
	Trigger     string
	Description string
	Rules       string
	Scope       string
	Body        []string
}

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
	if s.Rules != "" {
		sb.WriteString("Rules:       " + s.Rules + "\n")
	}

	sb.WriteString("Body:\n")
	length := len(s.Body) - 1
	for i, str := range s.Body {
		if i == length {
			sb.WriteString(str)
		} else {
			sb.WriteString(str + "\n")
		}
	}

	return sb.String()
}
