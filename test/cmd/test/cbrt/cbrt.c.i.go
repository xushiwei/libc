package cbrt

import (
	common "github.com/goplus/libc/test/common"
	unsafe "unsafe"
	libc "github.com/goplus/libc"
	testing "testing"
)

var _cgos_t_cbrt [20]common.Struct_d_d = [20]common.Struct_d_d{common.Struct_d_d{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'c', 'b', 'r', 't', '.', 'h', '\x00'})), int32(1), int32(0), -8.0668483905796808, -2.0055552545020245, float32(0.46667951345443726), int32(0)}, common.Struct_d_d{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'c', 'b', 'r', 't', '.', 'h', '\x00'})), int32(2), int32(0), 4.3452398493383049, 1.6318162410515635, float32(-0.081602714955806732), int32(0)}, common.Struct_d_d{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'c', 'b', 'r', 't', '.', 'h', '\x00'})), int32(3), int32(0), -8.3814334275552493, -2.031293910673361, float32(-0.048101816326379776), int32(0)}, common.Struct_d_d{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'c', 'b', 'r', 't', '.', 'h', '\x00'})), int32(4), int32(0), -6.5316735819134841, -1.8692820012204925, float32(0.086240187287330627), int32(0)}, common.Struct_d_d{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'c', 'b', 'r', 't', '.', 'h', '\x00'})), int32(5), int32(0), 9.2670569669725857, 2.1004577208597022, float32(-0.2722989022731781), int32(0)}, common.Struct_d_d{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'c', 'b', 'r', 't', '.', 'h', '\x00'})), int32(6), int32(0), 0.66198589809950448, 0.87153114704559731, float32(0.44149181246757507), int32(0)}, common.Struct_d_d{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'c', 'b', 'r', 't', '.', 'h', '\x00'})), int32(7), int32(0), -0.40660392238535531, -0.74083903030022302, float32(0.016453813761472702), int32(0)}, common.Struct_d_d{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'c', 'b', 'r', 't', '.', 'h', '\x00'})), int32(8), int32(0), 0.56175974622072411, 0.82511954005592858, float32(0.30680638551712036), int32(0)}, common.Struct_d_d{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'c', 'b', 'r', 't', '.', 'h', '\x00'})), int32(9), int32(0), 0.77415229659130369, 0.91821024789599137, float32(0.065439984202384949), int32(0)}, common.Struct_d_d{(*int8)(unsafe.Pointer(&[51]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'c', 'b', 'r', 't', '.', 'h', '\x00'})), int32(10), int32(0), -0.67876370263940244, -0.8788326906580094, float32(-0.20167131721973419), int32(0)}, common.Struct_d_d{(*int8)(unsafe.Pointer(&[52]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'c', 'b', 'r', 't', '.', 'h', '\x00'})), int32(1), int32(0), float64(libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'})))), float64(libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'})))), float32(0.0), int32(0)}, common.Struct_d_d{(*int8)(unsafe.Pointer(&[52]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'c', 'b', 'r', 't', '.', 'h', '\x00'})), int32(2), int32(0), float64(libc.X__builtin_inff()), float64(libc.X__builtin_inff()), float32(0.0), int32(0)}, common.Struct_d_d{(*int8)(unsafe.Pointer(&[52]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'c', 'b', 'r', 't', '.', 'h', '\x00'})), int32(3), int32(0), float64(-libc.X__builtin_inff()), float64(-libc.X__builtin_inff()), float32(0.0), int32(0)}, common.Struct_d_d{(*int8)(unsafe.Pointer(&[52]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'c', 'b', 'r', 't', '.', 'h', '\x00'})), int32(4), int32(0), 0.0, 0.0, float32(0.0), int32(0)}, common.Struct_d_d{(*int8)(unsafe.Pointer(&[52]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'c', 'b', 'r', 't', '.', 'h', '\x00'})), int32(5), int32(0), -0.0, -0.0, float32(0.0), int32(0)}, common.Struct_d_d{(*int8)(unsafe.Pointer(&[52]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'c', 'b', 'r', 't', '.', 'h', '\x00'})), int32(6), int32(0), 9.3132257461547852e-10, 9.765625e-4, float32(0.0), int32(0)}, common.Struct_d_d{(*int8)(unsafe.Pointer(&[52]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'c', 'b', 'r', 't', '.', 'h', '\x00'})), int32(7), int32(0), -9.3132257461547852e-10, -9.765625e-4, float32(0.0), int32(0)}, common.Struct_d_d{(*int8)(unsafe.Pointer(&[52]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'c', 'b', 'r', 't', '.', 'h', '\x00'})), int32(8), int32(0), 1.0, 1.0, float32(0.0), int32(0)}, common.Struct_d_d{(*int8)(unsafe.Pointer(&[52]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'c', 'b', 'r', 't', '.', 'h', '\x00'})), int32(9), int32(0), -1.0, -1.0, float32(0.0), int32(0)}, common.Struct_d_d{(*int8)(unsafe.Pointer(&[52]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'c', 'b', 'r', 't', '.', 'h', '\x00'})), int32(10), int32(0), 8.0, 2.0, float32(0.0), int32(0)}}

func _cgo_main() int32 {
	var y float64
	var d float32
	var e int32
	var i int32
	var err int32 = int32(0)
	var p *common.Struct_d_d
	for i = int32(0); uint64(i) < 20; i++ {
		p = (*common.Struct_d_d)(unsafe.Pointer(uintptr(unsafe.Pointer((*common.Struct_d_d)(unsafe.Pointer(&_cgos_t_cbrt)))) + uintptr(i)*40))
		if p.R < int32(0) {
			continue
		}
		libc.Fesetround(p.R)
		libc.Feclearexcept(int32(0))
		y = libc.Cbrt(p.X)
		e = libc.Fetestexcept(0)
		if !(common.Checkexcept(e, p.E, p.R) != 0) {
			libc.Printf((*int8)(unsafe.Pointer(&[49]int8{'%', 's', ':', '%', 'd', ':', ' ', 'b', 'a', 'd', ' ', 'f', 'p', ' ', 'e', 'x', 'c', 'e', 'p', 't', 'i', 'o', 'n', ':', ' ', '%', 's', ' ', 'c', 'b', 'r', 't', '(', '%', 'a', ')', '=', '%', 'a', ',', ' ', 'w', 'a', 'n', 't', ' ', '%', 's', '\x00'})), p.File, p.Line, common.Rstr(p.R), p.X, p.Y, common.Estr(p.E))
			libc.Printf((*int8)(unsafe.Pointer(&[9]int8{' ', 'g', 'o', 't', ' ', '%', 's', '\n', '\x00'})), common.Estr(e))
			err++
		}
		d = common.Ulperr(y, p.Y, p.Dy)
		if !(common.Checkulp(d, p.R) != 0) {
			libc.Printf((*int8)(unsafe.Pointer(&[57]int8{'%', 's', ':', '%', 'd', ':', ' ', '%', 's', ' ', 'c', 'b', 'r', 't', '(', '%', 'a', ')', ' ', 'w', 'a', 'n', 't', ' ', '%', 'a', ' ', 'g', 'o', 't', ' ', '%', 'a', ' ', 'u', 'l', 'p', 'e', 'r', 'r', ' ', '%', '.', '3', 'f', ' ', '=', ' ', '%', 'a', ' ', '+', ' ', '%', 'a', '\n', '\x00'})), p.File, p.Line, common.Rstr(p.R), p.X, p.Y, y, float64(d), float64(d-p.Dy), float64(p.Dy))
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
