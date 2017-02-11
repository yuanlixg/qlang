// +build go1.5

package trace

import "runtime/trace"

func init() {
	Exports["start"] = trace.Start
	Exports["stop"] = trace.Stop
}
