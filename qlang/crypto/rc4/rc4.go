package rc4

import (
	"crypto/rc4"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "crypto/rc4",

	"Cipher": qlang.StructOf((*rc4.Cipher)(nil)),
	"cipher": rc4.NewCipher,
}
