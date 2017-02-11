package crypto

import (
	"crypto"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "crypto",

	"MD4":       crypto.MD4,
	"MD5":       crypto.MD5,
	"MD5SHA1":   crypto.MD5SHA1,
	"RIPEMD160": crypto.RIPEMD160,
	"SHA1":      crypto.SHA1,
	"SHA224":    crypto.SHA224,
	"SHA256":    crypto.SHA256,
	"SHA384":    crypto.SHA384,
	"SHA3_224":  crypto.SHA3_224,
	"SHA3_256":  crypto.SHA3_256,
	"SHA3_384":  crypto.SHA3_384,
	"SHA3_512":  crypto.SHA3_512,
	"SHA512":    crypto.SHA512,

	"registerHash": crypto.RegisterHash,
}
