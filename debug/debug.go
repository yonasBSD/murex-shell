package debug

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Enabled is a flag used for debugging murex code. This can be enabled at
// startup by a `--debug` flag or during runtime with `debug on`.
var Enabled bool

// Log writes a debug message
func Log(data ...any) {
	if Enabled {
		log.Println(data...)
	}
}

// Logf writes a debug message using [fmt.Printf] arguments.
func Logf(format string, v ...any) {
	if Enabled {
		log.Println(fmt.Sprintf(format, v...))
	}
}

// Json encode an object then write it as a debug message
func Json(context string, data any) {
	if Enabled {
		b, _ := json.MarshalIndent(data, "", "\t")
		Log(context, "JSON:"+string(b))
	}
}

// Dump is used for runtime output of the status of various debug modes
func Dump() any {
	type status struct {
		Debug bool
	}

	return status{
		Debug: Enabled,
	}
}

func LogWriter(path string) error {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	log.SetOutput(f)
	log.SetPrefix(fmt.Sprintf("[PID: %d] ", os.Getpid()))
	return nil
}
