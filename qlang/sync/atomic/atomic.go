package atomic

import (
	"sync/atomic"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "sync/atomic",

	"addInt32":              atomic.AddInt32,
	"addInt64":              atomic.AddInt64,
	"addUint32":             atomic.AddUint32,
	"addUint64":             atomic.AddUint64,
	"addUintptr":            atomic.AddUintptr,
	"compareAndSwapInt32":   atomic.CompareAndSwapInt32,
	"compareAndSwapInt64":   atomic.CompareAndSwapInt64,
	"compareAndSwapPointer": atomic.CompareAndSwapPointer,
	"compareAndSwapUint32":  atomic.CompareAndSwapUint32,
	"compareAndSwapUint64":  atomic.CompareAndSwapUint64,
	"compareAndSwapUintptr": atomic.CompareAndSwapUintptr,
	"loadInt32":             atomic.LoadInt32,
	"loadInt64":             atomic.LoadInt64,
	"loadPointer":           atomic.LoadPointer,
	"loadUint32":            atomic.LoadUint32,
	"loadUint64":            atomic.LoadUint64,
	"loadUintptr":           atomic.LoadUintptr,
	"storeInt32":            atomic.StoreInt32,
	"storeInt64":            atomic.StoreInt64,
	"storePointer":          atomic.StorePointer,
	"storeUint32":           atomic.StoreUint32,
	"storeUint64":           atomic.StoreUint64,
	"storeUintptr":          atomic.StoreUintptr,
	"swapInt32":             atomic.SwapInt32,
	"swapInt64":             atomic.SwapInt64,
	"swapPointer":           atomic.SwapPointer,
	"swapUint32":            atomic.SwapUint32,
	"swapUint64":            atomic.SwapUint64,
	"swapUintptr":           atomic.SwapUintptr,

	"Value": qlang.StructOf((*atomic.Value)(nil)),
}
