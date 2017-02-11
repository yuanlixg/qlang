// +build go1.7

package flate

import "compress/flate"

func init() {
	Exports["HuffmanOnly"] = flate.HuffmanOnly
}
