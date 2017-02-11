// +build go1.7

package user

import (
	"os/user"

	qlang "qlang.io/spec"
)

func init() {
	Exports["Group"] = qlang.StructOf((*user.Group)(nil))
	Exports["lookupGroup"] = user.LookupGroup
	Exports["lookupGroupId"] = user.LookupGroupId
}
