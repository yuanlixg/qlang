package bytes

import (
	"bytes"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "bytes",

	"MinRead": bytes.MinRead,

	"ErrTooLarge": bytes.ErrTooLarge,

	"compare":        bytes.Compare,
	"contains":       bytes.Contains,
	"count":          bytes.Count,
	"equal":          bytes.Equal,
	"equalFold":      bytes.EqualFold,
	"fields":         bytes.Fields,
	"fieldsFunc":     bytes.FieldsFunc,
	"hasPrefix":      bytes.HasPrefix,
	"hasSuffix":      bytes.HasSuffix,
	"index":          bytes.Index,
	"indexAny":       bytes.IndexAny,
	"indexByte":      bytes.IndexByte,
	"indexFunc":      bytes.IndexFunc,
	"indexRune":      bytes.IndexRune,
	"join":           bytes.Join,
	"lastIndex":      bytes.LastIndex,
	"lastIndexAny":   bytes.LastIndexAny,
	"lastIndexFunc":  bytes.LastIndexFunc,
	"map":            bytes.Map,
	"repeat":         bytes.Repeat,
	"replace":        bytes.Replace,
	"runes":          bytes.Runes,
	"split":          bytes.Split,
	"splitAfter":     bytes.SplitAfter,
	"splitAfterN":    bytes.SplitAfterN,
	"splitN":         bytes.SplitN,
	"title":          bytes.Title,
	"toLower":        bytes.ToLower,
	"toLowerSpecial": bytes.ToLowerSpecial,
	"toTitle":        bytes.ToTitle,
	"toTitleSpecial": bytes.ToTitleSpecial,
	"toUpper":        bytes.ToUpper,
	"toUpperSpecial": bytes.ToUpperSpecial,
	"trim":           bytes.Trim,
	"trimFunc":       bytes.TrimFunc,
	"trimLeft":       bytes.TrimLeft,
	"trimLeftFunc":   bytes.TrimLeftFunc,
	"trimPrefix":     bytes.TrimPrefix,
	"trimRight":      bytes.TrimRight,
	"trimRightFunc":  bytes.TrimRightFunc,
	"trimSpace":      bytes.TrimSpace,
	"trimSuffix":     bytes.TrimSuffix,

	"Buffer":          qlang.StructOf((*bytes.Buffer)(nil)),
	"newBuffer":       bytes.NewBuffer,
	"newBufferString": bytes.NewBufferString,
	"Reader":          qlang.StructOf((*bytes.Reader)(nil)),
	"reader":          bytes.NewReader,
}
