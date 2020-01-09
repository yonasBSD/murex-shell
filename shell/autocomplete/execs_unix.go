// +build !windows,!plan9

package autocomplete

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/lmorg/murex/utils/consts"
	"github.com/phayes/permbits"
)

// SplitPath takes a $PATH (or similar) string and splits it into an array of paths
func SplitPath(envPath string) []string {
	split := strings.Split(envPath, ":")
	return split
}

func listExes(path string, exes map[string]bool) {
	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		perm := permbits.FileMode(f.Mode())
		switch {
		case perm.OtherExecute() && f.Mode().IsRegular():
			exes[f.Name()] = true
		case perm.OtherExecute() && f.Mode()&os.ModeSymlink != 0:
			ln, err := os.Readlink(path + consts.PathSlash + f.Name())
			if err != nil {
				continue
			}
			if ln[0] != consts.PathSlash[0] {
				ln = path + consts.PathSlash + ln
			}
			info, err := os.Stat(ln)
			if err != nil {
				continue
			}
			perm := permbits.FileMode(info.Mode())
			if perm.OtherExecute() && info.Mode().IsRegular() {
				exes[f.Name()] = true
			}

		default:
			/*|| perm.GroupExecute()||perm.UserExecute() need to check what user and groups you are in first */
		}
	}
}

func matchExes(s string, exes map[string]bool, includeColon bool) (items []string) {
	colon := ""
	// We only want a colon added if the exe is the function call rather than a
	// functions parameter (eg `some-exec` vs `sudo some-exec`).
	if includeColon {
		colon = ":"
	}

	for name := range exes {
		if strings.HasPrefix(name, s) {
			if name != consts.NamedPipeProcName {
				if !isSpecialBuiltin(name) {
					name = name + colon
				}
				items = append(items, name[len(s):])
			}
		}
	}

	sortColon(items, 0, len(items)-1)

	// I know it seems weird and inefficient to cycle through the array after
	// it has been created just to append a couple of characters (that easily
	// could have been appended in the former for loop) but this is so that the
	// colon isn't included as part of the sorting algorithm (eg otherwise
	// `manpath:` would precede `man:`). Ideally I would write my own sorting
	// function to take this into account but that can be part of the
	// optimisation stage - whenever I get there.
	/*for i := range items {
		if !isSpecialBuiltin(items[i]) {
			items[i] += colon
		}
	}*/
	return
}
