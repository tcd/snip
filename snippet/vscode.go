package snippet

import (
	"encoding/json"
	"log"
)

// ParseCodeFile parses a VS Code snippet file.
func ParseCodeFile(path string) []Snippet {
	bytes := bytesFromFile(path)
	emptyInterface := makeGenericInterface(bytes)
	codeSnippets := parseUnknownJSON(emptyInterface)
	return codeSnippets
}

// Unmarshal JSON bytes using a generic interface
func makeGenericInterface(bytes []byte) interface{} {
	var x interface{}
	err := json.Unmarshal(bytes, &x)
	if err != nil {
		log.Println("Error parsing JSON: ", err)
	}
	return x
}

// See: https://gist.github.com/mjohnsullivan/24647cae50928a34b5cc
func parseUnknownJSON(x interface{}) []Snippet {
	var snippets []Snippet

	itemsMap := x.(map[string]interface{})

	for _, v := range itemsMap {
		switch jsonObj := v.(type) {
		case interface{}:
			var snippet Snippet
			for itemKey, itemValue := range jsonObj.(map[string]interface{}) {
				switch itemKey {
				case "prefix":
					switch itemValue := itemValue.(type) {
					case string:
						snippet.Trigger = itemValue
					default:
						break
					}
				case "description":
					switch itemValue := itemValue.(type) {
					case string:
						snippet.Description = itemValue
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
						snippet.Body = body
					case string:
						snippet.Body = []string{itemValue}
					default:
						break
					}
				default:
					break
				}
			}
			snippets = append(snippets, snippet)
		default:
			break
		}
	}
	return snippets
}
