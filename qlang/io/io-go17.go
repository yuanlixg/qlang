// +build go1.7

package io

import "io"

func init() {
	Exports["SeekCurrent"] = io.SeekCurrent
	Exports["SeekEnd"] = io.SeekEnd
	Exports["SeekStart"] = io.SeekStart
}
