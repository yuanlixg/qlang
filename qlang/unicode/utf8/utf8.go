package utf8

import (
	"unicode/utf8"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "unicode/utf8",

	"MaxRune":   utf8.MaxRune,
	"RuneError": utf8.RuneError,
	"RuneSelf":  utf8.RuneSelf,
	"UTFMax":    utf8.UTFMax,

	"decodeLastRune":         utf8.DecodeLastRune,
	"decodeLastRuneInString": utf8.DecodeLastRuneInString,
	"decodeRune":             utf8.DecodeRune,
	"decodeRuneInString":     utf8.DecodeRuneInString,
	"encodeRune":             utf8.EncodeRune,
	"fullRune":               utf8.FullRune,
	"fullRuneInString":       utf8.FullRuneInString,
	"runeCount":              utf8.RuneCount,
	"runeCountInString":      utf8.RuneCountInString,
	"runeLen":                utf8.RuneLen,
	"runeStart":              utf8.RuneStart,
	"valid":                  utf8.Valid,
	"validRune":              utf8.ValidRune,
	"validString":            utf8.ValidString,
}
