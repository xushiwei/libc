package libc

var C0_cgo18___cosdf float64 = -0.499999997251031
var C1_cgo19___cosdf float64 = 0.041666623323739063
var C2_cgo20___cosdf float64 = -0.0013886763774609929
var C3_cgo21___cosdf float64 = 2.4390448796277409e-5

func __cosdf(x float64) float32 {
	var r float64
	var w float64
	var z float64
	z = x * x
	w = z * z
	r = C2_cgo20___cosdf + z*C3_cgo21___cosdf
	return float32(1 + z*C0_cgo18___cosdf + w*C1_cgo19___cosdf + w*z*r)
}