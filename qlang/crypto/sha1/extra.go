package sha1

import (
	"crypto/sha1"
	"encoding/hex"
)

func init() {
	Exports["sumstr"] = Sumstr
}

// -----------------------------------------------------------------------------

// Sumstr is hex.EncodeToString(sha1.Sum(b)).
//
func Sumstr(b []byte) string {

	v := sha1.Sum(b)
	return hex.EncodeToString(v[:])
}

// -----------------------------------------------------------------------------
