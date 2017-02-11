package dsa

import (
	"crypto/dsa"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "crypto/dsa",

	"L1024N160": dsa.L1024N160,
	"L2048N224": dsa.L2048N224,
	"L2048N256": dsa.L2048N256,
	"L3072N256": dsa.L3072N256,

	"ErrInvalidPublicKey": dsa.ErrInvalidPublicKey,

	"generateKey":        dsa.GenerateKey,
	"generateParameters": dsa.GenerateParameters,
	"sign":               dsa.Sign,
	"verify":             dsa.Verify,

	"Parameters": qlang.StructOf((*dsa.Parameters)(nil)),
	"PrivateKey": qlang.StructOf((*dsa.PrivateKey)(nil)),
	"PublicKey":  qlang.StructOf((*dsa.PublicKey)(nil)),
}
