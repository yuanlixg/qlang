package textproto

import (
	"net/textproto"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "net/textproto",

	"canonicalMIMEHeaderKey": textproto.CanonicalMIMEHeaderKey,
	"trimBytes":              textproto.TrimBytes,
	"trimString":             textproto.TrimString,

	"Conn":     qlang.StructOf((*textproto.Conn)(nil)),
	"conn":     textproto.NewConn,
	"dial":     textproto.Dial,
	"Pipeline": qlang.StructOf((*textproto.Pipeline)(nil)),
	"Reader":   qlang.StructOf((*textproto.Reader)(nil)),
	"reader":   textproto.NewReader,
	"Writer":   qlang.StructOf((*textproto.Writer)(nil)),
	"writer":   textproto.NewWriter,
}
