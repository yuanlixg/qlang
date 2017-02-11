// +build go1.7

package httptest

import "net/http/httptest"

func init() {
	Exports["newRequest"] = httptest.NewRequest
}
