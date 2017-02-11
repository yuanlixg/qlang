package mail

import (
	"net/mail"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "net/mail",

	"ErrHeaderNotPresent": mail.ErrHeaderNotPresent,

	"parseAddressList": mail.ParseAddressList,

	"Address":      qlang.StructOf((*mail.Address)(nil)),
	"parseAddress": mail.ParseAddress,
	"Message":      qlang.StructOf((*mail.Message)(nil)),
	"readMessage":  mail.ReadMessage,
}
