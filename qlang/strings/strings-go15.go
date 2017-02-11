// +build go1.5

package strings

import "strings"

func init() {
	Exports["compare"] = strings.Compare
	Exports["lastIndexByte"] = strings.LastIndexByte
}
