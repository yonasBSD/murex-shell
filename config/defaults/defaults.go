package defaults

import (
	"runtime"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/lang/types/define"
)

// Defaults defines the default config
func Defaults(c *config.Config, isInteractive bool) {
	c.Define("shell", "prompt", config.Properties{
		Description: "Interactive shell prompt.",
		//Default:     "{ exitnum->set: x; if { = x!=`0` } { set: prompt='\033[31m»\033[0m' } { set: prompt='\033[31m»\033[0m' }; out: murex $prompt }",
		Default:  "{ out 'murex » ' }",
		DataType: types.CodeBlock,
	})

	c.Define("shell", "prompt-multiline", config.Properties{
		Description: "Shell prompt when command line string spans multiple lines.",
		Default:     `{ out "$linenum » " }`,
		DataType:    types.CodeBlock,
	})

	c.Define("shell", "max-suggestions", config.Properties{
		Description: "Maximum number of lines with auto-completion suggestions to display.",
		Default:     10,
		DataType:    types.Integer,
	})

	c.Define("shell", "history", config.Properties{
		Description: "Write shell history (interactive shell) to disk.",
		Default:     true,
		DataType:    types.Boolean,
	})

	c.Define("shell", "add-colour", config.Properties{
		Description: "ANSI escape sequences in Murex builtins to highlight syntax errors, history completions, etc.",
		Default:     (runtime.GOOS != "windows" && isInteractive),
		DataType:    types.Boolean,
	})

	c.Define("shell", "syntax-highlighting", config.Properties{
		Description: "Syntax highlighting of murex code when in the interactive shell.",
		Default:     true,
		DataType:    types.Boolean,
	})

	c.Define("shell", "show-exts", config.Properties{
		Description: "Windows only! Auto-completes file extensions. This also affects the auto-completion parameters.",
		Default:     false,
		DataType:    types.Boolean,
	})

	//c.Define("shell", "strip-colour", config.Properties{
	//	Description: "Strips the colour codes (ANSI escape sequences) from all output destined for the terminal.",
	//	Default:     false,
	//	DataType:    types.Boolean,
	//})

	// Add config hooks for mime types
	c.Define("shell", "mime", config.Properties{
		Description: "Supported MIME types and their corresponding Murex data types.",
		Default:     define.GetMimes(),
		DataType:    types.Json,
	})

	// Add config hooks for mime types
	c.Define("shell", "extensions", config.Properties{
		Description: "Supported file extensions and their corresponding Murex data types.",
		Default:     define.GetFileExts(),
		DataType:    types.Json,
	})
}