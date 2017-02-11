package big

import (
	"math/big"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "math/big",

	"MaxBase": big.MaxBase,

	"Int": qlang.StructOf((*big.Int)(nil)),
	"int": big.NewInt,
	"Rat": qlang.StructOf((*big.Rat)(nil)),
	"rat": big.NewRat,
}
