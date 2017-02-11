// +build go1.6

package rand

import "math/rand"

func init() {
	Exports["read"] = rand.Read
}
