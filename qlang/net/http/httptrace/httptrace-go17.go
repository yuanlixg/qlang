// +build go1.7

package httptrace

import (
	"net/http/httptrace"

	qlang "qlang.io/spec"
)

func init() {
	Exports["withClientTrace"] = httptrace.WithClientTrace
	Exports["ClientTrace"] = qlang.StructOf((*httptrace.ClientTrace)(nil))
	Exports["contextClientTrace"] = httptrace.ContextClientTrace
	Exports["DNSDoneInfo"] = qlang.StructOf((*httptrace.DNSDoneInfo)(nil))
	Exports["DNSStartInfo"] = qlang.StructOf((*httptrace.DNSStartInfo)(nil))
	Exports["GotConnInfo"] = qlang.StructOf((*httptrace.GotConnInfo)(nil))
	Exports["WroteRequestInfo"] = qlang.StructOf((*httptrace.WroteRequestInfo)(nil))
}
