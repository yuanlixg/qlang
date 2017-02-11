// +build go1.6

package asn1

import "encoding/asn1"

func init() {
	Exports["ClassApplication"] = asn1.ClassApplication
	Exports["ClassContextSpecific"] = asn1.ClassContextSpecific
	Exports["ClassPrivate"] = asn1.ClassPrivate
	Exports["ClassUniversal"] = asn1.ClassUniversal
	Exports["TagBitString"] = asn1.TagBitString
	Exports["TagBoolean"] = asn1.TagBoolean
	Exports["TagEnum"] = asn1.TagEnum
	Exports["TagGeneralString"] = asn1.TagGeneralString
	Exports["TagGeneralizedTime"] = asn1.TagGeneralizedTime
	Exports["TagIA5String"] = asn1.TagIA5String
	Exports["TagInteger"] = asn1.TagInteger
	Exports["TagOID"] = asn1.TagOID
	Exports["TagOctetString"] = asn1.TagOctetString
	Exports["TagPrintableString"] = asn1.TagPrintableString
	Exports["TagSequence"] = asn1.TagSequence
	Exports["TagSet"] = asn1.TagSet
	Exports["TagT61String"] = asn1.TagT61String
	Exports["TagUTCTime"] = asn1.TagUTCTime
	Exports["TagUTF8String"] = asn1.TagUTF8String
}
