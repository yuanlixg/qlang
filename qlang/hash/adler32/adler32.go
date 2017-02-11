package adler32

import (
	"hash/adler32"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "hash/adler32",

	"Size": adler32.Size,

	"checksum": adler32.Checksum,
	"new":      adler32.New,
}
