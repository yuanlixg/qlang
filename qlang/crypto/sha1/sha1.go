package sha1

import (
	"crypto/sha1"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "crypto/sha1",

	"BlockSize": sha1.BlockSize,
	"Size":      sha1.Size,

	"new": sha1.New,
	"sum": sha1.Sum,
}
