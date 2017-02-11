package cgi

import (
	"net/http/cgi"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "net/http/cgi",

	"request":        cgi.Request,
	"requestFromMap": cgi.RequestFromMap,
	"serve":          cgi.Serve,

	"Handler": qlang.StructOf((*cgi.Handler)(nil)),
}
