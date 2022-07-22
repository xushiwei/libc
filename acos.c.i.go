package libc

import unsafe "unsafe"

var pio2_hi_cgo18_acos float64 = 1.5707963267948966
var pio2_lo_cgo19_acos float64 = 6.123233995736766e-17
var pS0_cgo20_acos float64 = 0.16666666666666666
var pS1_cgo21_acos float64 = -0.32556581862240092
var pS2_cgo22_acos float64 = 0.20121253213486293
var pS3_cgo23_acos float64 = -0.040055534500679411
var pS4_cgo24_acos float64 = 7.9153499428981453e-4
var pS5_cgo25_acos float64 = 3.4793310759602117e-5
var qS1_cgo26_acos float64 = -2.4033949117344142
var qS2_cgo27_acos float64 = 2.0209457602335057
var qS3_cgo28_acos float64 = -0.68828397160545329
var qS4_cgo29_acos float64 = 0.077038150555901935

func R_cgo30_acos(z float64) float64 {
	var p float64
	var q float64
	p = z * (pS0_cgo20_acos + z*(pS1_cgo21_acos+z*(pS2_cgo22_acos+z*(pS3_cgo23_acos+z*(pS4_cgo24_acos+z*pS5_cgo25_acos)))))
	q = 1 + z*(qS1_cgo26_acos+z*(qS2_cgo27_acos+z*(qS3_cgo28_acos+z*qS4_cgo29_acos)))
	return p / q
}
func Acos(x float64) float64 {
	var z float64
	var w float64
	var s float64
	var c float64
	var df float64
	var hx uint32
	var ix uint32
	for {
		hx = uint32(*(*uint64)(unsafe.Pointer(&_cgoz_31_acos{x})) >> int32(32))
		if true {
			break
		}
	}
	ix = hx & uint32(2147483647)
	if ix >= uint32(1072693248) {
		var lx uint32
		for {
			lx = uint32(*(*uint64)(unsafe.Pointer(&_cgoz_32_acos{x})))
			if true {
				break
			}
		}
		if ix-uint32(1072693248)|lx == uint32(0) {
			if hx>>int32(31) != 0 {
				return float64(int32(2))*pio2_hi_cgo18_acos + float64(7.52316385e-37)
			}
			return float64(int32(0))
		}
		return float64(int32(0)) / (x - x)
	}
	if ix < uint32(1071644672) {
		if ix <= uint32(1012924416) {
			return pio2_hi_cgo18_acos + float64(7.52316385e-37)
		}
		return pio2_hi_cgo18_acos - (x - (pio2_lo_cgo19_acos - x*R_cgo30_acos(x*x)))
	}
	if hx>>int32(31) != 0 {
		z = (1 + x) * 0.5
		s = Sqrt(z)
		w = R_cgo30_acos(z)*s - pio2_lo_cgo19_acos
		return float64(int32(2)) * (pio2_hi_cgo18_acos - (s + w))
	}
	z = (1 - x) * 0.5
	s = Sqrt(z)
	df = s
	for {
		df = *(*float64)(unsafe.Pointer(&_cgoz_33_acos{uint64(*(*uint64)(unsafe.Pointer(&_cgoz_34_acos{df}))>>int32(32))<<int32(32) | uint64(0)}))
		if true {
			break
		}
	}
	c = (z - df*df) / (s + df)
	w = R_cgo30_acos(z)*s + c
	return float64(int32(2)) * (df + w)
}

type _cgoz_31_acos struct {
	_f float64
}
type _cgoz_32_acos struct {
	_f float64
}
type _cgoz_33_acos struct {
	_i uint64
}
type _cgoz_34_acos struct {
	_f float64
}