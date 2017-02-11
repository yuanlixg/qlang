package cookiejar

import (
	"net/http/cookiejar"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "net/http/cookiejar",

	"Jar":     qlang.StructOf((*cookiejar.Jar)(nil)),
	"new":     cookiejar.New,
	"Options": qlang.StructOf((*cookiejar.Options)(nil)),
}
