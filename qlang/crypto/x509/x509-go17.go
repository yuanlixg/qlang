// +build go1.7

package x509

import "crypto/x509"

func init() {
	Exports["systemCertPool"] = x509.SystemCertPool
}
