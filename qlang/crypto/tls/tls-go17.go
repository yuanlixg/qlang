// +build go1.7

package tls

import "crypto/tls"

func init() {
	Exports["RenegotiateFreelyAsClient"] = tls.RenegotiateFreelyAsClient
	Exports["RenegotiateNever"] = tls.RenegotiateNever
	Exports["RenegotiateOnceAsClient"] = tls.RenegotiateOnceAsClient
}
