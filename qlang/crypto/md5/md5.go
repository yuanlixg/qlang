package md5

import (
	"crypto/md5"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "crypto/md5",

	"BlockSize": md5.BlockSize,
	"Size":      md5.Size,

	"new": md5.New,
	"sum": md5.Sum,
}
