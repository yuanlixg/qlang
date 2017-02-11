package fnv

import (
	"hash/fnv"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "hash/fnv",

	"new32":  fnv.New32,
	"new32a": fnv.New32a,
	"new64":  fnv.New64,
	"new64a": fnv.New64a,
}
