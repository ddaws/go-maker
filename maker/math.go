package maker

import "math/big"

// https://github.com/makerdao/dss/blob/master/DEVELOPING.md#units
const (
	// WadScale is the fixed-point precision of a wad unit (for basic quantities, e.g balances)
	WadScale = 18
	// RayScale is the fixed-point precision of the ray unit (for precise quantities, e.g. ratios)
	RayScale = 27
	// RadScale is the fixed-point precision of the rad unit (result of integer multiplication of a wad and ray)
	RadScale = 45
)

type Ray struct {
	*big.Int
}

type Wad struct {
	*big.Int
}

type Rad struct {
	*big.Int
}

// Mul implements standard integer multiplication of fixed-point precision integers in units of wad.
func (w *Wad) Mul(r *Ray) *Rad {
	return nil
}

// RMul implements fixed-point precision multiplication of a Wad and a Ray where precision is lost.
func (w *Wad) RMul(r *Ray) *Wad {
	return nil
}

// RMul implements fixed-point precision multiplication of a Ray and a Ray where precision is lost.
func (r *Ray) RMul(ray *Ray) *Ray {
	return nil
}

// RMul implements fixed-point precision multiplication of a Rad and a Ray where precision is lost.
func (r *Rad) RMul(ray *Ray) *Rad {
	return nil
}
