package docs

func init() {

	Definition["read"] = "# `read` - Command Reference\n\n> `read` a line of input from the user and store as a variable\n\n## Description\n\nA readline function to allow a line of data inputed from the terminal.\n\n## Usage\n\nClassic usage:\n\n    read: \"prompt\" var_name\n    \n    <stdin> -> read: var_name\n    \nScript usage:\n\n    read: [ --prompt \"prompt\"         ]\n          [ --variable var_name       ]\n          [ --default \"default value\" ]\n          [ --datatype data-type      ]\n          [ --mask character          ]\n\n## Examples\n\nClassic usage:\n\n    read: \"What is your name? \" name\n    out: \"Hello $name\"\n    \n    out: What is your name? -> read: name\n    out: \"Hello $name\"\n    \nScript usage:\n\n    read: --prompt \"Are you sure? [Y/n]\" \\\n          --variable yn \\\n          --default Y\n    \nSecrets:\n\n    read: --prompt \"Password: \" --variable pw --mask *\n\n## Flags\n\n* `--datatype`\n    Murex data-type for the read data (default: str)\n* `--default`\n    If a zero length string is returned but neither ctrl+c nor ctrl+d were pressed, then the default value defined here will be returned\n* `--mask`\n    Optional password mask, for reading secrets\n* `--prompt`\n    User notification to display\n* `--variable`\n    Variable name to store the read data (default: read)\n\n## Detail\n\n### Classic Usage\n\nIf `read` is called as a method then the prompt string is taken from STDIN.\nOtherwise the prompt string will be the first parameter. However if no prompt\nstring is given then `read` will not write a prompt.\n\nThe last parameter will be the variable name to store the string read by `read`.\nThis variable cannot be prefixed by dollar, `$`, otherwise the shell will write\nthe output of that variable as the last parameter rather than the name of the\nvariable.\n\nThe data type the `read` line will be stored as is `str` (string). If you\nrequire this to be different then please use `tread` (typed read) or call `read`\nwith the `--datatype` flag as per the **script usage**.\n\n## See Also\n\n* [`(` (brace quote)](../commands/brace-quote.md):\n  Write a string to the STDOUT without new line\n* [`>>` (append file)](../commands/greater-than-greater-than.md):\n  Writes STDIN to disk - appending contents if file already exists\n* [`>` (truncate file)](../commands/greater-than.md):\n  Writes STDIN to disk - overwriting contents if file already exists\n* [`cast`](../commands/cast.md):\n  Alters the data type of the previous function without altering it's output\n* [`err`](../commands/err.md):\n  Print a line to the STDERR\n* [`out`](../commands/out.md):\n  Print a string to the STDOUT with a trailing new line character\n* [`tout`](../commands/tout.md):\n  Print a string to the STDOUT and set it's data-type\n* [`tread`](../commands/tread.md):\n  `read` a line of input from the user and store as a user defined *typed* variable"

}
