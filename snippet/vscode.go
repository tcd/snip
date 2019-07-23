package snippet

import (
	"encoding/json"
	"fmt"
)

// CodeSnippet models raw a snippet parsed from a JSON file.
type CodeSnippet struct {
	Prefix      string
	Body        []string
	Description string
}

// ParseCodeFile parses a VS Code snippet file.
func ParseCodeFile(path string) []Snippet {
	bytes := getBytes(path)
	emptyInterface := makeGenericInterface(bytes)
	codeSnippets := parseUnknownJSON(emptyInterface)

	var snippets []Snippet

	for _, snippet := range codeSnippets {
		snippets = append(snippets, codeSnippetToSnippet(snippet))
	}
	return snippets
}

func codeSnippetToSnippet(snippet CodeSnippet) Snippet {
	newSnippet := Snippet{
		Name:        snippet.Prefix,
		Body:        snippet.Body,
		Description: snippet.Description,
	}
	return newSnippet
}

// Unmarshal JSON bytes using a generic interface
func makeGenericInterface(bytes []byte) interface{} {
	var x interface{}
	err := json.Unmarshal(bytes, &x)
	if err != nil {
		fmt.Println("Error parsing JSON: ", err)
	}
	return x
}

// See: https://gist.github.com/mjohnsullivan/24647cae50928a34b5cc
func parseUnknownJSON(x interface{}) []CodeSnippet {
	var snippets []CodeSnippet

	itemsMap := x.(map[string]interface{})

	for _, v := range itemsMap {
		switch jsonObj := v.(type) {
		case interface{}:
			var item CodeSnippet
			for itemKey, itemValue := range jsonObj.(map[string]interface{}) {
				switch itemKey {
				case "prefix":
					switch itemValue := itemValue.(type) {
					case string:
						item.Prefix = itemValue
					default:
						break
					}
				case "description":
					switch itemValue := itemValue.(type) {
					case string:
						item.Description = itemValue
					default:
						break
					}
				case "body":
					switch itemValue := itemValue.(type) {
					case []interface{}:
						var body []string
						for _, v := range itemValue {
							body = append(body, v.(string))
						}
						item.Body = body
					case string:
						item.Body = []string{itemValue}
					default:
						break
					}
				default:
					break
				}
			}
			snippets = append(snippets, item)
		default:
			break
		}
	}
	return snippets
}
