// +build go1.5

package io

import "io"

func init() {
	Exports["copyBuffer"] = io.CopyBuffer
}
