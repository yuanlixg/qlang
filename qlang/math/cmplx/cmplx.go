package cmplx

import (
	"math/cmplx"
)

// Exports is the export table of this module.
//
var Exports = map[string]interface{}{
	"_name": "math/cmplx",

	"abs":   cmplx.Abs,
	"acos":  cmplx.Acos,
	"acosh": cmplx.Acosh,
	"asin":  cmplx.Asin,
	"asinh": cmplx.Asinh,
	"atan":  cmplx.Atan,
	"atanh": cmplx.Atanh,
	"conj":  cmplx.Conj,
	"cos":   cmplx.Cos,
	"cosh":  cmplx.Cosh,
	"cot":   cmplx.Cot,
	"exp":   cmplx.Exp,
	"inf":   cmplx.Inf,
	"isInf": cmplx.IsInf,
	"isNaN": cmplx.IsNaN,
	"log":   cmplx.Log,
	"log10": cmplx.Log10,
	"naN":   cmplx.NaN,
	"phase": cmplx.Phase,
	"polar": cmplx.Polar,
	"pow":   cmplx.Pow,
	"rect":  cmplx.Rect,
	"sin":   cmplx.Sin,
	"sinh":  cmplx.Sinh,
	"sqrt":  cmplx.Sqrt,
	"tan":   cmplx.Tan,
	"tanh":  cmplx.Tanh,
}
