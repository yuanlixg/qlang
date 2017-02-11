package proxy

import (
	"golang.org/x/net/proxy"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "golang.org/x/net/proxy",

	"Direct": proxy.Direct,

	"registerDialerType": proxy.RegisterDialerType,

	"fromEnvironment": proxy.FromEnvironment,
	"fromURL":         proxy.FromURL,
	"SOCKS5":          proxy.SOCKS5,

	"Auth":    qlang.StructOf((*proxy.Auth)(nil)),
	"PerHost": qlang.StructOf((*proxy.PerHost)(nil)),
	"perHost": proxy.NewPerHost,
}
