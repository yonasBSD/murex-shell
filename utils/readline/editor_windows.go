//go:build windows
// +build windows

package readline

import "errors"

func (rl *Instance) launchEditor(multiline []rune) ([]rune, error) {
	return rl.line.Runes(), errors.New("Not currently supported on Windows")
}
