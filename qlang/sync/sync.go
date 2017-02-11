package sync

import (
	"sync"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "sync",

	"Cond":      qlang.StructOf((*sync.Cond)(nil)),
	"cond":      sync.NewCond,
	"Mutex":     qlang.StructOf((*sync.Mutex)(nil)),
	"Once":      qlang.StructOf((*sync.Once)(nil)),
	"Pool":      qlang.StructOf((*sync.Pool)(nil)),
	"RWMutex":   qlang.StructOf((*sync.RWMutex)(nil)),
	"WaitGroup": qlang.StructOf((*sync.WaitGroup)(nil)),
}
