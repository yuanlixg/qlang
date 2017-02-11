package reflect

import (
	"reflect"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "reflect",

	"Array":         reflect.Array,
	"Bool":          reflect.Bool,
	"BothDir":       reflect.BothDir,
	"Chan":          reflect.Chan,
	"Complex128":    reflect.Complex128,
	"Complex64":     reflect.Complex64,
	"Float32":       reflect.Float32,
	"Float64":       reflect.Float64,
	"Func":          reflect.Func,
	"Int":           reflect.Int,
	"Int16":         reflect.Int16,
	"Int32":         reflect.Int32,
	"Int64":         reflect.Int64,
	"Int8":          reflect.Int8,
	"Interface":     reflect.Interface,
	"Invalid":       reflect.Invalid,
	"Map":           reflect.Map,
	"Ptr":           reflect.Ptr,
	"RecvDir":       reflect.RecvDir,
	"SelectDefault": reflect.SelectDefault,
	"SelectRecv":    reflect.SelectRecv,
	"SelectSend":    reflect.SelectSend,
	"SendDir":       reflect.SendDir,
	"Slice":         reflect.Slice,
	"String":        reflect.String,
	"Struct":        reflect.Struct,
	"Uint":          reflect.Uint,
	"Uint16":        reflect.Uint16,
	"Uint32":        reflect.Uint32,
	"Uint64":        reflect.Uint64,
	"Uint8":         reflect.Uint8,
	"Uintptr":       reflect.Uintptr,
	"UnsafePointer": reflect.UnsafePointer,

	"copy":      reflect.Copy,
	"deepEqual": reflect.DeepEqual,
	"select":    reflect.Select,

	"chanOf":  reflect.ChanOf,
	"mapOf":   reflect.MapOf,
	"ptrTo":   reflect.PtrTo,
	"sliceOf": reflect.SliceOf,
	"typeOf":  reflect.TypeOf,

	"Method":       qlang.StructOf((*reflect.Method)(nil)),
	"SelectCase":   qlang.StructOf((*reflect.SelectCase)(nil)),
	"SliceHeader":  qlang.StructOf((*reflect.SliceHeader)(nil)),
	"StringHeader": qlang.StructOf((*reflect.StringHeader)(nil)),
	"StructField":  qlang.StructOf((*reflect.StructField)(nil)),
	"Value":        qlang.StructOf((*reflect.Value)(nil)),
	"append":       reflect.Append,
	"appendSlice":  reflect.AppendSlice,
	"indirect":     reflect.Indirect,
	"makeChan":     reflect.MakeChan,
	"makeFunc":     reflect.MakeFunc,
	"makeMap":      reflect.MakeMap,
	"makeSlice":    reflect.MakeSlice,
	"new":          reflect.New,
	"newAt":        reflect.NewAt,
	"valueOf":      reflect.ValueOf,
	"zero":         reflect.Zero,
}
