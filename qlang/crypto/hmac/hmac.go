package hmac

import (
	"crypto/hmac"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "crypto/hmac",

	"equal": hmac.Equal,
	"new":   hmac.New,
}
