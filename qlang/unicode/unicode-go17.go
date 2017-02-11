// +build go1.7

package unicode

import "unicode"

func init() {
	Exports["Adlam"] = unicode.Adlam
	Exports["Bhaiksuki"] = unicode.Bhaiksuki
	Exports["Marchen"] = unicode.Marchen
	Exports["Newa"] = unicode.Newa
	Exports["Osage"] = unicode.Osage
	Exports["Prepended_Concatenation_Mark"] = unicode.Prepended_Concatenation_Mark
	Exports["Sentence_Terminal"] = unicode.Sentence_Terminal
	Exports["Tangut"] = unicode.Tangut
}
