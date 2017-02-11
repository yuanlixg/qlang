package pprof

import (
	"net/http/pprof"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "net/http/pprof",

	"cmdline": pprof.Cmdline,
	"handler": pprof.Handler,
	"index":   pprof.Index,
	"profile": pprof.Profile,
	"symbol":  pprof.Symbol,
}
