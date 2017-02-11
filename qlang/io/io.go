package io

import (
	"io"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "io",

	"EOF":              io.EOF,
	"ErrClosedPipe":    io.ErrClosedPipe,
	"ErrNoProgress":    io.ErrNoProgress,
	"ErrShortBuffer":   io.ErrShortBuffer,
	"ErrShortWrite":    io.ErrShortWrite,
	"ErrUnexpectedEOF": io.ErrUnexpectedEOF,

	"copy":        io.Copy,
	"copyN":       io.CopyN,
	"readAtLeast": io.ReadAtLeast,
	"readFull":    io.ReadFull,
	"writeString": io.WriteString,

	"limitReader": io.LimitReader,
	"multiReader": io.MultiReader,
	"teeReader":   io.TeeReader,
	"multiWriter": io.MultiWriter,

	"LimitedReader": qlang.StructOf((*io.LimitedReader)(nil)),
	"PipeReader":    qlang.StructOf((*io.PipeReader)(nil)),
	"pipe":          io.Pipe,
	"PipeWriter":    qlang.StructOf((*io.PipeWriter)(nil)),
	"SectionReader": qlang.StructOf((*io.SectionReader)(nil)),
	"sectionReader": io.NewSectionReader,
}
