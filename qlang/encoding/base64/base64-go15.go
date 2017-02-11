// +build go1.5

package base64

import "encoding/base64"

func init() {
	Exports["NoPadding"] = base64.NoPadding
	Exports["StdPadding"] = base64.StdPadding
	Exports["RawStdEncoding"] = base64.RawStdEncoding
	Exports["RawURLEncoding"] = base64.RawURLEncoding
}
