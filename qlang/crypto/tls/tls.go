package tls

import (
	"crypto/tls"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "crypto/tls",

	"CurveP256":                               tls.CurveP256,
	"CurveP384":                               tls.CurveP384,
	"CurveP521":                               tls.CurveP521,
	"NoClientCert":                            tls.NoClientCert,
	"RequestClientCert":                       tls.RequestClientCert,
	"RequireAndVerifyClientCert":              tls.RequireAndVerifyClientCert,
	"RequireAnyClientCert":                    tls.RequireAnyClientCert,
	"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA":    tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
	"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256": tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
	"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA":    tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
	"TLS_ECDHE_ECDSA_WITH_RC4_128_SHA":        tls.TLS_ECDHE_ECDSA_WITH_RC4_128_SHA,
	"TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA":     tls.TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA,
	"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA":      tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
	"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256":   tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
	"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA":      tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
	"TLS_ECDHE_RSA_WITH_RC4_128_SHA":          tls.TLS_ECDHE_RSA_WITH_RC4_128_SHA,
	"TLS_FALLBACK_SCSV":                       tls.TLS_FALLBACK_SCSV,
	"TLS_RSA_WITH_3DES_EDE_CBC_SHA":           tls.TLS_RSA_WITH_3DES_EDE_CBC_SHA,
	"TLS_RSA_WITH_AES_128_CBC_SHA":            tls.TLS_RSA_WITH_AES_128_CBC_SHA,
	"TLS_RSA_WITH_AES_256_CBC_SHA":            tls.TLS_RSA_WITH_AES_256_CBC_SHA,
	"TLS_RSA_WITH_RC4_128_SHA":                tls.TLS_RSA_WITH_RC4_128_SHA,
	"VerifyClientCertIfGiven":                 tls.VerifyClientCertIfGiven,
	"VersionSSL30":                            tls.VersionSSL30,
	"VersionTLS10":                            tls.VersionTLS10,
	"VersionTLS11":                            tls.VersionTLS11,
	"VersionTLS12":                            tls.VersionTLS12,

	"listen":      tls.Listen,
	"newListener": tls.NewListener,

	"newLRUClientSessionCache": tls.NewLRUClientSessionCache,

	"Certificate":        qlang.StructOf((*tls.Certificate)(nil)),
	"loadX509KeyPair":    tls.LoadX509KeyPair,
	"X509KeyPair":        tls.X509KeyPair,
	"ClientHelloInfo":    qlang.StructOf((*tls.ClientHelloInfo)(nil)),
	"ClientSessionState": qlang.StructOf((*tls.ClientSessionState)(nil)),
	"Config":             qlang.StructOf((*tls.Config)(nil)),
	"Conn":               qlang.StructOf((*tls.Conn)(nil)),
	"client":             tls.Client,
	"dial":               tls.Dial,
	"dialWithDialer":     tls.DialWithDialer,
	"server":             tls.Server,
	"ConnectionState":    qlang.StructOf((*tls.ConnectionState)(nil)),
}
