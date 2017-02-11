package des

import (
	"crypto/des"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "crypto/des",

	"BlockSize": des.BlockSize,

	"newCipher":          des.NewCipher,
	"newTripleDESCipher": des.NewTripleDESCipher,
}
