package log1pf

import (
	common "github.com/goplus/libc/test/common"
	unsafe "unsafe"
	libc "github.com/goplus/libc"
	testing "testing"
)

var _cgos_t_log1pf [19]common.Struct_f_f = [19]common.Struct_f_f{common.Struct_f_f{(*int8)(unsafe.Pointer(&[53]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'l', 'o', 'g', '1', 'p', 'f', '.', 'h', '\x00'})), int32(1), int32(0), float32(-8.0668487548828125), libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'}))), float32(0.0), int32(0)}, common.Struct_f_f{(*int8)(unsafe.Pointer(&[53]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'l', 'o', 'g', '1', 'p', 'f', '.', 'h', '\x00'})), int32(2), int32(0), float32(4.3452396392822266), float32(1.6762063503265381), float32(-0.23014859855175018), int32(0)}, common.Struct_f_f{(*int8)(unsafe.Pointer(&[53]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'l', 'o', 'g', '1', 'p', 'f', '.', 'h', '\x00'})), int32(3), int32(0), float32(-8.3814334869384765), libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'}))), float32(0.0), int32(0)}, common.Struct_f_f{(*int8)(unsafe.Pointer(&[53]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'l', 'o', 'g', '1', 'p', 'f', '.', 'h', '\x00'})), int32(4), int32(0), float32(-6.5316734313964844), libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'}))), float32(0.0), int32(0)}, common.Struct_f_f{(*int8)(unsafe.Pointer(&[53]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'l', 'o', 'g', '1', 'p', 'f', '.', 'h', '\x00'})), int32(5), int32(0), float32(9.2670574188232421), float32(2.3289403915405273), float32(-0.29075589776039124), int32(0)}, common.Struct_f_f{(*int8)(unsafe.Pointer(&[53]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'l', 'o', 'g', '1', 'p', 'f', '.', 'h', '\x00'})), int32(6), int32(0), float32(0.66198587417602539), float32(0.50801318883895874), float32(-0.13867668807506561), int32(0)}, common.Struct_f_f{(*int8)(unsafe.Pointer(&[53]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'l', 'o', 'g', '1', 'p', 'f', '.', 'h', '\x00'})), int32(7), int32(0), float32(-0.40660393238067627), float32(-0.5218932032585144), float32(-0.088044330477714539), int32(0)}, common.Struct_f_f{(*int8)(unsafe.Pointer(&[53]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'l', 'o', 'g', '1', 'p', 'f', '.', 'h', '\x00'})), int32(8), int32(0), float32(0.56175976991653442), float32(0.44581323862075806), float32(-0.15101368725299835), int32(0)}, common.Struct_f_f{(*int8)(unsafe.Pointer(&[53]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'l', 'o', 'g', '1', 'p', 'f', '.', 'h', '\x00'})), int32(9), int32(0), float32(0.77415227890014648), float32(0.57332271337509155), float32(-0.10264533013105392), int32(0)}, common.Struct_f_f{(*int8)(unsafe.Pointer(&[53]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'a', 'n', 'i', 't', 'y', '/', 'l', 'o', 'g', '1', 'p', 'f', '.', 'h', '\x00'})), int32(10), int32(0), float32(-0.67876368761062622), float32(-1.1355782747268677), float32(-0.19879481196403503), int32(0)}, common.Struct_f_f{(*int8)(unsafe.Pointer(&[54]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'l', 'o', 'g', '1', 'p', 'f', '.', 'h', '\x00'})), int32(1), int32(0), float32(0.0), float32(0.0), float32(0.0), int32(0)}, common.Struct_f_f{(*int8)(unsafe.Pointer(&[54]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'l', 'o', 'g', '1', 'p', 'f', '.', 'h', '\x00'})), int32(2), int32(0), float32(-0.0), float32(-0.0), float32(0.0), int32(0)}, common.Struct_f_f{(*int8)(unsafe.Pointer(&[54]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'l', 'o', 'g', '1', 'p', 'f', '.', 'h', '\x00'})), int32(3), int32(0), float32(-7.8886090522101181e-31), float32(-7.8886090522101181e-31), float32(3.3087224502121107e-24), int32(0)}, common.Struct_f_f{(*int8)(unsafe.Pointer(&[54]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'l', 'o', 'g', '1', 'p', 'f', '.', 'h', '\x00'})), int32(4), int32(0), float32(1.0), float32(0.69314718246459961), float32(0.031954795122146606), int32(0)}, common.Struct_f_f{(*int8)(unsafe.Pointer(&[54]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'l', 'o', 'g', '1', 'p', 'f', '.', 'h', '\x00'})), int32(5), int32(0), float32(-1.0), -libc.X__builtin_inff(), float32(0.0), int32(0)}, common.Struct_f_f{(*int8)(unsafe.Pointer(&[54]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'l', 'o', 'g', '1', 'p', 'f', '.', 'h', '\x00'})), int32(6), int32(0), libc.X__builtin_inff(), libc.X__builtin_inff(), float32(0.0), int32(0)}, common.Struct_f_f{(*int8)(unsafe.Pointer(&[54]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'l', 'o', 'g', '1', 'p', 'f', '.', 'h', '\x00'})), int32(7), int32(0), -libc.X__builtin_inff(), libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'}))), float32(0.0), int32(0)}, common.Struct_f_f{(*int8)(unsafe.Pointer(&[54]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'l', 'o', 'g', '1', 'p', 'f', '.', 'h', '\x00'})), int32(8), int32(0), libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'}))), libc.X__builtin_nanf((*int8)(unsafe.Pointer(&[1]int8{'\x00'}))), float32(0.0), int32(0)}, common.Struct_f_f{(*int8)(unsafe.Pointer(&[54]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'm', 'a', 't', 'h', '/', 's', 'p', 'e', 'c', 'i', 'a', 'l', '/', 'l', 'o', 'g', '1', 'p', 'f', '.', 'h', '\x00'})), int32(9), int32(0), float32(-1.1754942106924411e-38), float32(-1.1754942106924411e-38), float32(4.9303806576313238e-32), 0}}

func _cgo_main() int32 {
	var y float32
	var d float32
	var e int32
	var i int32
	var err int32 = int32(0)
	var p *common.Struct_f_f
	for i = int32(0); uint64(i) < 19; i++ {
		p = (*common.Struct_f_f)(unsafe.Pointer(uintptr(unsafe.Pointer((*common.Struct_f_f)(unsafe.Pointer(&_cgos_t_log1pf)))) + uintptr(i)*32))
		if p.R < int32(0) {
			continue
		}
		libc.Fesetround(p.R)
		libc.Feclearexcept(int32(0))
		y = libc.Log1pf(p.X)
		e = libc.Fetestexcept(0)
		if !(common.Checkexcept(e, p.E, p.R) != 0) {
			libc.Printf((*int8)(unsafe.Pointer(&[51]int8{'%', 's', ':', '%', 'd', ':', ' ', 'b', 'a', 'd', ' ', 'f', 'p', ' ', 'e', 'x', 'c', 'e', 'p', 't', 'i', 'o', 'n', ':', ' ', '%', 's', ' ', 'l', 'o', 'g', '1', 'p', 'f', '(', '%', 'a', ')', '=', '%', 'a', ',', ' ', 'w', 'a', 'n', 't', ' ', '%', 's', '\x00'})), p.File, p.Line, common.Rstr(p.R), float64(p.X), float64(p.Y), common.Estr(p.E))
			libc.Printf((*int8)(unsafe.Pointer(&[9]int8{' ', 'g', 'o', 't', ' ', '%', 's', '\n', '\x00'})), common.Estr(e))
			err++
		}
		d = common.Ulperrf(y, p.Y, p.Dy)
		if !(common.Checkulp(d, p.R) != 0) {
			libc.Printf((*int8)(unsafe.Pointer(&[59]int8{'%', 's', ':', '%', 'd', ':', ' ', '%', 's', ' ', 'l', 'o', 'g', '1', 'p', 'f', '(', '%', 'a', ')', ' ', 'w', 'a', 'n', 't', ' ', '%', 'a', ' ', 'g', 'o', 't', ' ', '%', 'a', ' ', 'u', 'l', 'p', 'e', 'r', 'r', ' ', '%', '.', '3', 'f', ' ', '=', ' ', '%', 'a', ' ', '+', ' ', '%', 'a', '\n', '\x00'})), p.File, p.Line, common.Rstr(p.R), float64(p.X), float64(p.Y), float64(y), float64(d), float64(d-p.Dy), float64(p.Dy))
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
