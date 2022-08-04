package ldexpl

import (
	common "github.com/goplus/libc/test/common"
	unsafe "unsafe"
	libc "github.com/goplus/libc"
	testing "testing"
)

var _cgos_t_ldexpl [24]common.Struct_li_l = [24]common.Struct_li_l{common.Struct_li_l{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'l', 'd', 'e', 'x', 'p', '.', 'h', '\x00'})), int32(1), int32(0), float64(-8.0668483905796808), int64(-2), float64(-2.0167120976449202), float32(0), int32(0)}, common.Struct_li_l{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'l', 'd', 'e', 'x', 'p', '.', 'h', '\x00'})), int32(2), int32(0), float64(4.3452398493383049), int64(-1), float64(2.1726199246691524), float32(0), int32(0)}, common.Struct_li_l{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'l', 'd', 'e', 'x', 'p', '.', 'h', '\x00'})), int32(3), int32(0), float64(-8.3814334275552493), int64(0), float64(-8.3814334275552493), float32(0), int32(0)}, common.Struct_li_l{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'l', 'd', 'e', 'x', 'p', '.', 'h', '\x00'})), int32(4), int32(0), float64(-6.5316735819134841), int64(1), float64(-13.063347163826968), float32(0), int32(0)}, common.Struct_li_l{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'l', 'd', 'e', 'x', 'p', '.', 'h', '\x00'})), int32(5), int32(0), float64(9.2670569669725857), int64(2), float64(37.068227867890343), float32(0), int32(0)}, common.Struct_li_l{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'l', 'd', 'e', 'x', 'p', '.', 'h', '\x00'})), int32(6), int32(0), float64(0.66198589809950448), int64(3), float64(5.2958871847960358), float32(0), int32(0)}, common.Struct_li_l{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'l', 'd', 'e', 'x', 'p', '.', 'h', '\x00'})), int32(7), int32(0), float64(-0.40660392238535531), int64(4), float64(-6.505662758165685), float32(0), int32(0)}, common.Struct_li_l{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'l', 'd', 'e', 'x', 'p', '.', 'h', '\x00'})), int32(8), int32(0), float64(0.56175974622072411), int64(5), float64(17.976311879063172), float32(0), int32(0)}, common.Struct_li_l{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'l', 'd', 'e', 'x', 'p', '.', 'h', '\x00'})), int32(9), int32(0), float64(0.77415229659130369), int64(6), float64(49.545746981843436), float32(0), int32(0)}, common.Struct_li_l{(*int8)(unsafe.Pointer(&[55]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'l', 'd', 'e', 'x', 'p', '.', 'h', '\x00'})), int32(10), int32(0), float64(-0.67876370263940244), int64(7), float64(-86.881753937843513), float32(0), int32(0)}, common.Struct_li_l{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'l', 'd', 'e', 'x', 'p', '.', 'h', '\x00'})), int32(1), int32(0), float64(0), int64(2147483647), float64(0), float32(0), int32(0)}, common.Struct_li_l{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'l', 'd', 'e', 'x', 'p', '.', 'h', '\x00'})), int32(2), int32(0), float64(0), int64(-2147483647), float64(0), float32(0), int32(0)}, common.Struct_li_l{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'l', 'd', 'e', 'x', 'p', '.', 'h', '\x00'})), int32(3), int32(0), float64(-0), int64(2147483647), float64(-0), float32(0), int32(0)}, common.Struct_li_l{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'l', 'd', 'e', 'x', 'p', '.', 'h', '\x00'})), int32(4), int32(0), float64(libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'})))), int64(0), float64(libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'})))), float32(0), int32(0)}, common.Struct_li_l{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'l', 'd', 'e', 'x', 'p', '.', 'h', '\x00'})), int32(5), int32(0), float64(libc.X__builtin_inff()), int64(0), float64(libc.X__builtin_inff()), float32(0), int32(0)}, common.Struct_li_l{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'l', 'd', 'e', 'x', 'p', '.', 'h', '\x00'})), int32(6), int32(0), float64(-libc.X__builtin_inff()), int64(0), float64(-libc.X__builtin_inff()), float32(0), int32(0)}, common.Struct_li_l{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'l', 'd', 'e', 'x', 'p', '.', 'h', '\x00'})), int32(7), int32(0), float64(1), int64(0), float64(1), float32(0), int32(0)}, common.Struct_li_l{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'l', 'd', 'e', 'x', 'p', '.', 'h', '\x00'})), int32(8), int32(0), float64(1), int64(1), float64(2), float32(0), int32(0)}, common.Struct_li_l{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'l', 'd', 'e', 'x', 'p', '.', 'h', '\x00'})), int32(9), int32(0), float64(1), int64(-1), float64(0.5), float32(0), int32(0)}, common.Struct_li_l{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'l', 'd', 'e', 'x', 'p', '.', 'h', '\x00'})), int32(10), int32(0), float64(1), int64(2147483647), float64(libc.X__builtin_inff()), float32(0), 0}, common.Struct_li_l{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'l', 'd', 'e', 'x', 'p', '.', 'h', '\x00'})), int32(11), int32(0), float64(libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'})))), int64(1), float64(libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'})))), float32(0), int32(0)}, common.Struct_li_l{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'l', 'd', 'e', 'x', 'p', '.', 'h', '\x00'})), int32(12), int32(0), float64(libc.X__builtin_inff()), int64(2147483647), float64(libc.X__builtin_inff()), float32(0), int32(0)}, common.Struct_li_l{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'l', 'd', 'e', 'x', 'p', '.', 'h', '\x00'})), int32(13), int32(0), float64(libc.X__builtin_inff()), int64(-2147483647), float64(libc.X__builtin_inff()), float32(0), int32(0)}, common.Struct_li_l{(*int8)(unsafe.Pointer(&[56]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'l', 'd', 'e', 'x', 'p', '.', 'h', '\x00'})), int32(14), int32(0), float64(-libc.X__builtin_inff()), int64(2147483647), float64(-libc.X__builtin_inff()), float32(0), int32(0)}}

func _cgo_main() int32 {
	var y float64
	var d float32
	var e int32
	var i int32
	var err int32 = int32(0)
	var p *common.Struct_li_l
	for i = int32(0); uint64(i) < 24; i++ {
		p = (*common.Struct_li_l)(unsafe.Pointer(uintptr(unsafe.Pointer((*common.Struct_li_l)(unsafe.Pointer(&_cgos_t_ldexpl)))) + uintptr(i)*48))
		if p.R < int32(0) {
			continue
		}
		libc.Fesetround(p.R)
		libc.Feclearexcept(int32(0))
		y = libc.Ldexpl(p.X, int32(p.I))
		e = libc.Fetestexcept(0)
		if !(common.Checkexceptall(e, p.E, p.R) != 0) {
			libc.Printf((*int8)(unsafe.Pointer(&[59]int8{'%', 's', ':', '%', 'd', ':', ' ', 'b', 'a', 'd', ' ', 'f', 'p', ' ', 'e', 'x', 'c', 'e', 'p', 't', 'i', 'o', 'n', ':', ' ', '%', 's', ' ', 'l', 'd', 'e', 'x', 'p', 'l', '(', '%', 'L', 'a', ',', ' ', '%', 'l', 'l', 'd', ')', '=', '%', 'L', 'a', ',', ' ', 'w', 'a', 'n', 't', ' ', '%', 's', '\x00'})), p.File, p.Line, common.Rstr(p.R), p.X, p.I, p.Y, common.Estr(p.E))
			libc.Printf((*int8)(unsafe.Pointer(&[9]int8{' ', 'g', 'o', 't', ' ', '%', 's', '\n', '\x00'})), common.Estr(e))
			err++
		}
		d = common.Ulperrl(y, p.Y, p.Dy)
		if !(common.Checkcr(y, p.Y, p.R) != 0) {
			libc.Printf((*int8)(unsafe.Pointer(&[68]int8{'%', 's', ':', '%', 'd', ':', ' ', '%', 's', ' ', 'l', 'd', 'e', 'x', 'p', 'l', '(', '%', 'L', 'a', ',', ' ', '%', 'l', 'l', 'd', ')', ' ', 'w', 'a', 'n', 't', ' ', '%', 'L', 'a', ' ', 'g', 'o', 't', ' ', '%', 'L', 'a', ' ', 'u', 'l', 'p', 'e', 'r', 'r', ' ', '%', '.', '3', 'f', ' ', '=', ' ', '%', 'a', ' ', '+', ' ', '%', 'a', '\n', '\x00'})), p.File, p.Line, common.Rstr(p.R), p.X, p.I, p.Y, y, float64(d), float64(d-p.Dy), float64(p.Dy))
			err++
		}
	}
	return func() int32 {
		if !!(err != 0) {
			return 1
		} else {
			return 0
		}
	}()
}
func TestMain(t *testing.T) {
	if _cgo_ret := _cgo_main(); _cgo_ret != 0 {
		t.Fatal("exit status", _cgo_ret)
	}
}
