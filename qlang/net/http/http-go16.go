// +build go1.6

package http

import "net/http"

func init() {
	Exports["MethodConnect"] = http.MethodConnect
	Exports["MethodDelete"] = http.MethodDelete
	Exports["MethodGet"] = http.MethodGet
	Exports["MethodHead"] = http.MethodHead
	Exports["MethodOptions"] = http.MethodOptions
	Exports["MethodPatch"] = http.MethodPatch
	Exports["MethodPost"] = http.MethodPost
	Exports["MethodPut"] = http.MethodPut
	Exports["MethodTrace"] = http.MethodTrace
	Exports["StatusNetworkAuthenticationRequired"] = http.StatusNetworkAuthenticationRequired
	Exports["StatusPreconditionRequired"] = http.StatusPreconditionRequired
	Exports["StatusRequestHeaderFieldsTooLarge"] = http.StatusRequestHeaderFieldsTooLarge
	Exports["StatusTooManyRequests"] = http.StatusTooManyRequests
	Exports["StatusUnavailableForLegalReasons"] = http.StatusUnavailableForLegalReasons
	Exports["ErrSkipAltProtocol"] = http.ErrSkipAltProtocol
}
