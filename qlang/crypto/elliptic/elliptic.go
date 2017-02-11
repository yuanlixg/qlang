package elliptic

import (
	"crypto/elliptic"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "crypto/elliptic",

	"generateKey": elliptic.GenerateKey,
	"marshal":     elliptic.Marshal,
	"unmarshal":   elliptic.Unmarshal,

	"P224": elliptic.P224,
	"P256": elliptic.P256,
	"P384": elliptic.P384,
	"P521": elliptic.P521,

	"CurveParams": qlang.StructOf((*elliptic.CurveParams)(nil)),
}
