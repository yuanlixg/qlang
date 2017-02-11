// +build go1.5

package pprof

import "net/http/pprof"

func init() {
	Exports["trace"] = pprof.Trace
}
