package base64

import (
	"encoding/base64"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "encoding/base64",

	"StdEncoding": base64.StdEncoding,
	"URLEncoding": base64.URLEncoding,

	"newDecoder": base64.NewDecoder,
	"newEncoder": base64.NewEncoder,

	"Encoding": qlang.StructOf((*base64.Encoding)(nil)),
	"encoding": base64.NewEncoding,
}
