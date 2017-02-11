package list

import (
	"container/list"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "container/list",

	"Element": qlang.StructOf((*list.Element)(nil)),
	"List":    qlang.StructOf((*list.List)(nil)),
	"new":     list.New,
}
