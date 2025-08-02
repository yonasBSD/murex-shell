package streams

import (
	"context"
	"io"
	"os"

	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang/stdio"
)

/*func init() {
	stdio.RegisterPipe("tee", func(string) (stdio.Io, error) {
		return nil, errors.New("`tee` is a system device used for `test`. It's user creation isn't yet supported but might be included in a future release")
	})
}*/

// Tee is a stream interface with two output streams
// (like the `tee` command on UNIX/Linux)
type Tee struct {
	primary   stdio.Io
	secondary Stdin
}

// NewTee creates a new tee stdio interface
func NewTee(primary stdio.Io) (primaryTee *Tee, secondaryTee *Stdin) {
	primaryTee = new(Tee)
	primaryTee.primary = primary
	primaryTee.secondary.max = 0
	primaryTee.secondary.ctx = context.Background()
	secondaryTee = &primaryTee.secondary
	return
}

// IsTTY calls the primary STDOUT stream in tee to see if it's a TTY
func (tee *Tee) IsTTY() bool {
	return tee.primary.IsTTY()
}

func (tee *Tee) File() *os.File {
	return tee.primary.File()
}

// Stats is stored against the primary STDOUT stream in tee
func (tee *Tee) Stats() (uint64, uint64) {
	return tee.primary.Stats()
}

// Read from STDIN (uses primary tee stream)
func (tee *Tee) Read(p []byte) (int, error) {
	return tee.primary.Read(p)
}

// ReadLine reads a line from STDIN (uses the primary tee stream)
func (tee *Tee) ReadLine(callback func([]byte)) error {
	return tee.primary.ReadLine(callback)
}

// ReadArray reads an array from STDIN (uses the primary tee stream)
func (tee *Tee) ReadArray(ctx context.Context, callback func([]byte)) error {
	return tee.primary.ReadArray(ctx, callback)
}

// ReadArrayWithType reads an array from STDIN (uses the primary tee stream)
func (tee *Tee) ReadArrayWithType(ctx context.Context, callback func(any, string)) error {
	return tee.primary.ReadArrayWithType(ctx, callback)
}

// ReadMap reads a hash table from STDIN (uses the primary tee stream)
func (tee *Tee) ReadMap(config *config.Config, callback func(*stdio.Map)) error {
	return tee.primary.ReadMap(config, callback)
}

// ReadAll from STDIN (uses the primary tee stream)
func (tee *Tee) ReadAll() ([]byte, error) {
	return tee.primary.ReadAll()
}

// Write is the standard Writer interface Write() method.
func (tee *Tee) Write(p []byte) (int, error) {
	tee.secondary.Write(p)
	return tee.primary.Write(p)
}

// Writeln just calls Write() but with an appended, OS specific, new line.
func (tee *Tee) Writeln(p []byte) (int, error) {
	tee.secondary.Writeln(p)
	return tee.primary.Writeln(p)
}

// WriteArray performs data type specific buffered writes to an stdio.Io interface
func (tee *Tee) WriteArray(dataType string) (stdio.ArrayWriter, error) {
	return stdio.WriteArray(tee, dataType)
}

// Open the stream.Io interface for another dependant
func (tee *Tee) Open() {
	tee.primary.Open()
}

// Close the stream.Io interface
func (tee *Tee) Close() {
	tee.primary.Close()
}

// ForceClose forces the stream.Io interface to close. This should only be called by a STDIN reader
func (tee *Tee) ForceClose() {
	tee.primary.ForceClose()
}

// WriteTo reads from the stream.Io interface and writes to a destination
// io.Writer interface
func (tee *Tee) WriteTo(w io.Writer) (n int64, err error) {
	return tee.primary.WriteTo(w)
}

// GetDataType returns the murex data type for the stream.Io interface
func (tee *Tee) GetDataType() (dt string) {
	return tee.primary.GetDataType()
}

// SetDataType defines the murex data type for the stream.Io interface
func (tee *Tee) SetDataType(dt string) {
	tee.secondary.SetDataType(dt)
	tee.primary.SetDataType(dt)
}
