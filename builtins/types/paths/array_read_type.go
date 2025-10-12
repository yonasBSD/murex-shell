package paths

import (
	"bytes"
	"context"

	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
	"github.com/lmorg/murex/utils/path"
)

func readArrayWithTypePath(ctx context.Context, read stdio.Io, callback func(any, string)) error {
	b, err := read.ReadAll()
	if err != nil {
		return err
	}

	s := string(b)
	split := path.Split(s)

	for i := range split {
		select {
		case <-ctx.Done():
			return nil
		default:
			callback(split[i], types.String)
		}
	}

	return nil
}

func readArrayWithTypePaths(ctx context.Context, read stdio.Io, callback func(any, string)) error {
	b, err := read.ReadAll()
	if err != nil {
		return err
	}

	split := bytes.Split(b, pathsSeparator)
	for i := range split {

		select {
		case <-ctx.Done():
			return nil

		default:
			callback(string(split[i]), types.Path)
		}

	}

	return nil
}
