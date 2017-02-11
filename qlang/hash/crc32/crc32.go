package crc32

import (
	"hash/crc32"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "hash/crc32",

	"Castagnoli": uint64(crc32.Castagnoli),
	"IEEE":       uint64(crc32.IEEE),
	"Koopman":    uint64(crc32.Koopman),
	"Size":       crc32.Size,

	"IEEETable": crc32.IEEETable,

	"checksum":     crc32.Checksum,
	"checksumIEEE": crc32.ChecksumIEEE,
	"new":          crc32.New,
	"newIEEE":      crc32.NewIEEE,
	"update":       crc32.Update,
}
