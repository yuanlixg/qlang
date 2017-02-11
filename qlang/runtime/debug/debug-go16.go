// +build go1.6

package debug

import "runtime/debug"

func init() {
	Exports["setTraceback"] = debug.SetTraceback
}
