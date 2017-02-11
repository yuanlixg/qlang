package hex

import (
	"encoding/hex"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "encoding/hex",

	"ErrLength": hex.ErrLength,

	"decode":         hex.Decode,
	"decodeString":   hex.DecodeString,
	"decodedLen":     hex.DecodedLen,
	"dump":           hex.Dump,
	"dumper":         hex.Dumper,
	"encode":         hex.Encode,
	"encodeToString": hex.EncodeToString,
	"encodedLen":     hex.EncodedLen,
}
