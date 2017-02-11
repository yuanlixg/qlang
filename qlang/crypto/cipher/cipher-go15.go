// +build go1.5

package cipher

import "crypto/cipher"

func init() {
	Exports["newGCMWithNonceSize"] = cipher.NewGCMWithNonceSize
}
