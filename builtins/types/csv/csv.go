package csv

import (
	"github.com/lmorg/murex/config"
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/stdio"
	"github.com/lmorg/murex/lang/types"
)

const typeName = "csv"

func init() {
	//stdio.RegisterReadArray(typeName, readArray)
	stdio.RegisterReadMap(typeName, readMap)

	lang.ReadIndexes[typeName] = readIndex
	lang.ReadNotIndexes[typeName] = readIndex

	lang.RegisterMarshaller(typeName, marshal)
	lang.RegisterUnmarshaller(typeName, unmarshal)

	// `application/csv` and `text/csv` are the common ones. `x-csv` is added just in case anyone decides to use
	// something non-standard.
	lang.SetMime(typeName,
		"application/csv",
		"application/x-csv",
		"text/csv",
		"text/x-csv",
		"+csv",
	)

	lang.SetFileExtensions(typeName, "csv")

	config.InitConf.Define("csv", "separator", config.Properties{
		Description: "The delimiter for records in a CSV file.",
		Default:     `,`,
		DataType:    types.String,
	})

	config.InitConf.Define("csv", "comment", config.Properties{
		Description: "The prefix token for comments in a CSV table.",
		Default:     `#`,
		DataType:    types.String,
	})

	config.InitConf.Define("csv", "left-trim", config.Properties{
		Description: "Trim leading whitespace if input is generic or string, and when heading contains whitespace",
		Default:     true,
		DataType:    types.Boolean,
	})
}
