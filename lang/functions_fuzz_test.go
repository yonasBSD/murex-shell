package lang_test

import (
	"testing"

	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/test/count"
)

func FuzzFuncParseDataTypes(f *testing.F) {
	tests := []string{"name: str, age: int", "", "!12345", `age: int "how old are you?" [123]`}
	for _, tc := range tests {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, orig string) {
		count.Tests(t, 1)
		lang.ParseMxFunctionParameters(orig)
		// we are just testing we can't cause an unhandled panic
	})
}
