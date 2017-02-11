package multipart

import (
	"mime/multipart"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "mime/multipart",

	"FileHeader": qlang.StructOf((*multipart.FileHeader)(nil)),
	"Form":       qlang.StructOf((*multipart.Form)(nil)),
	"Part":       qlang.StructOf((*multipart.Part)(nil)),
	"Reader":     qlang.StructOf((*multipart.Reader)(nil)),
	"reader":     multipart.NewReader,
	"Writer":     qlang.StructOf((*multipart.Writer)(nil)),
	"writer":     multipart.NewWriter,
}
