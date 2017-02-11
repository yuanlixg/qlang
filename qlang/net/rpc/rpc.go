package rpc

import (
	"net/rpc"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "net/rpc",

	"DefaultDebugPath": rpc.DefaultDebugPath,
	"DefaultRPCPath":   rpc.DefaultRPCPath,

	"DefaultServer": rpc.DefaultServer,
	"ErrShutdown":   rpc.ErrShutdown,

	"accept":       rpc.Accept,
	"handleHTTP":   rpc.HandleHTTP,
	"register":     rpc.Register,
	"registerName": rpc.RegisterName,
	"serveCodec":   rpc.ServeCodec,
	"serveConn":    rpc.ServeConn,
	"serveRequest": rpc.ServeRequest,

	"Call":               qlang.StructOf((*rpc.Call)(nil)),
	"Client":             qlang.StructOf((*rpc.Client)(nil)),
	"newClient":          rpc.NewClient,
	"newClientWithCodec": rpc.NewClientWithCodec,
	"dial":               rpc.Dial,
	"dialHTTP":           rpc.DialHTTP,
	"dialHTTPPath":       rpc.DialHTTPPath,
	"Request":            qlang.StructOf((*rpc.Request)(nil)),
	"Response":           qlang.StructOf((*rpc.Response)(nil)),
	"Server":             qlang.StructOf((*rpc.Server)(nil)),
	"server":             rpc.NewServer,
}
