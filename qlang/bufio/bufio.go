package bufio

import (
	"bufio"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "bufio",

	"MaxScanTokenSize": bufio.MaxScanTokenSize,

	"ErrAdvanceTooFar":     bufio.ErrAdvanceTooFar,
	"ErrBufferFull":        bufio.ErrBufferFull,
	"ErrInvalidUnreadByte": bufio.ErrInvalidUnreadByte,
	"ErrInvalidUnreadRune": bufio.ErrInvalidUnreadRune,
	"ErrNegativeAdvance":   bufio.ErrNegativeAdvance,
	"ErrNegativeCount":     bufio.ErrNegativeCount,
	"ErrTooLong":           bufio.ErrTooLong,

	"scanBytes": bufio.ScanBytes,
	"scanLines": bufio.ScanLines,
	"scanRunes": bufio.ScanRunes,
	"scanWords": bufio.ScanWords,

	"ReadWriter":    qlang.StructOf((*bufio.ReadWriter)(nil)),
	"readWriter":    bufio.NewReadWriter,
	"Reader":        qlang.StructOf((*bufio.Reader)(nil)),
	"newReader":     bufio.NewReader,
	"newReaderSize": bufio.NewReaderSize,
	"Scanner":       qlang.StructOf((*bufio.Scanner)(nil)),
	"scanner":       bufio.NewScanner,
	"Writer":        qlang.StructOf((*bufio.Writer)(nil)),
	"newWriter":     bufio.NewWriter,
	"newWriterSize": bufio.NewWriterSize,
}
