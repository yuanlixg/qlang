package user

import (
	"os/user"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "os/user",

	"User":     qlang.StructOf((*user.User)(nil)),
	"current":  user.Current,
	"lookup":   user.Lookup,
	"lookupId": user.LookupId,
}
