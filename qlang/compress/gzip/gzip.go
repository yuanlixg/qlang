package gzip

import (
	"compress/gzip"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "compress/gzip",

	"BestCompression":    gzip.BestCompression,
	"BestSpeed":          gzip.BestSpeed,
	"DefaultCompression": gzip.DefaultCompression,
	"NoCompression":      gzip.NoCompression,

	"ErrChecksum": gzip.ErrChecksum,
	"ErrHeader":   gzip.ErrHeader,

	"Header":         qlang.StructOf((*gzip.Header)(nil)),
	"Reader":         qlang.StructOf((*gzip.Reader)(nil)),
	"reader":         gzip.NewReader,
	"Writer":         qlang.StructOf((*gzip.Writer)(nil)),
	"newWriter":      gzip.NewWriter,
	"newWriterLevel": gzip.NewWriterLevel,
}
