package sqlselect

import (
	"encoding/json"
	"testing"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/test/count"
)

func inlineJson(v any) string {
	b, _ := json.Marshal(v)
	return string(b)
}

type Str2IfaceT struct {
	Input  []string
	Max    int
	Output []string
}

func TestStringToInterfaceTrim(t *testing.T) {
	tests := []Str2IfaceT{
		{
			Input:  []string{"a", "b", "c", "d", "e"},
			Max:    0,
			Output: []string{},
		},
		{
			Input:  []string{"a", "b", "c", "d", "e"},
			Max:    1,
			Output: []string{"a"},
		},
		{
			Input:  []string{"a", "b", "c", "d", "e"},
			Max:    2,
			Output: []string{"a", "b"},
		},
		{
			Input:  []string{"a", "b", "c", "d", "e"},
			Max:    3,
			Output: []string{"a", "b", "c"},
		},
		{
			Input:  []string{"a", "b", "c", "d", "e"},
			Max:    4,
			Output: []string{"a", "b", "c", "d"},
		},
		{
			Input:  []string{"a", "b", "c", "d", "e"},
			Max:    5,
			Output: []string{"a", "b", "c", "d", "e"},
		},
		{
			Input:  []string{"a", "b", "c", "d", "e"},
			Max:    6,
			Output: []string{"a", "b", "c", "d", "e", ""},
		},
		{
			Input:  []string{"a", "b", "c", "d", "e"},
			Max:    7,
			Output: []string{"a", "b", "c", "d", "e", "", ""},
		},
		{
			Input:  []string{"a", "b", "c", "d", "e"},
			Max:    8,
			Output: []string{"a", "b", "c", "d", "e", "", "", ""},
		},
		/////
		{
			Input:  []string{},
			Max:    5,
			Output: []string{"", "", "", "", ""},
		},
		{
			Input:  []string{"a"},
			Max:    5,
			Output: []string{"a", "", "", "", ""},
		},
		{
			Input:  []string{"a", "b"},
			Max:    5,
			Output: []string{"a", "b", "", "", ""},
		},
		{
			Input:  []string{"a", "b", "c"},
			Max:    5,
			Output: []string{"a", "b", "c", "", ""},
		},
		{
			Input:  []string{"a", "b", "c", "d"},
			Max:    5,
			Output: []string{"a", "b", "c", "d", ""},
		},
		{
			Input:  []string{"a", "b", "c", "d", "e"},
			Max:    5,
			Output: []string{"a", "b", "c", "d", "e"},
		},
		{
			Input:  []string{"a", "b", "c", "d", "e", "f"},
			Max:    5,
			Output: []string{"a", "b", "c", "d", "e"},
		},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		actual := stringToInterfaceTrim(test.Input, test.Max)

		if len(test.Output) != len(actual) {
			t.Errorf("Length mismatch in test %d", i)
			t.Logf("  Input:    %v", test.Input)
			t.Logf("  Max:      %d", test.Max)
			t.Logf("  Expected: %v", test.Output)
			t.Logf("  Actual:   %v", actual)
		}

		for j := range test.Output {
			if test.Output[j] != actual[j].(string) {
				t.Errorf("Value mismatch in test %d[%d]", i, j)
				t.Logf("  Input:    %v", test.Input)
				t.Logf("  Max:      %d", test.Max)
				t.Logf("  Expected: %v", test.Output)
				t.Logf("  Actual:   %v", actual)
				t.Logf("  Expected: '%s'", test.Output[j])
				t.Logf("  Actual:   '%s'", actual[j].(string))
			}
		}
	}
}

type DissectParametersT struct {
	Input      []string
	IsMethod   bool
	Output     string
	Error      bool
	FileName   string
	NamedPipes []string
	Variables  []string
}

func TestDissectParameters(t *testing.T) {
	tests := []DissectParametersT{

		// is a method

		{
			Input:    []string{"FROM", "file.csv", "ORDER BY", "1"},
			IsMethod: true,
			Output:   "",
			FileName: "",
			Error:    true,
		},
		{
			Input:    []string{"*", "FROM", "file.csv", "ORDER BY", "1"},
			IsMethod: true,
			Output:   "",
			FileName: "",
			Error:    true,
		},
		{
			Input:    []string{"*", "FROM", "file.csv"},
			IsMethod: true,
			Output:   "",
			FileName: "",
			Error:    true,
		},
		{
			Input:    []string{"FROM", "file.csv"},
			IsMethod: true,
			Output:   "",
			FileName: "",
			Error:    true,
		},
		{
			Input:    []string{"FROM", "file name with space.csv"},
			IsMethod: true,
			Output:   "",
			FileName: "",
			Error:    true,
		},
		{
			Input:    []string{"a", "b", "c", "FROM", "file.csv"},
			IsMethod: true,
			Output:   "",
			FileName: "",
			Error:    true,
		},
		{
			Input:    []string{"FROM", "file.csv", "ORDER BY", "1", "2", "3"},
			IsMethod: true,
			Output:   "",
			FileName: "",
			Error:    true,
		},
		{
			Input:    []string{"a", "b", "c", "FROM", "file.csv", "ORDER BY", "1", "2", "3"},
			IsMethod: true,
			Output:   "",
			FileName: "",
			Error:    true,
		},
		{
			Input:    []string{"a", "b", "c", "FROM", "file name with space.csv", "ORDER BY", "1", "2", "3"},
			IsMethod: true,
			Output:   "",
			FileName: "",
			Error:    true,
		},

		{
			Input:    []string{"a", "b", "c", "ORDER BY", "1", "2", "3"},
			IsMethod: true,
			Output:   "a b c ORDER BY 1 2 3",
			FileName: "",
		},

		// not a method

		{
			Input:    []string{"FROM", "file.csv", "ORDER BY", "1"},
			IsMethod: false,
			Output:   "* ORDER BY 1",
			FileName: "file.csv",
		},
		{
			Input:    []string{"*", "FROM", "file.csv", "ORDER BY", "1"},
			IsMethod: false,
			Output:   "* ORDER BY 1",
			FileName: "file.csv",
		},
		{
			Input:    []string{"*", "FROM", "file.csv"},
			IsMethod: false,
			Output:   "*",
			FileName: "file.csv",
		},
		{
			Input:    []string{"FROM", "file.csv"},
			IsMethod: false,
			Output:   "*",
			FileName: "file.csv",
		},
		{
			Input:    []string{"FROM", "file name with space.csv"},
			IsMethod: false,
			Output:   "*",
			FileName: "file name with space.csv",
		},
		{
			Input:    []string{"a", "b", "c", "FROM", "file.csv"},
			IsMethod: false,
			Output:   "a b c",
			FileName: "file.csv",
		},
		{
			Input:    []string{"FROM", "file.csv", "ORDER BY", "1", "2", "3"},
			IsMethod: false,
			Output:   "* ORDER BY 1 2 3",
			FileName: "file.csv",
		},
		{
			Input:    []string{"a", "b", "c", "FROM", "file.csv", "ORDER BY", "1", "2", "3"},
			IsMethod: false,
			Output:   "a b c ORDER BY 1 2 3",
			FileName: "file.csv",
		},
		{
			Input:    []string{"a", "b", "c", "FROM", "file name with space.csv", "ORDER BY", "1", "2", "3"},
			IsMethod: false,
			Output:   "a b c ORDER BY 1 2 3",
			FileName: "file name with space.csv",
		},

		{
			Input:      []string{"FROM", "<foo>,", "<bar>", "ORDER BY", "1"},
			IsMethod:   false,
			Output:     "* ORDER BY 1",
			NamedPipes: []string{"foo", "bar"},
		},
		{
			Input:      []string{"*", "FROM", "<foo>,", "<bar>", "ORDER BY", "1"},
			IsMethod:   false,
			Output:     "* ORDER BY 1",
			NamedPipes: []string{"foo", "bar"},
		},
		{
			Input:      []string{"*", "FROM", "<foo>,", "<bar>"},
			IsMethod:   false,
			Output:     "*",
			NamedPipes: []string{"foo", "bar"},
		},
		{
			Input:      []string{"FROM", "<foo>,", "<bar>"},
			IsMethod:   false,
			Output:     "*",
			NamedPipes: []string{"foo", "bar"},
		},
		{
			Input:      []string{"a", "b", "c", "FROM", "<foo>,", "<bar>"},
			IsMethod:   false,
			Output:     "a b c",
			NamedPipes: []string{"foo", "bar"},
		},
		{
			Input:      []string{"FROM", "<foo>,", "<bar>", "ORDER BY", "1", "2", "3"},
			IsMethod:   false,
			Output:     "* ORDER BY 1 2 3",
			NamedPipes: []string{"foo", "bar"},
		},
		{
			Input:      []string{"a", "b", "c", "FROM", "<fee>", "ORDER BY", "1", "2", "3"},
			IsMethod:   false,
			Output:     "a b c ORDER BY 1 2 3",
			NamedPipes: []string{"fee"},
		},
		{
			Input:      []string{"a", "b", "c", "FROM", "<fee>,", "<fii>", "ORDER BY", "1", "2", "3"},
			IsMethod:   false,
			Output:     "a b c ORDER BY 1 2 3",
			NamedPipes: []string{"fee", "fii"},
		},
		{
			Input:      []string{"a", "b", "c", "FROM", "<fee>,", "<fii>,", "<fo>", "ORDER BY", "1", "2", "3"},
			IsMethod:   false,
			Output:     "a b c ORDER BY 1 2 3",
			NamedPipes: []string{"fee", "fii", "fo"},
		},
		{
			Input:      []string{"a", "b", "c", "FROM", "<fee>,", "<fii>,", "<fo>,", "<fum>", "ORDER BY", "1", "2", "3"},
			IsMethod:   false,
			Output:     "a b c ORDER BY 1 2 3",
			NamedPipes: []string{"fee", "fii", "fo", "fum"},
		},

		{
			Input:     []string{"FROM", "$foo,", "$bar", "ORDER BY", "1"},
			IsMethod:  false,
			Output:    "* ORDER BY 1",
			Variables: []string{"foo", "bar"},
		},
		{
			Input:     []string{"*", "FROM", "$foo,", "$bar", "ORDER BY", "1"},
			IsMethod:  false,
			Output:    "* ORDER BY 1",
			Variables: []string{"foo", "bar"},
		},
		{
			Input:     []string{"*", "FROM", "$foo,", "$bar"},
			IsMethod:  false,
			Output:    "*",
			Variables: []string{"foo", "bar"},
		},
		{
			Input:     []string{"FROM", "$foo,", "$bar"},
			IsMethod:  false,
			Output:    "*",
			Variables: []string{"foo", "bar"},
		},
		{
			Input:     []string{"a", "b", "c", "FROM", "$foo,", "$bar"},
			IsMethod:  false,
			Output:    "a b c",
			Variables: []string{"foo", "bar"},
		},
		{
			Input:     []string{"FROM", "$foo,", "$bar", "ORDER BY", "1", "2", "3"},
			IsMethod:  false,
			Output:    "* ORDER BY 1 2 3",
			Variables: []string{"foo", "bar"},
		},
		{
			Input:     []string{"a", "b", "c", "FROM", "$fee", "ORDER BY", "1", "2", "3"},
			IsMethod:  false,
			Output:    "a b c ORDER BY 1 2 3",
			Variables: []string{"fee"},
		},
		{
			Input:     []string{"a", "b", "c", "FROM", "$fee,", "$fii", "ORDER BY", "1", "2", "3"},
			IsMethod:  false,
			Output:    "a b c ORDER BY 1 2 3",
			Variables: []string{"fee", "fii"},
		},
		{
			Input:     []string{"a", "b", "c", "FROM", "$fee,", "$fii,", "$fo", "ORDER BY", "1", "2", "3"},
			IsMethod:  false,
			Output:    "a b c ORDER BY 1 2 3",
			Variables: []string{"fee", "fii", "fo"},
		},
		{
			Input:     []string{"a", "b", "c", "FROM", "$fee,", "$fii,", "$fo,", "$fum", "ORDER BY", "1", "2", "3"},
			IsMethod:  false,
			Output:    "a b c ORDER BY 1 2 3",
			Variables: []string{"fee", "fii", "fo", "fum"},
		},

		{
			Input:    []string{"a", "b", "c", "ORDER BY", "1", "2", "3"},
			IsMethod: false,
			Output:   "",
			FileName: "",
			Error:    true,
		},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		p := lang.NewTestProcess()
		p.IsMethod = test.IsMethod
		p.Parameters.DefineParsed(test.Input)
		actOutput, actFileName, actPipes, actVars, err := dissectParameters(p)

		if actOutput != test.Output {
			t.Errorf("Parameter output does not match expected in test %d", i)
			t.Logf("  Input:      %v", inlineJson(test.Input))
			t.Logf("  IsMethod:   %v", test.IsMethod)
			t.Logf("  exp param: '%s'", test.Output)
			t.Logf("  act param: '%s'", actOutput)
			t.Logf("  exp file:  '%s'", test.FileName)
			t.Logf("  act file:  '%s'", actFileName)
			t.Logf("  exp pipes:  %s", inlineJson(test.NamedPipes))
			t.Logf("  act pipes:  %s", inlineJson(actPipes))
			t.Logf("  exp vars:   %s", inlineJson(test.Variables))
			t.Logf("  act vars:   %s", inlineJson(actVars))
			t.Logf("  exp error:  %v", test.Error)
			t.Logf("  act error:  %v", err)
		}

		if actFileName != test.FileName {
			t.Errorf("FileName output does not match expected in test %d", i)
			t.Logf("  Input:      %v", inlineJson(test.Input))
			t.Logf("  IsMethod:   %v", test.IsMethod)
			t.Logf("  exp param: '%s'", test.Output)
			t.Logf("  act param: '%s'", actOutput)
			t.Logf("  exp file:  '%s'", test.FileName)
			t.Logf("  act file:  '%s'", actFileName)
			t.Logf("  exp pipes:  %s", inlineJson(test.NamedPipes))
			t.Logf("  act pipes:  %s", inlineJson(actPipes))
			t.Logf("  exp vars:   %s", inlineJson(test.Variables))
			t.Logf("  act vars:   %s", inlineJson(actVars))
			t.Logf("  exp error:  %v", test.Error)
			t.Logf("  act error:  %v", err)
		}

		if (err != nil) != test.Error {
			t.Errorf("Output does not match expected in test %d", i)
			t.Logf("  Input:      %v", inlineJson(test.Input))
			t.Logf("  IsMethod:   %v", test.IsMethod)
			t.Logf("  exp param: '%s'", test.Output)
			t.Logf("  act param: '%s'", actOutput)
			t.Logf("  exp file:  '%s'", test.FileName)
			t.Logf("  act file:  '%s'", actFileName)
			t.Logf("  exp pipes:  %s", inlineJson(test.NamedPipes))
			t.Logf("  act pipes:  %s", inlineJson(actPipes))
			t.Logf("  exp vars:   %s", inlineJson(test.Variables))
			t.Logf("  act vars:   %s", inlineJson(actVars))
			t.Logf("  exp error:  %v", test.Error)
			t.Logf("  act error:  %v", err)
		}

		if inlineJson(actPipes) != inlineJson(test.NamedPipes) {
			t.Errorf("Pipes do not match expected in test %d", i)
			t.Logf("  Input:      %v", inlineJson(test.Input))
			t.Logf("  IsMethod:   %v", test.IsMethod)
			t.Logf("  exp param: '%s'", test.Output)
			t.Logf("  act param: '%s'", actOutput)
			t.Logf("  exp file:  '%s'", test.FileName)
			t.Logf("  act file:  '%s'", actFileName)
			t.Logf("  exp pipes:  %s", inlineJson(test.NamedPipes))
			t.Logf("  act pipes:  %s", inlineJson(actPipes))
			t.Logf("  exp vars:   %s", inlineJson(test.Variables))
			t.Logf("  act vars:   %s", inlineJson(actVars))
			t.Logf("  exp error:  %v", test.Error)
			t.Logf("  act error:  %v", err)
		}

		if inlineJson(actVars) != inlineJson(test.Variables) {
			t.Errorf("Variables do not match expected in test %d", i)
			t.Logf("  Input:      %v", inlineJson(test.Input))
			t.Logf("  IsMethod:   %v", test.IsMethod)
			t.Logf("  exp param: '%s'", test.Output)
			t.Logf("  act param: '%s'", actOutput)
			t.Logf("  exp file:  '%s'", test.FileName)
			t.Logf("  act file:  '%s'", actFileName)
			t.Logf("  exp pipes:  %s", inlineJson(test.NamedPipes))
			t.Logf("  act pipes:  %s", inlineJson(actPipes))
			t.Logf("  exp vars:   %s", inlineJson(test.Variables))
			t.Logf("  act vars:   %s", inlineJson(actVars))
			t.Logf("  exp error:  %v", test.Error)
			t.Logf("  act error:  %v", err)
		}
	}
}
