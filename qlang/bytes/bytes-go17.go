// +build go1.7

package bytes

import "bytes"

func init() {
	Exports["containsAny"] = bytes.ContainsAny
	Exports["containsRune"] = bytes.ContainsRune
}
