// +build go1.5

package mail

import (
	"net/mail"

	qlang "qlang.io/spec"
)

func init() {
	Exports["AddressParser"] = qlang.StructOf((*mail.AddressParser)(nil))
}
