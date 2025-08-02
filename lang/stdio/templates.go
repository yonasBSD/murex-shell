package stdio

// template functions for stdio.Io methods to call
// (saves reinventing the wheel lots of times)

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/types"
)

// ReadArray is a template function for stdio.Io
func ReadArray(ctx context.Context, read Io, callback func([]byte)) error {
	dt := read.GetDataType()

	fnReadArray := readArray[dt]
	if fnReadArray != nil {
		return fnReadArray(ctx, read, callback)
	}

	return readArray[types.Generic](ctx, read, callback)
}

// ReadArrayWithType is a template function for stdio.Io
func ReadArrayWithType(ctx context.Context, read Io, callback func(any, string)) error {
	dt := read.GetDataType()

	fnReadArray := readArrayWithType[dt]
	if fnReadArray != nil {
		return fnReadArray(ctx, read, callback)
	}

	return readArrayWithType[types.Generic](ctx, read, callback)
}

// ReadMap is a template function for stdio.Io
func ReadMap(read Io, config *config.Config, callback func(*Map)) error {
	dataType := read.GetDataType()

	fnReadMap := readMap[dataType]
	if fnReadMap != nil {
		return fnReadMap(read, config, callback)
	}

	return readMap[types.Generic](read, config, callback)
}

// WriteArray is a template function for stdio.Io
func WriteArray(writer Io, dt string) (ArrayWriter, error) {
	if writeArray[dt] != nil {
		return writeArray[dt](writer)
	}

	return nil, fmt.Errorf("murex data type `%s` has not implemented WriteArray() method", dt)
}

// WriteTo is a template function for stdio.Io
func WriteTo(std Io, w io.Writer) (int64, error) {
	var (
		total int64
		i, n  int
		p     = make([]byte, 1024*10)
		err   error
	)

	for {
		i, err = std.Read(p)

		if err == io.EOF {
			return total, nil
		}

		if err != nil {
			return total, err
		}

		n, err = w.Write(p[:i])
		total += int64(n)

		if err != nil {
			return total, err
		}

	}
}

// WriteToFromFile is a template function for stdio.Io
func WriteToFromFile(f *os.File, w io.Writer) (int64, error) {
	var (
		total int64
		i, n  int
		p     = make([]byte, 1024*10)
		err   error
	)

	for {
		i, err = f.Read(p)

		if err == io.EOF {
			return total, nil
		}

		if err != nil {
			return total, err
		}

		n, err = w.Write(p[:i])
		total += int64(n)

		if err != nil {
			return total, err
		}

	}
}
