package snippet

type xcodeSnippet struct {
	Trigger          string   `plist:"IDECodeSnippetCompletionPrefix"`
	Title            string   `plist:"IDECodeSnippetTitle"`
	Scopes           string   `plist:"IDECodeSnippetLanguage"`
	Contents         string   `plist:"IDECodeSnippetContents"`
	CompletionScopes []string `plist:"IDECodeSnippetCompletionScopes"`
	UUID             string   `plist:"IDECodeSnippetIdentifier"`
	UserSnippet      bool     `plist:"IDECodeSnippetUserSnippet"`
	SnippetVersion   int      `plist:"IDECodeSnippetVersion"`
}
