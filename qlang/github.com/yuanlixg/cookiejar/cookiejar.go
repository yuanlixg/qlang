package cookiejar

import (
	"github.com/yuanlixg/cookiejar"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "github.com/yuanlixg/cookiejar",

	"Jar":     qlang.StructOf((*cookiejar.Jar)(nil)),
	"clone":   cookiejar.Clone,
	"new":     cookiejar.New,
	"Options": qlang.StructOf((*cookiejar.Options)(nil)),
}
