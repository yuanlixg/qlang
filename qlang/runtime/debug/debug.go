package debug

import (
	"runtime/debug"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "runtime/debug",

	"freeOSMemory":    debug.FreeOSMemory,
	"printStack":      debug.PrintStack,
	"readGCStats":     debug.ReadGCStats,
	"setGCPercent":    debug.SetGCPercent,
	"setMaxStack":     debug.SetMaxStack,
	"setMaxThreads":   debug.SetMaxThreads,
	"setPanicOnFault": debug.SetPanicOnFault,
	"stack":           debug.Stack,
	"writeHeapDump":   debug.WriteHeapDump,

	"GCStats": qlang.StructOf((*debug.GCStats)(nil)),
}
