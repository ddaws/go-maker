package maker

import (
	"math"
	"math/big"
)

// https://github.com/makerdao/dss/blob/master/DEVELOPING.md#units
const (
	// WadScale is the fixed-point precision of a wad unit (for basic quantities, e.g balances)
	WadScale = 18
	// RayScale is the fixed-point precision of the ray unit (for precise quantities, e.g. ratios)
	RayScale = 27
	// RadScale is the fixed-point precision of the rad unit (result of integer multiplication of a wad and ray)
	RadScale = 45
)

var (
	// One unit of Ray
	One = big.NewInt(int64(math.Pow(10, RayScale)))
)
