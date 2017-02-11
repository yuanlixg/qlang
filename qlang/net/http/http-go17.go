// +build go1.7

package http

import "net/http"

func init() {
	Exports["StatusAlreadyReported"] = http.StatusAlreadyReported
	Exports["StatusFailedDependency"] = http.StatusFailedDependency
	Exports["StatusIMUsed"] = http.StatusIMUsed
	Exports["StatusInsufficientStorage"] = http.StatusInsufficientStorage
	Exports["StatusLocked"] = http.StatusLocked
	Exports["StatusLoopDetected"] = http.StatusLoopDetected
	Exports["StatusMultiStatus"] = http.StatusMultiStatus
	Exports["StatusNotExtended"] = http.StatusNotExtended
	Exports["StatusPermanentRedirect"] = http.StatusPermanentRedirect
	Exports["StatusProcessing"] = http.StatusProcessing
	Exports["StatusUnprocessableEntity"] = http.StatusUnprocessableEntity
	Exports["StatusUpgradeRequired"] = http.StatusUpgradeRequired
	Exports["StatusVariantAlsoNegotiates"] = http.StatusVariantAlsoNegotiates
	Exports["ErrUseLastResponse"] = http.ErrUseLastResponse
	Exports["LocalAddrContextKey"] = http.LocalAddrContextKey
	Exports["ServerContextKey"] = http.ServerContextKey
}
