// +build go1.5

package tls

import "crypto/tls"

func init() {
	Exports["TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384"] = tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384
	Exports["TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384"] = tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384
}
