// +build go1.5

package sha512

import "crypto/sha512"

func init() {
	Exports["Size224"] = sha512.Size224
	Exports["Size256"] = sha512.Size256
	Exports["new512_224"] = sha512.New512_224
	Exports["new512_256"] = sha512.New512_256
	Exports["sum512_224"] = sha512.Sum512_224
	Exports["sum512_256"] = sha512.Sum512_256
}
