package os

import (
	"os"
	"strconv"

	qlang "qlang.io/spec"
)

func init() {
	Exports["_initSafe"] = _initSafe
}

func _initSafe(mod qlang.Module) {

	mod.Disable("open")
	mod.Disable("getenv")
	mod.Exports["exit"] = SafeExit
}

// SafeExit is a safe way to quit qlang application.
//
func SafeExit(code int) {

	panic("exit " + strconv.Itoa(code))
}

func exit() {
	os.Exit(0)
}

func safeExit() {
	panic("exit")
}

func _initSafe2(mod qlang.Module) {
	mod.Exports["exit"] = safeExit
}

// InlineExports is the export table of this module.
//
var InlineExports = map[string]interface{}{
	"exit":      exit,
	"_initSafe": _initSafe2,
}

