package docs

func init() {

	Definition["private"] = "# `private`\n\n> Define a private function block\n\n## Description\n\n`private` defines a function who's scope is limited to that module or source\nfile.\n\nPrivates cannot be called from one module to another (unless they're wrapped\naround a global `function`) and nor can they be called from the interactive\ncommand line. The purpose of a `private` is to reduce repeated code inside\na module or source file without cluttering up the global namespace.\n\n## Usage\n\n```\nprivate: name { code-block }\n```\n\n## Examples\n\n```\n# The following cannot be entered via the command line. You need to write\n# it to a file and execute it from there.\n\nprivate hw {\n    out \"Hello, World!\"\n}\n\nfunction tom {\n    hw\n    out \"My name is Tom.\"\n}\n\nfunction dick {\n    hw\n    out \"My name is Dick.\"\n}\n\nfunction harry {\n    hw\n    out \"My name is Harry.\"\n}\n```\n\n## Detail\n\n### Allowed characters\n\nPrivate names can only include any characters apart from dollar (`$`).\nThis is to prevent functions from overwriting variables (see the order of\npreference below).\n\n### Undefining a private\n\nBecause private functions are fixed to the source file that declares them,\nthere isn't much point in undefining them. Thus at this point in time, it\nis not possible to do so.\n\n### Order of preference\n\nThere is an order of precedence for which commands are looked up:\n\n1. `runmode`: this is executed before the rest of the script. It is invoked by\n   the pre-compiler forking process and is required to sit at the top of any\n   scripts.\n\n1. `test` and `pipe` functions also alter the behavior of the compiler and thus\n   are executed ahead of any scripts.\n\n4. private functions - defined via `private`. Private's cannot be global and\n   are scoped only to the module or source that defined them. For example, You\n   cannot call a private function directly from the interactive command line\n   (however you can force an indirect call via `fexec`).\n\n2. Aliases - defined via `alias`. All aliases are global.\n\n3. Murex functions - defined via `function`. All functions are global.\n\n5. Variables (dollar prefixed) which are declared via `global`, `set` or `let`.\n   Also environmental variables too, declared via `export`.\n\n6. globbing: however this only applies for commands executed in the interactive\n   shell.\n\n7. Murex builtins.\n\n8. External executable files\n\nYou can override this order of precedence via the `fexec` and `exec` builtins.\n\n## See Also\n\n* [`alias`](../commands/alias.md):\n  Create an alias for a command\n* [`break`](../commands/break.md):\n  Terminate execution of a block within your processes scope\n* [`exec`](../commands/exec.md):\n  Runs an executable\n* [`export`](../commands/export.md):\n  Define an environmental variable and set it's value\n* [`fexec` ](../commands/fexec.md):\n  Execute a command or function, bypassing the usual order of precedence.\n* [`function`](../commands/function.md):\n  Define a function block\n* [`g`](../commands/g.md):\n  Glob pattern matching for file system objects (eg `*.txt`)\n* [`global`](../commands/global.md):\n  Define a global variable and set it's value\n* [`let`](../commands/let.md):\n  Evaluate a mathematical function and assign to variable (deprecated)\n* [`method`](../commands/method.md):\n  Define a methods supported data-types\n* [`set`](../commands/set.md):\n  Define a local variable and set it's value\n* [`source`](../commands/source.md):\n  Import Murex code from another file of code block"

}
