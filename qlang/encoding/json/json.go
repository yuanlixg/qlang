package json

import (
	"encoding/json"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "encoding/json",

	"compact":       json.Compact,
	"HTMLEscape":    json.HTMLEscape,
	"indent":        json.Indent,
	"marshal":       json.Marshal,
	"marshalIndent": json.MarshalIndent,
	"unmarshal":     json.Unmarshal,

	"Decoder": qlang.StructOf((*json.Decoder)(nil)),
	"decoder": json.NewDecoder,
	"Encoder": qlang.StructOf((*json.Encoder)(nil)),
	"encoder": json.NewEncoder,
}
