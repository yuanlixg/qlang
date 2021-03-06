package pem

import (
	"encoding/pem"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "encoding/pem",

	"encode":         pem.Encode,
	"encodeToMemory": pem.EncodeToMemory,

	"Block":  qlang.StructOf((*pem.Block)(nil)),
	"decode": pem.Decode,
}
