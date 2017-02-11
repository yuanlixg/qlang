package ioutil

import (
	qlang "qlang.io/spec"
)

func init() {
	Exports["_initSafe"] = _initSafe
}

func _initSafe(mod qlang.Module) {

	mod.Disable("readDir", "readFile", "writeFile")
}
