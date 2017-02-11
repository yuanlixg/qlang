package mime

import (
	"mime"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "mime",

	"addExtensionType": mime.AddExtensionType,
	"formatMediaType":  mime.FormatMediaType,
	"parseMediaType":   mime.ParseMediaType,
	"typeByExtension":  mime.TypeByExtension,
}
