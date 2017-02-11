// +build go1.6

package tls

import "crypto/tls"

func init() {
	Exports["TLS_RSA_WITH_AES_128_GCM_SHA256"] = tls.TLS_RSA_WITH_AES_128_GCM_SHA256
	Exports["TLS_RSA_WITH_AES_256_GCM_SHA384"] = tls.TLS_RSA_WITH_AES_256_GCM_SHA384
}
