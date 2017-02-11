package runtime

import (
	"runtime"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "runtime",

	"Compiler": runtime.Compiler,
	"GOARCH":   runtime.GOARCH,
	"GOOS":     runtime.GOOS,

	"MemProfileRate": runtime.MemProfileRate,

	"blockProfile":        runtime.BlockProfile,
	"breakpoint":          runtime.Breakpoint,
	"CPUProfile":          runtime.CPUProfile,
	"caller":              runtime.Caller,
	"callers":             runtime.Callers,
	"GC":                  runtime.GC,
	"GOMAXPROCS":          runtime.GOMAXPROCS,
	"GOROOT":              runtime.GOROOT,
	"goexit":              runtime.Goexit,
	"goroutineProfile":    runtime.GoroutineProfile,
	"gosched":             runtime.Gosched,
	"lockOSThread":        runtime.LockOSThread,
	"memProfile":          runtime.MemProfile,
	"numCPU":              runtime.NumCPU,
	"numCgoCall":          runtime.NumCgoCall,
	"numGoroutine":        runtime.NumGoroutine,
	"readMemStats":        runtime.ReadMemStats,
	"setBlockProfileRate": runtime.SetBlockProfileRate,
	"setCPUProfileRate":   runtime.SetCPUProfileRate,
	"setFinalizer":        runtime.SetFinalizer,
	"stack":               runtime.Stack,
	"threadCreateProfile": runtime.ThreadCreateProfile,
	"unlockOSThread":      runtime.UnlockOSThread,
	"version":             runtime.Version,

	"BlockProfileRecord": qlang.StructOf((*runtime.BlockProfileRecord)(nil)),
	"Func":               qlang.StructOf((*runtime.Func)(nil)),
	"funcForPC":          runtime.FuncForPC,
	"MemProfileRecord":   qlang.StructOf((*runtime.MemProfileRecord)(nil)),
	"MemStats":           qlang.StructOf((*runtime.MemStats)(nil)),
	"StackRecord":        qlang.StructOf((*runtime.StackRecord)(nil)),
}
