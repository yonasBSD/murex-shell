package lang

import (
	"testing"

	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/json"
)

type testFuncParseDataTypesT struct {
	Parameters string
	Error      bool
	Expected   []MurexFuncParam
}

func testFuncParseDataTypes(t *testing.T, tests []testFuncParseDataTypesT) {
	t.Helper()
	count.Tests(t, len(tests))

	for i := range tests {
		actual, err := ParseMxFunctionParameters(tests[i].Parameters)
		if (err == nil) == tests[i].Error {
			t.Errorf("Unexpected error raised in test %d", i)
			t.Logf("Parameters: %s", tests[i].Parameters)
			t.Logf("Expected:   %s", json.LazyLogging(tests[i].Expected))
			t.Logf("Actual:     %s", json.LazyLogging(actual))
			t.Logf("exp err:    %v", tests[i].Error)
			t.Logf("act err:    %s", err)
		}

		if json.LazyLogging(tests[i].Expected) != json.LazyLogging(actual) {
			t.Errorf("Unexpected error raised in test %d", i)
			t.Logf("Parameters: %s", tests[i].Parameters)
			t.Logf("Expected:   %s", json.LazyLogging(tests[i].Expected))
			t.Logf("Actual:     %s", json.LazyLogging(actual))
			t.Logf("exp err:    %v", tests[i].Error)
			t.Logf("act err:    %s", err)
		}
	}
}

func TestFuncParseDataTypes(t *testing.T) {
	tests := []testFuncParseDataTypesT{
		{
			Parameters: `name, age`,
			Expected: []MurexFuncParam{{
				Name:     "name",
				DataType: "str",
			}, {
				Name:     "age",
				DataType: "str",
			}},
		},
		{
			Parameters: `name: str, age: int`,
			Expected: []MurexFuncParam{{
				Name:     "name",
				DataType: "str",
			}, {
				Name:     "age",
				DataType: "int",
			}},
		},
		{
			Parameters: `name: str "What is your name?", age: int "How old are you?"`,
			Expected: []MurexFuncParam{{
				Name:        "name",
				DataType:    "str",
				Description: "What is your name?",
			}, {
				Name:        "age",
				DataType:    "int",
				Description: "How old are you?",
			}},
		},
		{
			Parameters: `name: str [Bob], age: int [100]`,
			Expected: []MurexFuncParam{{
				Name:       "name",
				DataType:   "str",
				Default:    "Bob",
				HasDefault: true,
			}, {
				Name:       "age",
				DataType:   "int",
				Default:    "100",
				HasDefault: true,
			}},
		},
		{
			Parameters: `name: str "What is your name?" [Bob], age: int "How old are you?" [100]`,
			Expected: []MurexFuncParam{{
				Name:        "name",
				DataType:    "str",
				Description: "What is your name?",
				Default:     "Bob",
				HasDefault:  true,
			}, {
				Name:        "age",
				DataType:    "int",
				Description: "How old are you?",
				Default:     "100",
				HasDefault:  true,
			}},
		},
		{
			Parameters: `name: str [Bob]`,
			Expected: []MurexFuncParam{{
				Name:       "name",
				DataType:   "str",
				Default:    "Bob",
				HasDefault: true,
			}},
		},
		{
			Parameters: `colon: str ":" [:]`,
			Expected: []MurexFuncParam{{
				Name:        "colon",
				DataType:    "str",
				Description: ":",
				Default:     ":",
				HasDefault:  true,
			}},
		},
		{
			Parameters: `quote: str "'" ["]`,
			Expected: []MurexFuncParam{{
				Name:        "quote",
				DataType:    "str",
				Description: "'",
				Default:     "\"",
				HasDefault:  true,
			}},
		},
		{
			Parameters: `square: str "[" [[]`,
			Expected: []MurexFuncParam{{
				Name:        "square",
				DataType:    "str",
				Description: "[",
				Default:     "[",
				HasDefault:  true,
			}},
		},
		{
			Parameters: `square: str "]" [square]`,
			Expected: []MurexFuncParam{{
				Name:        "square",
				DataType:    "str",
				Description: "]",
				Default:     "square",
				HasDefault:  true,
			}},
		},
		{
			Parameters: `comma: str "," [,]`,
			Expected: []MurexFuncParam{{
				Name:        "comma",
				DataType:    "str",
				Description: ",",
				Default:     ",",
				HasDefault:  true,
			}},
		},
	}

	testFuncParseDataTypes(t, tests)
}

func TestFuncParseDataTypesErrorOutOfOrder(t *testing.T) {
	tests := []testFuncParseDataTypesT{
		{
			Parameters: `foo baz [bar]`,
			Error:      true,
			/*Expected: []MurexFuncParam{{
				Name:     "foo",
				DataType: "baz",
				Default:  "bar",
			}},*/
		},
		{
			Parameters: `foo "baz" [bar]`,
			Error:      true,
			/*Expected: []MurexFuncParam{{
				Name:        "foo",
				Description: "baz",
				Default:     "bar",
			}},*/
		},
		{
			Parameters: `foo [bar]`,
			Error:      true,
			/*Expected: []MurexFuncParam{{
				Name:     "foo",
				DataType: "str",
				Default:  "bar",
			}},*/
		},
	}

	testFuncParseDataTypes(t, tests)
}

func TestFuncParseDataTypesErrorCrLf(t *testing.T) {
	tests := []testFuncParseDataTypesT{
		{
			Parameters: `\nname: str, age: int`,
			Error:      true,
		},
		{
			Parameters: `\r\nname: str, age: int`,
			Error:      true,
		},

		{
			Parameters: `na\nme: str, age: int`,
			Error:      true,
		},
		{
			Parameters: `na\r\nme: str, age: int`,
			Error:      true,
		},

		{
			Parameters: `name\n: str, age: int`,
			Error:      true,
		},
		{
			Parameters: `name\r\n: str, age: int`,
			Error:      true,
		},

		{
			Parameters: `name: \nstr, age: int`,
			Error:      true,
		},
		{
			Parameters: `name: \r\nstr, age: int`,
			Error:      true,
		},

		{
			Parameters: `name: str, age: int\nname: str, age: int`,
			Error:      true,
		},
		{
			Parameters: `name: str, age: int\r\nname: str, age: int`,
			Error:      true,
		},

		{
			Parameters: `name: str, age: int\nname\n: str, age: int`,
			Error:      true,
		},
		{
			Parameters: `name: str, age: int\r\nname\r\n str, age: int`,
			Error:      true,
		},

		{
			Parameters: `name: str, age: int\r\nname: \nstr, age: int`,
			Error:      true,
		},
		{
			Parameters: `name: str, age: int\r\nname: \r\nstr, age: int`,
			Error:      true,
		},
	}

	testFuncParseDataTypes(t, tests)
}

func TestFuncParseDataTypesErrorSpaceTab(t *testing.T) {
	tests := []testFuncParseDataTypesT{
		{
			Parameters: `name : str, age: int`,
			Error:      true,
		},
		{
			Parameters: `name\t: str, age: int`,
			Error:      true,
		},
	}

	testFuncParseDataTypes(t, tests)
}

func TestFuncParseDataTypesErrorColon(t *testing.T) {
	tests := []testFuncParseDataTypesT{
		{
			Parameters: `name: str, age: int :`,
			Error:      true,
		},
	}

	testFuncParseDataTypes(t, tests)
}

func TestFuncParseDataTypesErrorQuote(t *testing.T) {
	tests := []testFuncParseDataTypesT{
		{
			Parameters: `"name": str`,
			Error:      true,
		},
		{
			Parameters: `name: "str"`,
			Error:      true,
		},
	}

	testFuncParseDataTypes(t, tests)
}

func TestFuncParseDataTypesErrorSquare(t *testing.T) {
	tests := []testFuncParseDataTypesT{
		{
			Parameters: `]: str`,
			Error:      true,
		},
		{
			Parameters: `name: ]`,
			Error:      true,
		},
		{
			Parameters: `name: str ]`,
			Error:      true,
		},
	}

	testFuncParseDataTypes(t, tests)
}

func TestFuncParseDataTypesErrorComma(t *testing.T) {
	tests := []testFuncParseDataTypesT{
		{
			Parameters: `,name`,
			Error:      true,
		},
		{
			Parameters: `name: ,`,
			Error:      true,
		},
		{
			Parameters: `name,,`,
			Error:      true,
		},
	}

	testFuncParseDataTypes(t, tests)
}

func TestFuncSummary(t *testing.T) {
	tests := []struct {
		Block   string
		Summary string
	}{
		{
			Block:   "",
			Summary: "",
		},
		{
			Block:   "hello world",
			Summary: "",
		},
		{
			Block:   "foo\nbar",
			Summary: "",
		},
		{
			Block:   "foo\n#bar",
			Summary: "",
		},
		{
			Block:   "{\n#bar",
			Summary: "bar",
		},
		{
			Block:   "foo\n\n#bar",
			Summary: "",
		},
		{
			Block:   "{\n\n#bar",
			Summary: "",
		},
		{
			Block:   "foo\nbar\nbaz",
			Summary: "",
		},
		{
			Block:   "{\nbar\nbaz",
			Summary: "",
		},
		{
			Block:   "foo\nbar\n#baz",
			Summary: "",
		},
		{
			Block:   "{\nbar\n#baz",
			Summary: "",
		},
		{
			Block:   "foo\n#bar\n#baz",
			Summary: "",
		},
		{
			Block:   "{\n#bar\n#baz",
			Summary: "bar",
		},
		{
			Block:   "\n#bar\n#baz",
			Summary: "bar",
		},
		{
			Block:   "\n\n#bar\n#baz",
			Summary: "",
		},
		{
			Block:   "foo\n# bar\n# baz",
			Summary: "",
		},
		{
			Block:   "{\n# bar\n# baz",
			Summary: "bar",
		},
		{
			Block:   "foo\n#  bar\n#  baz",
			Summary: "",
		},
		{
			Block:   "{\n#  bar\n#  baz",
			Summary: "bar",
		},
		{
			Block:   "foo\n#\tbar\n#\tbaz",
			Summary: "",
		},
		{
			Block:   "{\n#\tbar\n#\tbaz",
			Summary: "bar",
		},
		{
			Block:   "#\tbar",
			Summary: "bar",
		},
		{
			Block:   "# bar\t",
			Summary: "bar",
		},
		{
			Block:   "# foo bar",
			Summary: "foo bar",
		},
		{
			Block:   "# foo\tbar",
			Summary: "foo    bar",
		},
		{
			Block:   " # foo\tbar",
			Summary: "foo    bar",
		},
		{
			Block:   "\t# foo\tbar",
			Summary: "foo    bar",
		},
		{
			Block:   "baz # foo\tbar",
			Summary: "",
		},
		{
			Block:   "{ # foo\tbar",
			Summary: "foo    bar",
		},
		{
			Block:   "baz\n # foo\tbar",
			Summary: "",
		},
		{
			Block:   "{\n # foo\tbar",
			Summary: "foo    bar",
		},
		{
			Block:   "1\n2\n # foo\tbar",
			Summary: "",
		},
		{
			Block:   "{\n # foo\tbar\n}",
			Summary: "foo    bar",
		},
		{
			Block:   "\r\n",
			Summary: "",
		},
		{
			Block:   "\r\n# foo\r\nbar",
			Summary: "foo",
		},
	}

	count.Tests(t, len(tests))

	for i, test := range tests {
		actual := funcSummary([]rune(test.Block))
		if actual != test.Summary {
			t.Errorf("Summary mismatch in test %d", i)
			t.Logf("  Block:\n'%s'", test.Block)
			t.Logf("  Expected: '%s'", test.Summary)
			t.Logf("  Actual:   '%s'", actual)
		}
	}
}

func TestFuncParseDocExamples(t *testing.T) {
	tests := []testFuncParseDataTypesT{
		{
			Parameters: `var:datatype [default-value] "description"`,
			Expected: []MurexFuncParam{
				{
					Name:        "var",
					DataType:    "datatype",
					Default:     "default-value",
					Description: "description",
					HasDefault:  true,
				},
			},
		},
		{
			Parameters: `
				variable1: data-type [default-value] "description",
				variable2: data-type [default-value] "description"
			`,
			Expected: []MurexFuncParam{
				{
					Name:        "variable1",
					DataType:    "data-type",
					Default:     "default-value",
					Description: "description",
					HasDefault:  true,
				},
				{
					Name:        "variable2",
					DataType:    "data-type",
					Default:     "default-value",
					Description: "description",
					HasDefault:  true,
				},
			},
		},
	}

	testFuncParseDataTypes(t, tests)
}

func TestFuncParseOptional(t *testing.T) {
	tests := []testFuncParseDataTypesT{
		{
			Parameters: `!var:datatype [default-value] "description"`,
			Expected: []MurexFuncParam{
				{
					Name:        "var",
					DataType:    "datatype",
					Default:     "default-value",
					Description: "description",
					HasDefault:  true,
					Optional:    true,
				},
			},
		},
		{
			Parameters: `
				!variable1: data-type [default-value] "description",
				!variable2: data-type [default-value] "description"
			`,
			Expected: []MurexFuncParam{
				{
					Name:        "variable1",
					DataType:    "data-type",
					Default:     "default-value",
					Description: "description",
					HasDefault:  true,
					Optional:    true,
				},
				{
					Name:        "variable2",
					DataType:    "data-type",
					Default:     "default-value",
					Description: "description",
					HasDefault:  true,
					Optional:    true,
				},
			},
		},
		{
			Parameters: `
				variable1: data-type [default-value] "description",
				!variable2: data-type [default-value] "description"
			`,
			Expected: []MurexFuncParam{
				{
					Name:        "variable1",
					DataType:    "data-type",
					Default:     "default-value",
					Description: "description",
					HasDefault:  true,
				},
				{
					Name:        "variable2",
					DataType:    "data-type",
					Default:     "default-value",
					Description: "description",
					HasDefault:  true,
					Optional:    true,
				},
			},
		},
		{
			Parameters: `
				!variable1: data-type [default-value] "description",
				variable2: data-type [default-value] "description"
			`,
			Error: true,
		},
		{
			Parameters: `var: !data-type [default-value] "description"`,
			Error:      true,
		},
		{
			Parameters: `var: data-!type [default-value] "description"`,
			Error:      true,
		},
		{
			Parameters: `var: data-type [!default-value] "description"`,
			Expected: []MurexFuncParam{
				{
					Name:        "var",
					DataType:    "data-type",
					Default:     "!default-value",
					Description: "description",
					HasDefault:  true,
				},
			},
		},
		{
			Parameters: `var: data-type [default-value] "!description"`,
			Expected: []MurexFuncParam{
				{
					Name:        "var",
					DataType:    "data-type",
					Default:     "default-value",
					Description: "!description",
					HasDefault:  true,
				},
			},
		},
	}
	testFuncParseDataTypes(t, tests)
}

func TestFuncParseHasDefault(t *testing.T) {
	tests := []testFuncParseDataTypesT{
		{
			Parameters: `!var: data-type "description"`,
			Expected: []MurexFuncParam{
				{
					Name:        "var",
					DataType:    "data-type",
					Default:     "",
					Description: "description",
					HasDefault:  false,
					Optional:    true,
				},
			},
		},
		{
			Parameters: `!var: data-type`,
			Expected: []MurexFuncParam{
				{
					Name:        "var",
					DataType:    "data-type",
					Default:     "",
					Description: "",
					HasDefault:  false,
					Optional:    true,
				},
			},
		},
		/////
		{
			Parameters: `!var: data-type "description" []`,
			Expected: []MurexFuncParam{
				{
					Name:        "var",
					DataType:    "data-type",
					Default:     "",
					Description: "description",
					HasDefault:  true,
					Optional:    true,
				},
			},
		},
		{
			Parameters: `!var: data-type [] "description"`,
			Expected: []MurexFuncParam{
				{
					Name:        "var",
					DataType:    "data-type",
					Default:     "",
					Description: "description",
					HasDefault:  true,
					Optional:    true,
				},
			},
		},
		{
			Parameters: `!var: data-type []`,
			Expected: []MurexFuncParam{
				{
					Name:        "var",
					DataType:    "data-type",
					Default:     "",
					Description: "",
					HasDefault:  true,
					Optional:    true,
				},
			},
		},
	}

	testFuncParseDataTypes(t, tests)
}
