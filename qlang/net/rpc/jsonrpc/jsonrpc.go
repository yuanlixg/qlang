package jsonrpc

import (
	"net/rpc/jsonrpc"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "net/rpc/jsonrpc",

	"dial":           jsonrpc.Dial,
	"newClient":      jsonrpc.NewClient,
	"newClientCodec": jsonrpc.NewClientCodec,
	"newServerCodec": jsonrpc.NewServerCodec,
	"serveConn":      jsonrpc.ServeConn,
}
