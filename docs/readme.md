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

|                   |       Sublime       |      TextMate       |       VS Code       |       Ultisnips        |          JetBrains           |
| ----------------- | ------------------- | ------------------- | ------------------- | ---------------------- | ---------------------------- |
| filename          | `$TM_FILENAME`      |                     | `$TM_FILENAME`      | `!v expand('%:t')`     | `fileName()`                 |
| base filename     |                     |                     | `$TM_FILENAME_BASE` | `!v expand('%:r')`     | `fileNameWithoutExtension()` |
| filepath          | `$TM_FILEPATH`      | `$TM_FILEPATH`      | `$TM_FILEPATH`      | `!v expand('%:p')`     |                              |
| current directory | `$TM_DIRECTORY`     | `$TM_DIRECTORY`     | `$TM_DIRECTORY`     | `!v expand('%:p:h:t')` |                              |
| current line      | `$TM_CURRENT_LINE`  | `$TM_CURRENT_LINE`  | `$TM_CURRENT_LINE`  | `!v getline('.')`      |                              |
| current word      | `$TM_CURRENT_WORD`  | `$TM_CURRENT_WORD`  | `$TM_CURRENT_WORD`  | `!v expand('<cword>')` |                              |
| line number       | `$TM_LINE_NUMBER`   | `$TM_LINE_NUMBER`   | `$TM_LINE_NUMBER`   | `!v line('.')`         |                              |
| current highlight | `$TM_SELECTED_TEXT` | `$TM_SELECTED_TEXT` | `$TM_SELECTED_TEXT` | `${VISUAL}`            |                              |
| tabs or spaces    | `$TM_SOFT_TABS`     | `$TM_SOFT_TABS`     |                     |                        |                              |
| tab size          | `$TM_TAB_SIZE`      | `$TM_TAB_SIZE`      |                     | `!v &tabsize`          |                              |
| clipboard         |                     |                     | `$CLIPBOARD`        | `!v getreg('+')`       | `clipboard()`                |

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
