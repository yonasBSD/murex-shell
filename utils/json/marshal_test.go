package json_test

import (
	"testing"

	"github.com/lmorg/murex/test/count"
	"github.com/lmorg/murex/utils/json"
)

// TestJsonMap tests the the JSON wrapper can marshal any maps which the
// core library cannot
func TestJsonMap(t *testing.T) {
	count.Tests(t, 1)

	obj := make(map[any]any)
	obj["a"] = "b"
	obj[1] = 2

	b, err := json.Marshal(obj, false)
	if err != nil {
		t.Error("Error marshalling: " + err.Error())
	}

	if string(b) != `{"a":"b","1":2}` && string(b) != `{"1":2,"a":"b"}` {
		t.Error("Unexpected JSON returned: " + string(b))
	}
}
