package binary

import (
	"encoding/binary"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "encoding/binary",

	"MaxVarintLen16": binary.MaxVarintLen16,
	"MaxVarintLen32": binary.MaxVarintLen32,
	"MaxVarintLen64": binary.MaxVarintLen64,

	"BigEndian":    &binary.BigEndian,
	"LittleEndian": &binary.LittleEndian,

	"putUvarint":  binary.PutUvarint,
	"putVarint":   binary.PutVarint,
	"read":        binary.Read,
	"readUvarint": binary.ReadUvarint,
	"readVarint":  binary.ReadVarint,
	"size":        binary.Size,
	"uvarint":     binary.Uvarint,
	"varint":      binary.Varint,
	"write":       binary.Write,
}
