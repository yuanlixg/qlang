// +build go1.5

package bytes

import "bytes"

func init() {
	Exports["lastIndexByte"] = bytes.LastIndexByte
}
