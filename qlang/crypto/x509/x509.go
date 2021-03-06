package x509

import (
	"crypto/x509"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "crypto/x509",

	"CANotAuthorizedForThisName": x509.CANotAuthorizedForThisName,
	"DSA":                                   x509.DSA,
	"DSAWithSHA1":                           x509.DSAWithSHA1,
	"DSAWithSHA256":                         x509.DSAWithSHA256,
	"ECDSA":                                 x509.ECDSA,
	"ECDSAWithSHA1":                         x509.ECDSAWithSHA1,
	"ECDSAWithSHA256":                       x509.ECDSAWithSHA256,
	"ECDSAWithSHA384":                       x509.ECDSAWithSHA384,
	"ECDSAWithSHA512":                       x509.ECDSAWithSHA512,
	"Expired":                               x509.Expired,
	"ExtKeyUsageAny":                        x509.ExtKeyUsageAny,
	"ExtKeyUsageClientAuth":                 x509.ExtKeyUsageClientAuth,
	"ExtKeyUsageCodeSigning":                x509.ExtKeyUsageCodeSigning,
	"ExtKeyUsageEmailProtection":            x509.ExtKeyUsageEmailProtection,
	"ExtKeyUsageIPSECEndSystem":             x509.ExtKeyUsageIPSECEndSystem,
	"ExtKeyUsageIPSECTunnel":                x509.ExtKeyUsageIPSECTunnel,
	"ExtKeyUsageIPSECUser":                  x509.ExtKeyUsageIPSECUser,
	"ExtKeyUsageMicrosoftServerGatedCrypto": x509.ExtKeyUsageMicrosoftServerGatedCrypto,
	"ExtKeyUsageNetscapeServerGatedCrypto":  x509.ExtKeyUsageNetscapeServerGatedCrypto,
	"ExtKeyUsageOCSPSigning":                x509.ExtKeyUsageOCSPSigning,
	"ExtKeyUsageServerAuth":                 x509.ExtKeyUsageServerAuth,
	"ExtKeyUsageTimeStamping":               x509.ExtKeyUsageTimeStamping,
	"IncompatibleUsage":                     x509.IncompatibleUsage,
	"KeyUsageCRLSign":                       x509.KeyUsageCRLSign,
	"KeyUsageCertSign":                      x509.KeyUsageCertSign,
	"KeyUsageContentCommitment":             x509.KeyUsageContentCommitment,
	"KeyUsageDataEncipherment":              x509.KeyUsageDataEncipherment,
	"KeyUsageDecipherOnly":                  x509.KeyUsageDecipherOnly,
	"KeyUsageDigitalSignature":              x509.KeyUsageDigitalSignature,
	"KeyUsageEncipherOnly":                  x509.KeyUsageEncipherOnly,
	"KeyUsageKeyAgreement":                  x509.KeyUsageKeyAgreement,
	"KeyUsageKeyEncipherment":               x509.KeyUsageKeyEncipherment,
	"MD2WithRSA":                            x509.MD2WithRSA,
	"MD5WithRSA":                            x509.MD5WithRSA,
	"NotAuthorizedToSign":                   x509.NotAuthorizedToSign,
	"PEMCipher3DES":                         x509.PEMCipher3DES,
	"PEMCipherAES128":                       x509.PEMCipherAES128,
	"PEMCipherAES192":                       x509.PEMCipherAES192,
	"PEMCipherAES256":                       x509.PEMCipherAES256,
	"PEMCipherDES":                          x509.PEMCipherDES,
	"RSA":                                   x509.RSA,
	"SHA1WithRSA":                           x509.SHA1WithRSA,
	"SHA256WithRSA":                         x509.SHA256WithRSA,
	"SHA384WithRSA":                         x509.SHA384WithRSA,
	"SHA512WithRSA":                         x509.SHA512WithRSA,
	"TooManyIntermediates":                  x509.TooManyIntermediates,
	"UnknownPublicKeyAlgorithm":             x509.UnknownPublicKeyAlgorithm,
	"UnknownSignatureAlgorithm":             x509.UnknownSignatureAlgorithm,

	"ErrUnsupportedAlgorithm": x509.ErrUnsupportedAlgorithm,
	"IncorrectPasswordError":  x509.IncorrectPasswordError,

	"createCertificate":        x509.CreateCertificate,
	"createCertificateRequest": x509.CreateCertificateRequest,
	"decryptPEMBlock":          x509.DecryptPEMBlock,
	"encryptPEMBlock":          x509.EncryptPEMBlock,
	"isEncryptedPEMBlock":      x509.IsEncryptedPEMBlock,
	"marshalECPrivateKey":      x509.MarshalECPrivateKey,
	"marshalPKCS1PrivateKey":   x509.MarshalPKCS1PrivateKey,
	"marshalPKIXPublicKey":     x509.MarshalPKIXPublicKey,
	"parseCRL":                 x509.ParseCRL,
	"parseCertificates":        x509.ParseCertificates,
	"parseDERCRL":              x509.ParseDERCRL,
	"parseECPrivateKey":        x509.ParseECPrivateKey,
	"parsePKCS1PrivateKey":     x509.ParsePKCS1PrivateKey,
	"parsePKCS8PrivateKey":     x509.ParsePKCS8PrivateKey,
	"parsePKIXPublicKey":       x509.ParsePKIXPublicKey,

	"CertPool":                   qlang.StructOf((*x509.CertPool)(nil)),
	"certPool":                   x509.NewCertPool,
	"Certificate":                qlang.StructOf((*x509.Certificate)(nil)),
	"parseCertificate":           x509.ParseCertificate,
	"CertificateRequest":         qlang.StructOf((*x509.CertificateRequest)(nil)),
	"parseCertificateRequest":    x509.ParseCertificateRequest,
	"UnhandledCriticalExtension": qlang.StructOf((*x509.UnhandledCriticalExtension)(nil)),
	"VerifyOptions":              qlang.StructOf((*x509.VerifyOptions)(nil)),
}
