// +build go1.5

package rsa

import (
	"crypto/rsa"

	qlang "qlang.io/spec"
)

func init() {
	Exports["OAEPOptions"] = qlang.StructOf((*rsa.OAEPOptions)(nil))
	Exports["PKCS1v15DecryptOptions"] = qlang.StructOf((*rsa.PKCS1v15DecryptOptions)(nil))
}
