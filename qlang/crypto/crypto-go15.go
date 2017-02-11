// +build go1.5

package crypto

import "crypto"

func init() {
	Exports["SHA512_224"] = crypto.SHA512_224
	Exports["SHA512_256"] = crypto.SHA512_256
}
