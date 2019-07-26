# snippet

## Features

- Tabstops
  - Placeholders
  - Mirroring
  - Variables
    - [JetBrains](https://www.jetbrains.com/help/idea/template-variables.html#predefined_functions)
    - [Sublime](http://docs.sublimetext.info/en/latest/extensibility/snippets.html#environment-variables)
    - [VS Code](https://code.visualstudio.com/docs/editor/userdefinedsnippets#_variables)
    - [TextMate (v1)](https://macromates.com/manual/en/environment_variables)
- Code Execution
  - TextMate
    - Bash
    - Perl
    - PHP
    - Python
    - Ruby
  - Ultisnips
    - Python
    - Shell
    - VimL
  - YASnippet
    - ELisp

### Variables

|                   |          JetBrains           |       Sublime       |      TextMate       |       VS Code       |       Ultisnips        |
| ----------------- | ---------------------------- | ------------------- | ------------------- | ------------------- | ---------------------- |
| filename          | `fileName()`                 | `$TM_FILENAME`      |                     | `$TM_FILENAME`      | `!v expand('%:t')`     |
| base filename     | `fileNameWithoutExtension()` |                     |                     | `$TM_FILENAME_BASE` | `!v expand('%:r')`     |
| filepath          |                              | `$TM_FILEPATH`      | `$TM_FILEPATH`      | `$TM_FILEPATH`      | `!v expand('%:p')`     |
| current directory |                              | `$TM_DIRECTORY`     | `$TM_DIRECTORY`     | `$TM_DIRECTORY`     | `!v expand('%:p:h:t')` |
| current line      |                              | `$TM_CURRENT_LINE`  | `$TM_CURRENT_LINE`  | `$TM_CURRENT_LINE`  | `!v getline('.')`      |
| current word      |                              | `$TM_CURRENT_WORD`  | `$TM_CURRENT_WORD`  | `$TM_CURRENT_WORD`  | `!v expand('<cword>')` |
| line number       |                              | `$TM_LINE_NUMBER`   | `$TM_LINE_NUMBER`   | `$TM_LINE_NUMBER`   | `!v line('.')`         |
| current highlight |                              | `$TM_SELECTED_TEXT` | `$TM_SELECTED_TEXT` | `$TM_SELECTED_TEXT` | `${VISUAL}`            |
| tabs or spaces    |                              | `$TM_SOFT_TABS`     | `$TM_SOFT_TABS`     |                     |                        |
| tab size          |                              | `$TM_TAB_SIZE`      | `$TM_TAB_SIZE`      |                     | `!v &tabsize`          |
| clipboard         | `clipboard()`                |                     |                     | `$CLIPBOARD`        | `!v getreg('+')`       |

---

## TextMate

- [TextMate Snippets](http://manual.macromates.com/en/snippets)
- Samples
  - [textmate/javascript.tmbundle](https://github.com/textmate/javascript.tmbundle/tree/master/Snippets)
  - [textmate/html.tmbundle](https://github.com/textmate/html.tmbundle/tree/master/Snippets)
  - [textmate/ruby.tmbundle](https://github.com/textmate/ruby.tmbundle/tree/master/Snippets)
  - [textmate/python.tmbundle](https://github.com/textmate/python.tmbundle/tree/master/Snippets)
  - [textmate/swift.tmbundle](https://github.com/textmate/swift.tmbundle/tree/master/Snippets)
  - [elixir-editors/elixir-tmbundle](https://github.com/elixir-editors/elixir-tmbundle/tree/master/Snippets)

### [Variables](https://macromates.com/manual/en/environment_variables)

| variable                | description                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |
| ----------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `$TM_BUNDLE_SUPPORT`    | shell commands which are (indirectly) triggered from a bundle item (which could be a Command, Drag Command, Macro, or Snippet) will have this variable pointing to the Support folder of the bundle that ran the item, if such a folder exists. In addition, $TM_BUNDLE_SUPPORT/bin will be added to the path.                                                                                                                                                                 |
| `$TM_CURRENT_LINE`      | textual content of the current line.                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| `$TM_CURRENT_WORD`      | the word in which the caret is located.                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| `$TM_DIRECTORY`         | the folder of the current document (may not be set).                                                                                                                                                                                                                                                                                                                                                                                                                           |
| `$TM_FILEPATH`          | path (including file name) for the current document (may not be set).                                                                                                                                                                                                                                                                                                                                                                                                          |
| `$TM_LINE_INDEX`        | the index in the current line which marks the caret’s location. This index is zero-based and takes the utf-8 encoding of the line (e.g. read as TM_CURRENT_LINE) into account. So to split a line into what is to the left and right of the caret you could do:                                                                                                                                                                                                                |
| `$TM_LINE_NUMBER`       | the carets line position (counting from 1). For example if you need to work with the part of the document above the caret you can set the commands input to “Entire Document” and use the following to cut off the part below and including the current line:                                                                                                                                                                                                                  |
| `$TM_PROJECT_DIRECTORY` | the top-level folder in the project drawer (may not be set).                                                                                                                                                                                                                                                                                                                                                                                                                   |
| `$TM_SCOPE`             | the scope that the caret is inside. See scope selectors for information about scopes.                                                                                                                                                                                                                                                                                                                                                                                          |
| `$TM_SELECTED_FILES`    | space separated list of the files and folders selected in the project drawer (may not be set). The paths are shell-escaped, so to use these, you need to prefix the line with eval (to make the shell re-evaluate the line, after expanding the variable). For example to run the file command on all selected files in the project drawer, the shell command would be:                                                                                                        |
| `$TM_SELECTED_FILE`     | full path of the first selected file or folder in the project drawer (may not be set).                                                                                                                                                                                                                                                                                                                                                                                         |
| `$TM_SELECTED_TEXT`     | full content of the selection (may not be set). Note that environment variables have a size limitation of roughly 64 KB, so if the user selects more than that, this variable will not reflect the actual selection (commands that need to work with the selection should generally set this to be the standard input).                                                                                                                                                        |
| `$TM_SOFT_TABS`         | this will have the value YES if the user has enabled soft tabs, otherwise it has the value NO. This is useful when a shell command generates an indented result and wants to match the users preferences with respect to tabs versus spaces for the indent.                                                                                                                                                                                                                    |
| `$TM_SUPPORT_PATH`      | the TextMate application bundle contains a support folder with several items which are used by some of the default commands (for example CocoaDialog, Markdown, the SCM commit window, Textile, tidy, etc.). This variable points to that support folder. Generally you would not need to use the variable directly since $TM_SUPPORT_PATH/bin is added to the path, so using some of the bundled commands can be done without having to specify their full path.              |
| `$TM_TAB_SIZE`          | the tab size as shown in the status bar. This is useful when creating commands which need to present the current document in another form (Tidy, convert to HTML or similar) or generate a result which needs to match the tab size of the document. See also TM_SOFT_TABS.                                                                                                                                                                                                    |

## Atom

- [Atom Flight Manual » Snippets](https://flight-manual.atom.io/using-atom/sections/snippets/)
- [atom/snippets](https://github.com/atom/snippets)
- Samples
  - [atom/language-html](https://github.com/atom/language-html/blob/master/snippets/language-html.cson)
  - [atom/language-javascript](https://github.com/atom/language-javascript/blob/master/snippets/language-javascript.cson)
  - [webbushka/atom-react-snippets](https://github.com/webbushka/atom-react-snippets)
  - [joseramonc/rails-snippets](https://github.com/joseramonc/rails-snippets)

## Emacs

- [joaotavora/yasnippet](https://github.com/joaotavora/yasnippet)
- [emacswiki - Yasnippet](https://www.emacswiki.org/emacs/Yasnippet)
- Samples

## Sublime

- [Extending Sublime Text » Snippets](http://docs.sublimetext.info/en/latest/extensibility/snippets.html)
- [Meduim - Sublime3 Snippets](https://medium.freecodecamp.org/a-guide-to-preserving-your-wrists-with-sublime-text-snippets-7541662a53f2)
- [~~bobthecow/sublime-sane-snippets~~](https://github.com/bobthecow/sublime-sane-snippets)
- Samples
  - [babel/babel-sublime-snippets](https://github.com/babel/babel-sublime-snippets)
  - [JasonMortonNZ/bs3-sublime-plugin](https://github.com/JasonMortonNZ/bs3-sublime-plugin)
  - [tadast/sublime-rails-snippets](https://github.com/tadast/sublime-rails-snippets)
  - [jprichardson/sublime-js-snippets](https://github.com/jprichardson/sublime-js-snippets)
  - [zenorocha/sublime-javascript-snippets](https://github.com/zenorocha/sublime-javascript-snippets)

### [Variables](http://docs.sublimetext.info/en/latest/extensibility/snippets.html#environment-variables)

| variable             | description                                                           |
|----------------------|-----------------------------------------------------------------------|
| `$PARAM1 .. $PARAMn` | Arguments passed to the `insert_snippet` command. (Not covered here.) |
| `$SELECTION`         | The text that was selected when the snippet was triggered.            |
| `$TM_CURRENT_LINE`   | Content of the cursor’s line when the snippet was triggered.          |
| `$TM_CURRENT_WORD`   | Word under the cursor when the snippet was triggered.                 |
| `$TM_DIRECTORY`      | Directory name of the file being edited. (since 3154)                 |
| `$TM_FILENAME`       | Name of the file being edited, including extension.                   |
| `$TM_FILEPATH`       | Path to the file being edited.                                        |
| `$TM_FULLNAME`       | User's user name.                                                     |
| `$TM_LINE_INDEX`     | Column where the snippet is being inserted, 0 based.                  |
| `$TM_LINE_NUMBER`    | Row where the snippet is being inserted, 1 based.                     |
| `$TM_SELECTED_TEXT`  | An alias for `$SELECTION`.                                            |
| `$TM_SCOPE`          | The scope of the beginning of each selected region. (since 3154)      |
| `$TM_SOFT_TABS`      | `YES` if `translate_tabs_to_spaces` is true, otherwise `NO`.          |
| `$TM_TAB_SIZE`       | Spaces per-tab (controlled by the `tab_size` option).                 |

## Vim
- [vim snippet engines](http://vim-wiki.mawercer.de/wiki/topic/text-snippets-skeletons-templates.html)
  - [garbas/vim-snipmate](https://github.com/garbas/vim-snipmate)
  - [SirVer/ultisnips](https://github.com/SirVer/ultisnips)
  - [Shougo/neosnippet.vim](https://github.com/Shougo/neosnippet.vim)
  - [Shougo/deoppet.nvim](https://github.com/Shougo/deoppet.nvim)
- Samples
  - [honza/vim-snippets](https://github.com/honza/vim-snippets)
  - [Shougo/neosnippet-snippets](https://github.com/Shougo/neosnippet-snippets)

### Ultisnips

#### Basic Syntax
```
snippet trigger_word [ "description" [ options ] ]
snippet "tab trigger" [ "description" [ options ] ]
snippet !tab trigger! [ "description" [ options ] ]
snippet !"tab trigger"! [ "description" [ options ] ]
```

#### Rules

`Abeimrstw`

#### Special Characters

If you want to insert one of these characters literally, escape them with a backslash, '\'.
```
'`', '{', '$' and '\'
```

## VS Code
- [Creating your own Snippets](https://code.visualstudio.com/docs/editor/userdefinedsnippets)
- Samples
  - [xabikos/vscode-react](https://github.com/xabikos/vscode-react)
  - [Snippet Extensions](https://marketplace.visualstudio.com/search?target=VSCode&category=Snippets&sortBy=Downloads)
  - [`microsoft/vscode/extensions/*/snippets/*`](https://github.com/microsoft/vscode/tree/master/extensions)

### [Variables](https://code.visualstudio.com/docs/editor/userdefinedsnippets#_variables)

TM_SELECTED_TEXT The currently selected text or the empty string
TM_CURRENT_LINE The contents of the current line
TM_CURRENT_WORD The contents of the word under cursor or the empty string
TM_LINE_INDEX The zero-index based line number
TM_LINE_NUMBER The one-index based line number
TM_FILENAME The filename of the current document
TM_FILENAME_BASE The filename of the current document without its extensions
TM_DIRECTORY The directory of the current document
TM_FILEPATH The full file path of the current document
CLIPBOARD The contents of your clipboard
WORKSPACE_NAME The name of the opened workspace or folder

BLOCK_COMMENT_START Example output: in PHP /* or in HTML <!--
BLOCK_COMMENT_END Example output: in PHP */ or in HTML -->
LINE_COMMENT Example output: in PHP // or in HTML <!-- -->

CURRENT_YEAR The current year
CURRENT_YEAR_SHORT The current year's last two digits
CURRENT_MONTH The month as two digits (example '02')
CURRENT_MONTH_NAME The full name of the month (example 'July')
CURRENT_MONTH_NAME_SHORT The short name of the month (example 'Jul')
CURRENT_DATE The day of the month
CURRENT_DAY_NAME The name of day (example 'Monday')
CURRENT_DAY_NAME_SHORT The short name of the day (example 'Mon')
CURRENT_HOUR The current hour in 24-hour clock format
CURRENT_MINUTE The current minute
CURRENT_SECOND The current second

## Xcode
- Samples
  - [burczyk/XcodeSwiftSnippets](https://github.com/burczyk/XcodeSwiftSnippets)

## Visual Studio
- [Visual Studio](https://docs.microsoft.com/en-us/visualstudio/ide/code-snippets?view=vs-2019)
- Samples

## JetBrains

- Samples
  - [blakedietz/js-live-template](https://github.com/blakedietz/js-live-template)
  - [MrZaYaC/ng2-webstorm-snippets](https://github.com/MrZaYaC/ng2-webstorm-snippets)
  - [Krbz/webstorm-react-snippets](https://github.com/Krbz/webstorm-react-snippets)

### [Variables](https://www.jetbrains.com/help/idea/template-variables.html#predefined_functions)
