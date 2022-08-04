package scalblnf

import (
	common "github.com/goplus/libc/test/common"
	unsafe "unsafe"
	libc "github.com/goplus/libc"
	testing "testing"
)

var _cgos_t_scalblnf [27]common.Struct_fi_f = [27]common.Struct_fi_f{common.Struct_fi_f{(*int8)(unsafe.Pointer(&[58]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 's', 'c', 'a', 'l', 'b', 'l', 'n', 'f', '.', 'h', '\x00'})), int32(1), int32(0), float32(-8.0668487548828125), int64(-2), float32(-2.0167121887207031), float32(0), int32(0)}, common.Struct_fi_f{(*int8)(unsafe.Pointer(&[58]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 's', 'c', 'a', 'l', 'b', 'l', 'n', 'f', '.', 'h', '\x00'})), int32(2), int32(0), float32(4.3452396392822266), int64(-1), float32(2.1726198196411133), float32(0), int32(0)}, common.Struct_fi_f{(*int8)(unsafe.Pointer(&[58]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 's', 'c', 'a', 'l', 'b', 'l', 'n', 'f', '.', 'h', '\x00'})), int32(3), int32(0), float32(-8.3814334869384765), int64(0), float32(-8.3814334869384765), float32(0), int32(0)}, common.Struct_fi_f{(*int8)(unsafe.Pointer(&[58]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 's', 'c', 'a', 'l', 'b', 'l', 'n', 'f', '.', 'h', '\x00'})), int32(4), int32(0), float32(-6.5316734313964844), int64(1), float32(-13.063346862792969), float32(0), int32(0)}, common.Struct_fi_f{(*int8)(unsafe.Pointer(&[58]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 's', 'c', 'a', 'l', 'b', 'l', 'n', 'f', '.', 'h', '\x00'})), int32(5), int32(0), float32(9.2670574188232421), int64(2), float32(37.068229675292969), float32(0), int32(0)}, common.Struct_fi_f{(*int8)(unsafe.Pointer(&[58]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 's', 'c', 'a', 'l', 'b', 'l', 'n', 'f', '.', 'h', '\x00'})), int32(6), int32(0), float32(0.66198587417602539), int64(3), float32(5.2958869934082031), float32(0), int32(0)}, common.Struct_fi_f{(*int8)(unsafe.Pointer(&[58]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 's', 'c', 'a', 'l', 'b', 'l', 'n', 'f', '.', 'h', '\x00'})), int32(7), int32(0), float32(-0.40660393238067627), int64(4), float32(-6.5056629180908203), float32(0), int32(0)}, common.Struct_fi_f{(*int8)(unsafe.Pointer(&[58]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 's', 'c', 'a', 'l', 'b', 'l', 'n', 'f', '.', 'h', '\x00'})), int32(8), int32(0), float32(0.56175976991653442), int64(5), float32(17.976312637329102), float32(0), int32(0)}, common.Struct_fi_f{(*int8)(unsafe.Pointer(&[58]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 's', 'c', 'a', 'l', 'b', 'l', 'n', 'f', '.', 'h', '\x00'})), int32(9), int32(0), float32(0.77415227890014648), int64(6), float32(49.545745849609375), float32(0), int32(0)}, common.Struct_fi_f{(*int8)(unsafe.Pointer(&[58]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 's', 'c', 'a', 'l', 'b', 'l', 'n', 'f', '.', 'h', '\x00'})), int32(10), int32(0), float32(-0.67876368761062622), int64(7), float32(-86.881752014160156), float32(0), int32(0)}, common.Struct_fi_f{(*int8)(unsafe.Pointer(&[59]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 's', 'c', 'a', 'l', 'b', 'l', 'n', 'f', '.', 'h', '\x00'})), int32(1), int32(0), float32(0), int64(2147483647), float32(0), float32(0), int32(0)}, common.Struct_fi_f{(*int8)(unsafe.Pointer(&[59]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 's', 'c', 'a', 'l', 'b', 'l', 'n', 'f', '.', 'h', '\x00'})), int32(2), int32(0), float32(0), int64(-2147483647), float32(0), float32(0), int32(0)}, common.Struct_fi_f{(*int8)(unsafe.Pointer(&[59]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 's', 'c', 'a', 'l', 'b', 'l', 'n', 'f', '.', 'h', '\x00'})), int32(3), int32(0), float32(-0), int64(2147483647), float32(-0), float32(0), int32(0)}, common.Struct_fi_f{(*int8)(unsafe.Pointer(&[59]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 's', 'c', 'a', 'l', 'b', 'l', 'n', 'f', '.', 'h', '\x00'})), int32(4), int32(0), libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'}))), int64(0), libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'}))), float32(0), int32(0)}, common.Struct_fi_f{(*int8)(unsafe.Pointer(&[59]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 's', 'c', 'a', 'l', 'b', 'l', 'n', 'f', '.', 'h', '\x00'})), int32(5), int32(0), libc.X__builtin_inff(), int64(0), libc.X__builtin_inff(), float32(0), int32(0)}, common.Struct_fi_f{(*int8)(unsafe.Pointer(&[59]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 's', 'c', 'a', 'l', 'b', 'l', 'n', 'f', '.', 'h', '\x00'})), int32(6), int32(0), -libc.X__builtin_inff(), int64(0), -libc.X__builtin_inff(), float32(0), int32(0)}, common.Struct_fi_f{(*int8)(unsafe.Pointer(&[59]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 's', 'c', 'a', 'l', 'b', 'l', 'n', 'f', '.', 'h', '\x00'})), int32(7), int32(0), float32(1), int64(0), float32(1), float32(0), int32(0)}, common.Struct_fi_f{(*int8)(unsafe.Pointer(&[59]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 's', 'c', 'a', 'l', 'b', 'l', 'n', 'f', '.', 'h', '\x00'})), int32(8), int32(0), float32(1), int64(1), float32(2), float32(0), int32(0)}, common.Struct_fi_f{(*int8)(unsafe.Pointer(&[59]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 's', 'c', 'a', 'l', 'b', 'l', 'n', 'f', '.', 'h', '\x00'})), int32(9), int32(0), float32(1), int64(-1), float32(0.5), float32(0), int32(0)}, common.Struct_fi_f{(*int8)(unsafe.Pointer(&[59]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 's', 'c', 'a', 'l', 'b', 'l', 'n', 'f', '.', 'h', '\x00'})), int32(10), int32(0), float32(1), int64(2147483647), libc.X__builtin_inff(), float32(0), 0}, common.Struct_fi_f{(*int8)(unsafe.Pointer(&[59]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 's', 'c', 'a', 'l', 'b', 'l', 'n', 'f', '.', 'h', '\x00'})), int32(11), int32(0), libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'}))), int64(1), libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'}))), float32(0), int32(0)}, common.Struct_fi_f{(*int8)(unsafe.Pointer(&[59]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 's', 'c', 'a', 'l', 'b', 'l', 'n', 'f', '.', 'h', '\x00'})), int32(12), int32(0), libc.X__builtin_inff(), int64(2147483647), libc.X__builtin_inff(), float32(0), int32(0)}, common.Struct_fi_f{(*int8)(unsafe.Pointer(&[59]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 's', 'c', 'a', 'l', 'b', 'l', 'n', 'f', '.', 'h', '\x00'})), int32(13), int32(0), libc.X__builtin_inff(), int64(-2147483647), libc.X__builtin_inff(), float32(0), int32(0)}, common.Struct_fi_f{(*int8)(unsafe.Pointer(&[59]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 's', 'c', 'a', 'l', 'b', 'l', 'n', 'f', '.', 'h', '\x00'})), int32(14), int32(0), -libc.X__builtin_inff(), int64(2147483647), -libc.X__builtin_inff(), float32(0), int32(0)}, common.Struct_fi_f{(*int8)(unsafe.Pointer(&[59]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 's', 'c', 'a', 'l', 'b', 'l', 'n', 'f', '.', 'h', '\x00'})), int32(15), int32(0), float32(1.7014118346046923e+38), int64(-276), float32(1.4012984643248171e-45), float32(0), int32(0)}, common.Struct_fi_f{(*int8)(unsafe.Pointer(&[59]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 's', 'c', 'a', 'l', 'b', 'l', 'n', 'f', '.', 'h', '\x00'})), int32(16), int32(0), float32(1.4012984643248171e-45), int64(276), float32(1.7014118346046923e+38), float32(0), int32(0)}, common.Struct_fi_f{(*int8)(unsafe.Pointer(&[59]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'x', 'u', 's', 'h', 'i', 'w', 'e', 'i', '/', 'w', 'o', 'r', 'k', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 's', 'c', 'a', 'l', 'b', 'l', 'n', 'f', '.', 'h', '\x00'})), int32(17), int32(0), float32(1.000244140625), int64(-149), float32(1.4012984643248171e-45), float32(0), 0}}

func _cgo_main() int32 {
	var y float64
	var d float32
	var e int32
	var i int32
	var err int32 = int32(0)
	var p *common.Struct_fi_f
	for i = int32(0); uint64(i) < 29; i++ {
		p = (*common.Struct_fi_f)(unsafe.Pointer(uintptr(unsafe.Pointer((*common.Struct_fi_f)(unsafe.Pointer(&_cgos_t_scalblnf)))) + uintptr(i)*44))
		if p.R < int32(0) {
			continue
		}
		libc.Fesetround(p.R)
		libc.Feclearexcept(int32(0))
		y = float64(libc.Scalblnf(p.X, int64(p.I)))
		e = libc.Fetestexcept(0)
		if !(common.Checkexceptall(e, p.E, p.R) != 0) {
			libc.Printf((*int8)(unsafe.Pointer(&[59]int8{'%', 's', ':', '%', 'd', ':', ' ', 'b', 'a', 'd', ' ', 'f', 'p', ' ', 'e', 'x', 'c', 'e', 'p', 't', 'i', 'o', 'n', ':', ' ', '%', 's', ' ', 's', 'c', 'a', 'l', 'b', 'l', 'n', 'f', '(', '%', 'a', ',', ' ', '%', 'l', 'l', 'd', ')', '=', '%', 'a', ',', ' ', 'w', 'a', 'n', 't', ' ', '%', 's', '\x00'})), p.File, p.Line, common.Rstr(p.R), float64(p.X), p.I, float64(p.Y), common.Estr(p.E))
			libc.Printf((*int8)(unsafe.Pointer(&[9]int8{' ', 'g', 'o', 't', ' ', '%', 's', '\n', '\x00'})), common.Estr(e))
			err++
		}
		d = common.Ulperrf(float32(y), p.Y, p.Dy)
		if !(common.Checkcr(float64(y), float64(p.Y), p.R) != 0) {
			libc.Printf((*int8)(unsafe.Pointer(&[68]int8{'%', 's', ':', '%', 'd', ':', ' ', '%', 's', ' ', 's', 'c', 'a', 'l', 'b', 'l', 'n', 'f', '(', '%', 'a', ',', ' ', '%', 'l', 'l', 'd', ')', ' ', 'w', 'a', 'n', 't', ' ', '%', 'a', ' ', 'g', 'o', 't', ' ', '%', 'a', ',', ' ', 'u', 'l', 'p', 'e', 'r', 'r', ' ', '%', '.', '3', 'f', ' ', '=', ' ', '%', 'a', ' ', '+', ' ', '%', 'a', '\n', '\x00'})), p.File, p.Line, common.Rstr(p.R), float64(p.X), p.I, float64(p.Y), y, float64(d), float64(d-p.Dy), float64(p.Dy))
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