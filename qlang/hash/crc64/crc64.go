package crc64

import (
	"hash/crc64"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "hash/crc64",

	"ECMA": uint64(crc64.ECMA),
	"ISO":  uint64(crc64.ISO),
	"Size": crc64.Size,

	"checksum": crc64.Checksum,
	"new":      crc64.New,
	"update":   crc64.Update,
}
