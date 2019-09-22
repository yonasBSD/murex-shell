package jsonlines

import (
	"github.com/lmorg/murex/lang"
	"github.com/lmorg/murex/lang/proc/stdio"
	"github.com/lmorg/murex/lang/types"
)

func init() {
	// Register data type
	lang.Marshallers[types.JsonLines] = marshal
	lang.Unmarshallers[types.JsonLines] = unmarshal
	lang.ReadIndexes[types.JsonLines] = index
	lang.ReadNotIndexes[types.JsonLines] = index

	stdio.RegesterReadArray(types.JsonLines, readArray)
	//stdio.RegesterReadMap(name, readMap)
	stdio.RegesterWriteArray(types.JsonLines, newArrayWriter)

	lang.SetMime(types.JsonLines,
		"application/jsonl",
		"application/x-jsonl",
		"text/jsonl",
		"text/x-jsonl",

		"application/jsonlines",
		"application/x-jsonlines",
		"text/jsonlines",
		"text/x-jsonlines",

		"application/json-lines",
		"application/x-json-lines",
		"text/json-lines",
		"text/x-json-lines",

		"application/jsonseq",
		"application/x-jsonseq",
		"text/jsonseq",
		"text/x-jsonseq",

		"application/json-seq",
		"application/x-json-seq",
		"text/json-seq",
		"text/x-json-seq",

		"application/ldjson",
		"application/x-ldjson",
		"text/ldjson",
		"text/x-ldjson",

		"application/ndjson",
		"application/x-ndjson",
		"text/ndjson",
		"text/x-ndjson",
	)

	lang.SetFileExtensions(types.JsonLines,
		"jsonl",
		"jsonlines",
		"json-lines",
		"jsons",
		"jsonseq",
		"json-seq",
		"ldjson",
		"ndjson",
		"murex_history",
	)
}
