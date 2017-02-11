package flate

import (
	"compress/flate"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "compress/flate",

	"BestCompression":    flate.BestCompression,
	"BestSpeed":          flate.BestSpeed,
	"DefaultCompression": flate.DefaultCompression,
	"NoCompression":      flate.NoCompression,

	"newReader":     flate.NewReader,
	"newReaderDict": flate.NewReaderDict,

	"Writer":        qlang.StructOf((*flate.Writer)(nil)),
	"newWriter":     flate.NewWriter,
	"newWriterDict": flate.NewWriterDict,
}
