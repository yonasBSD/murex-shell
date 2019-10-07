package shell

import (
	"strings"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/shell/autocomplete"
	"github.com/lmorg/murex/utils/ansi"
	"github.com/lmorg/readline"
)

func errCallback(err error) {
	s := err.Error()
	if ansi.IsAllowed() {
		s = ansi.Reset + ansi.FgRed + s
	}
	Prompt.SetHintText(s)
}

func tabCompletion(line []rune, pos int, dtc readline.DelayedTabContext) (prefix string, items []string, descriptions map[string]string, tdt readline.TabDisplayType) {
	descriptions = make(map[string]string)

	if len(line) > pos-1 {
		line = line[:pos]
	}

	pt, _ := parse(line)

	switch {
	case pt.Variable != "":
		var s string
		if pt.VarLoc < len(line) {
			s = strings.TrimSpace(string(line[pt.VarLoc:]))
		}
		s = pt.Variable + s
		prefix = s

		items = autocomplete.MatchVars(s)

	case pt.ExpectFunc:
		var s string
		if pt.Loc < len(line) {
			s = strings.TrimSpace(string(line[pt.Loc:]))
		}
		prefix = s
		items = autocomplete.MatchFunction(s, errCallback, &dtc)

	default:
		var s string
		if len(pt.Parameters) > 0 {
			s = pt.Parameters[len(pt.Parameters)-1]
		}
		prefix = s

		autocomplete.InitExeFlags(pt.FuncName)

		pIndex := 0
		items = autocomplete.MatchFlags(autocomplete.ExesFlags[pt.FuncName], s, pt.FuncName, pt.Parameters, &pIndex, &descriptions, &tdt, errCallback, &dtc)
	}

	v, err := lang.ShellProcess.Config.Get("shell", "max-suggestions", types.Integer)
	if err != nil {
		v = 4
	}
	Prompt.MaxTabCompleterRows = v.(int)

	autocomplete.FormatSuggestionsArray(pt, items)
	autocomplete.FormatSuggestionsMap(pt, descriptions)

	return
}
