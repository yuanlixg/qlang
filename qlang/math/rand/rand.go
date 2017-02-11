package rand

import (
	"math/rand"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "math/rand",

	"expFloat64":  rand.ExpFloat64,
	"float32":     rand.Float32,
	"float64":     rand.Float64,
	"int":         rand.Int,
	"int31":       rand.Int31,
	"int31n":      rand.Int31n,
	"int63":       rand.Int63,
	"int63n":      rand.Int63n,
	"intn":        rand.Intn,
	"normFloat64": rand.NormFloat64,
	"perm":        rand.Perm,
	"seed":        rand.Seed,
	"uint32":      rand.Uint32,

	"source": rand.NewSource,

	"Rand": qlang.StructOf((*rand.Rand)(nil)),
	"new":  rand.New,
	"Zipf": qlang.StructOf((*rand.Zipf)(nil)),
	"zipf": rand.NewZipf,
}
