package subtle

import (
	"crypto/subtle"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "crypto/subtle",

	"constantTimeByteEq":   subtle.ConstantTimeByteEq,
	"constantTimeCompare":  subtle.ConstantTimeCompare,
	"constantTimeCopy":     subtle.ConstantTimeCopy,
	"constantTimeEq":       subtle.ConstantTimeEq,
	"constantTimeLessOrEq": subtle.ConstantTimeLessOrEq,
	"constantTimeSelect":   subtle.ConstantTimeSelect,
}
