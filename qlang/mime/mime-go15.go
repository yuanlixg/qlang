// +build go1.5

package mime

import (
	"mime"

	qlang "qlang.io/spec"
)

func init() {
	Exports["BEncoding"] = mime.BEncoding
	Exports["QEncoding"] = mime.QEncoding
	Exports["extensionsByType"] = mime.ExtensionsByType
	Exports["WordDecoder"] = qlang.StructOf((*mime.WordDecoder)(nil))
}
