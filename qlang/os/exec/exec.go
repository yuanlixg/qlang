package exec

import (
	"os/exec"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "os/exec",

	"ErrNotFound": exec.ErrNotFound,

	"lookPath": exec.LookPath,

	"Cmd":     qlang.StructOf((*exec.Cmd)(nil)),
	"command": exec.Command,
}
