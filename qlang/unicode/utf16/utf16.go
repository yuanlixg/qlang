package utf16

import (
	"unicode/utf16"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "unicode/utf16",

	"decode":      utf16.Decode,
	"decodeRune":  utf16.DecodeRune,
	"encode":      utf16.Encode,
	"encodeRune":  utf16.EncodeRune,
	"isSurrogate": utf16.IsSurrogate,
}
