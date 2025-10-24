package userdictionary

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/lmorg/murex/app"
	"github.com/lmorg/murex/lang/types"
)

var dictionary = []string{
	app.Name,
}

// IsInDictionary checks if word is in user dictionary
func IsInDictionary(word string) bool {
	s := strings.ToLower(word)
	for i := range dictionary {
		if s == strings.ToLower(dictionary[i]) {
			return true
		}
	}

	return false
}

// Get returns a copy of the slice dictionary
func Get() []string {
	a := make([]string, len(dictionary))
	copy(a, dictionary)
	return a
}

// Read returns an any of the user dictionary.
// This is only intended to be used by `config.Properties.GoFunc.Read()`
func Read() (any, error) {
	return Get(), nil
}

// Write takes a JSON-encoded string and writes it to the dictionary slice.
// This is only intended to be used by `config.Properties.GoFunc.Write()`
func Write(v any) error {
	switch v := v.(type) {
	case string:
		return json.Unmarshal([]byte(v), &dictionary)

	default:
		return fmt.Errorf("invalid data-type. Expecting a %s encoded string", types.Json)
	}
}
