// +build go1.5

package runtime

import "runtime"

func init() {
	Exports["readTrace"] = runtime.ReadTrace
	Exports["startTrace"] = runtime.StartTrace
	Exports["stopTrace"] = runtime.StopTrace
}
