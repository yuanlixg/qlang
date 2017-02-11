package strconv

import (
	"strconv"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "strconv",

	"IntSize": strconv.IntSize,

	"ErrRange":  strconv.ErrRange,
	"ErrSyntax": strconv.ErrSyntax,

	"appendBool":             strconv.AppendBool,
	"appendFloat":            strconv.AppendFloat,
	"appendInt":              strconv.AppendInt,
	"appendQuote":            strconv.AppendQuote,
	"appendQuoteRune":        strconv.AppendQuoteRune,
	"appendQuoteRuneToASCII": strconv.AppendQuoteRuneToASCII,
	"appendQuoteToASCII":     strconv.AppendQuoteToASCII,
	"appendUint":             strconv.AppendUint,
	"atoi":                   strconv.Atoi,
	"canBackquote":           strconv.CanBackquote,
	"formatBool":             strconv.FormatBool,
	"formatFloat":            strconv.FormatFloat,
	"formatInt":              strconv.FormatInt,
	"formatUint":             strconv.FormatUint,
	"isPrint":                strconv.IsPrint,
	"itoa":                   strconv.Itoa,
	"parseBool":              strconv.ParseBool,
	"parseFloat":             strconv.ParseFloat,
	"parseInt":               strconv.ParseInt,
	"parseUint":              strconv.ParseUint,
	"quote":                  strconv.Quote,
	"quoteRune":              strconv.QuoteRune,
	"quoteRuneToASCII":       strconv.QuoteRuneToASCII,
	"quoteToASCII":           strconv.QuoteToASCII,
	"unquote":                strconv.Unquote,
	"unquoteChar":            strconv.UnquoteChar,
}
