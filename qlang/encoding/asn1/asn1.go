package asn1

import (
	"encoding/asn1"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "encoding/asn1",

	"marshal":             asn1.Marshal,
	"unmarshal":           asn1.Unmarshal,
	"unmarshalWithParams": asn1.UnmarshalWithParams,

	"BitString": qlang.StructOf((*asn1.BitString)(nil)),
	"RawValue":  qlang.StructOf((*asn1.RawValue)(nil)),
}
