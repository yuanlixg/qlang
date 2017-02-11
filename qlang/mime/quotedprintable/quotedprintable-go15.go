// +build go1.5

package quotedprintable

import (
	"mime/quotedprintable"

	qlang "qlang.io/spec"
)

func init() {
	Exports["Reader"] = qlang.StructOf((*quotedprintable.Reader)(nil))
	Exports["reader"] = quotedprintable.NewReader
	Exports["Writer"] = qlang.StructOf((*quotedprintable.Writer)(nil))
	Exports["writer"] = quotedprintable.NewWriter
}
