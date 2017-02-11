package pkix

import (
	"crypto/x509/pkix"

	qlang "qlang.io/spec"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "crypto/x509/pkix",

	"AlgorithmIdentifier":      qlang.StructOf((*pkix.AlgorithmIdentifier)(nil)),
	"AttributeTypeAndValue":    qlang.StructOf((*pkix.AttributeTypeAndValue)(nil)),
	"AttributeTypeAndValueSET": qlang.StructOf((*pkix.AttributeTypeAndValueSET)(nil)),
	"CertificateList":          qlang.StructOf((*pkix.CertificateList)(nil)),
	"Extension":                qlang.StructOf((*pkix.Extension)(nil)),
	"Name":                     qlang.StructOf((*pkix.Name)(nil)),
	"RevokedCertificate":       qlang.StructOf((*pkix.RevokedCertificate)(nil)),
	"TBSCertificateList":       qlang.StructOf((*pkix.TBSCertificateList)(nil)),
}
