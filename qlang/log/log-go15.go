// +build go1.5

package log

import "log"

func init() {
	Exports["LUTC"] = log.LUTC
	Exports["output"] = log.Output
}
