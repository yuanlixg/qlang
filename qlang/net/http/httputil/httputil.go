package httputil

import (
	"net/http/httputil"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "net/http/httputil",

	"ErrClosed":      httputil.ErrClosed,
	"ErrLineTooLong": httputil.ErrLineTooLong,
	"ErrPersistEOF":  httputil.ErrPersistEOF,
	"ErrPipeline":    httputil.ErrPipeline,

	"dumpRequest":      httputil.DumpRequest,
	"dumpRequestOut":   httputil.DumpRequestOut,
	"dumpResponse":     httputil.DumpResponse,
	"newChunkedReader": httputil.NewChunkedReader,
	"newChunkedWriter": httputil.NewChunkedWriter,

	"ClientConn":                qlang.StructOf((*httputil.ClientConn)(nil)),
	"clientConn":                httputil.NewClientConn,
	"newProxyClientConn":        httputil.NewProxyClientConn,
	"ReverseProxy":              qlang.StructOf((*httputil.ReverseProxy)(nil)),
	"newSingleHostReverseProxy": httputil.NewSingleHostReverseProxy,
	"ServerConn":                qlang.StructOf((*httputil.ServerConn)(nil)),
	"serverConn":                httputil.NewServerConn,
}
