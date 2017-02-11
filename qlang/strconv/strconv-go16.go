// +build go1.6

package strconv

import "strconv"

func init() {
	Exports["appendQuoteRuneToGraphic"] = strconv.AppendQuoteRuneToGraphic
	Exports["appendQuoteToGraphic"] = strconv.AppendQuoteToGraphic
	Exports["isGraphic"] = strconv.IsGraphic
	Exports["quoteRuneToGraphic"] = strconv.QuoteRuneToGraphic
	Exports["quoteToGraphic"] = strconv.QuoteToGraphic
}
