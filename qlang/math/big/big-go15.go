// +build go1.5

package big

import (
	"math/big"

	qlang "qlang.io/spec"
)

func init() {
	Exports["Above"] = big.Above
	Exports["AwayFromZero"] = big.AwayFromZero
	Exports["Below"] = big.Below
	Exports["Exact"] = big.Exact
	Exports["MaxExp"] = big.MaxExp
	Exports["MaxPrec"] = uint64(big.MaxPrec)
	Exports["MinExp"] = big.MinExp
	Exports["ToNearestAway"] = big.ToNearestAway
	Exports["ToNearestEven"] = big.ToNearestEven
	Exports["ToNegativeInf"] = big.ToNegativeInf
	Exports["ToPositiveInf"] = big.ToPositiveInf
	Exports["ToZero"] = big.ToZero
	Exports["jacobi"] = big.Jacobi
	Exports["ErrNaN"] = qlang.StructOf((*big.ErrNaN)(nil))
	Exports["Float"] = qlang.StructOf((*big.Float)(nil))
	Exports["float"] = big.NewFloat
	Exports["parseFloat"] = big.ParseFloat
}
