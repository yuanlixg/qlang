package httptest

import (
	"net/http/httptest"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "net/http/httptest",

	"DefaultRemoteAddr": httptest.DefaultRemoteAddr,

	"ResponseRecorder":   qlang.StructOf((*httptest.ResponseRecorder)(nil)),
	"newRecorder":        httptest.NewRecorder,
	"Server":             qlang.StructOf((*httptest.Server)(nil)),
	"server":             httptest.NewServer,
	"newTLSServer":       httptest.NewTLSServer,
	"newUnstartedServer": httptest.NewUnstartedServer,
}
