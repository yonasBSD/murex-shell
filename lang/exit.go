package lang

import (
	"os"
)

var (
	ProfCpuCleanUp   func() = func() {}
	ProfMemCleanUp   func() = func() {}
	ProfTraceCleanUp func() = func() {}
)

func Exit(exitNum int) {
	ProfCpuCleanUp()
	ProfMemCleanUp()
	ProfTraceCleanUp()

	os.Exit(exitNum)
}
