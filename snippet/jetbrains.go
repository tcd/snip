package snippet

import (
	"encoding/xml"
	"io/ioutil"
	"strings"
)

type jbTemplateSet struct {
	Group     string       `xml:"group,attr"`
	Templates []jbTemplate `xml:"template"`
}

type jbTemplate struct {
	Name             string `xml:"name,attr"`        // name
	Value            string `xml:"value,attr"`       // body
	Description      string `xml:"description,attr"` // description
	Shortcut         string `xml:"shortcut,attr"`    // trigger
	ToReformat       string `xml:"toReformat,attr"`
	ToShortenFQNames string `xml:"toShortenFQNames,attr"`
	Variables        []jbVariable
	Context          jbContext
}

type jbVariable struct {
	Name         string `xml:"name,attr"`
	Expression   string `xml:"expression,attr"`
	DefaultValue string `xml:"defaultValue,attr"`
	AlwaysStopAt string `xml:"alwaysStopAt,attr"`
}

type jbContext struct {
	Options []jbOption
}

type jbOption struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

// ParseJetBrainsFile returns an array of snip.Snippets parsed
// from the JetBrains snippet xml file at the given path,
// or an error if something goes wrong.
func ParseJetBrainsFile(path string) ([]Snippet, error) {
	var snippets []Snippet
	var ts jbTemplateSet

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return snippets, err
	}

	err = xml.Unmarshal(bytes, &ts)
	if err != nil {
		return snippets, err
	}

	for _, snip := range ts.Templates {
		s := Snippet{
			Name:        strings.TrimSpace(snip.Name),
			Trigger:     strings.TrimSpace(snip.Name),
			Description: strings.TrimSpace(snip.Description),
			Body:        strings.Split(snip.Value, "\n"),
		}
		snippets = append(snippets, s)
	}

	return snippets, nil
}
