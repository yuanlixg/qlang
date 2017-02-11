// +build go1.5

package fcgi

import "net/http/fcgi"

func init() {
	Exports["ErrConnClosed"] = fcgi.ErrConnClosed
	Exports["ErrRequestAborted"] = fcgi.ErrRequestAborted
}
