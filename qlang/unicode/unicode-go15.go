// +build go1.5

package unicode

import "unicode"

func init() {
	Exports["Ahom"] = unicode.Ahom
	Exports["Anatolian_Hieroglyphs"] = unicode.Anatolian_Hieroglyphs
	Exports["Hatran"] = unicode.Hatran
	Exports["Multani"] = unicode.Multani
	Exports["Old_Hungarian"] = unicode.Old_Hungarian
	Exports["SignWriting"] = unicode.SignWriting
}
