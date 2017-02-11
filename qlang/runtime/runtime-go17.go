// +build go1.7

package runtime

import (
	"runtime"

	qlang "qlang.io/spec"
)

func init() {
	Exports["keepAlive"] = runtime.KeepAlive
	Exports["setCgoTraceback"] = runtime.SetCgoTraceback
	Exports["Frame"] = qlang.StructOf((*runtime.Frame)(nil))
	Exports["Frames"] = qlang.StructOf((*runtime.Frames)(nil))
	Exports["callersFrames"] = runtime.CallersFrames
}
