package libc

import unsafe "unsafe"

func __fpclassify(x float64) int32 {
	type _cgoa_118 struct {
		f float64
	}
	var u _cgoa_118
	u.f = x
	var e int32 = int32(*(*uint64)(unsafe.Pointer(&u)) >> int32(52) & uint64(2047))
	if !(e != 0) {
		return func() int32 {
			if *(*uint64)(unsafe.Pointer(&u))<<int32(1) != 0 {
				return int32(3)
			} else {
				return int32(2)
			}
		}()
	}
	if e == int32(2047) {
		return func() int32 {
			if *(*uint64)(unsafe.Pointer(&u))<<int32(12) != 0 {
				return int32(0)
			} else {
				return int32(1)
			}
		}()
	}
	return int32(4)
}
