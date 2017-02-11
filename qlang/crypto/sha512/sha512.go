package sha512

import (
	"crypto/sha512"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "crypto/sha512",

	"BlockSize": sha512.BlockSize,
	"Size":      sha512.Size,
	"Size384":   sha512.Size384,

	"new":    sha512.New,
	"new384": sha512.New384,
	"sum384": sha512.Sum384,
	"sum512": sha512.Sum512,
}
