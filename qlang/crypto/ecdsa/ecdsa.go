package ecdsa

import (
	"crypto/ecdsa"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "crypto/ecdsa",

	"sign":   ecdsa.Sign,
	"verify": ecdsa.Verify,

	"PrivateKey":  qlang.StructOf((*ecdsa.PrivateKey)(nil)),
	"generateKey": ecdsa.GenerateKey,
	"PublicKey":   qlang.StructOf((*ecdsa.PublicKey)(nil)),
}
