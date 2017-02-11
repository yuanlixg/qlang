package cipher

import (
	"crypto/cipher"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "crypto/cipher",

	"newGCM":          cipher.NewGCM,
	"newCBCDecrypter": cipher.NewCBCDecrypter,
	"newCBCEncrypter": cipher.NewCBCEncrypter,
	"newCFBDecrypter": cipher.NewCFBDecrypter,
	"newCFBEncrypter": cipher.NewCFBEncrypter,
	"newCTR":          cipher.NewCTR,
	"newOFB":          cipher.NewOFB,

	"StreamReader": qlang.StructOf((*cipher.StreamReader)(nil)),
	"StreamWriter": qlang.StructOf((*cipher.StreamWriter)(nil)),
}
