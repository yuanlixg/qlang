package pprof

import (
	"runtime/pprof"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "runtime/pprof",

	"profiles":         pprof.Profiles,
	"startCPUProfile":  pprof.StartCPUProfile,
	"stopCPUProfile":   pprof.StopCPUProfile,
	"writeHeapProfile": pprof.WriteHeapProfile,

	"Profile": qlang.StructOf((*pprof.Profile)(nil)),
	"profile": pprof.NewProfile,
	"lookup":  pprof.Lookup,
}
