package md5

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

func init() {
	Exports["sumstr"] = Sumstr
	Exports["hash"] = Hash
}

// -----------------------------------------------------------------------------

// Hash returns md5 sum of (sep, args...) serialization.
//
func Hash(sep interface{}, args ...interface{}) string {

	h := md5.New()

	var bsep []byte
	switch v := sep.(type) {
	case []byte:
		bsep = v
	case string:
		bsep = []byte(v)
	default:
		panic("md5.Hash: invalid argument type, require []byte or string")
	}

	for i, arg := range args {
		if i > 0 {
			h.Write(bsep)
		}
		switch v := arg.(type) {
		case []byte:
			h.Write(v)
		case string:
			io.WriteString(h, v)
		case error:
		default:
			panic("md5.Hash: invalid argument type, require []byte or string")
		}
	}

	return hex.EncodeToString(h.Sum(nil))
}

// Sumstr is hex.EncodeToString(md5.Sum(b)).
//
func Sumstr(b []byte) string {

	v := md5.Sum(b)
	return hex.EncodeToString(v[:])
}

// -----------------------------------------------------------------------------
