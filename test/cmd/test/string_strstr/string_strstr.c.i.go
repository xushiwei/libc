package string_strstr

import (
	unsafe "unsafe"
	libc "github.com/goplus/libc"
	common "github.com/goplus/libc/test/common"
	testing "testing"
)

func _cgo_main() int32 {
	{
		var p *int8 = (*int8)(unsafe.Pointer(&[1]int8{'\x00'}))
		var q *int8 = libc.Strstr(p, (*int8)(unsafe.Pointer(&[2]int8{'a', '\x00'})))
		if q != nil {
			common.T_printf((*int8)(unsafe.Pointer(&[104]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '2', '2', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', '0', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[3]int8{'"', '"', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'"', 'a', '"', '\x00'})), uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)))
		}
	}
	{
		var p *int8 = (*int8)(unsafe.Pointer(&[2]int8{'a', '\x00'}))
		var q *int8 = libc.Strstr(p, (*int8)(unsafe.Pointer(&[3]int8{'a', 'a', '\x00'})))
		if q != nil {
			common.T_printf((*int8)(unsafe.Pointer(&[104]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '2', '3', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', '0', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'"', 'a', '"', '\x00'})), (*int8)(unsafe.Pointer(&[5]int8{'"', 'a', 'a', '"', '\x00'})), uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)))
		}
	}
	{
		var p *int8 = (*int8)(unsafe.Pointer(&[2]int8{'a', '\x00'}))
		var q *int8 = libc.Strstr(p, (*int8)(unsafe.Pointer(&[2]int8{'b', '\x00'})))
		if q != nil {
			common.T_printf((*int8)(unsafe.Pointer(&[104]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '2', '4', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', '0', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'"', 'a', '"', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'"', 'b', '"', '\x00'})), uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)))
		}
	}
	{
		var p *int8 = (*int8)(unsafe.Pointer(&[3]int8{'a', 'a', '\x00'}))
		var q *int8 = libc.Strstr(p, (*int8)(unsafe.Pointer(&[3]int8{'a', 'b', '\x00'})))
		if q != nil {
			common.T_printf((*int8)(unsafe.Pointer(&[104]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '2', '5', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', '0', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[5]int8{'"', 'a', 'a', '"', '\x00'})), (*int8)(unsafe.Pointer(&[5]int8{'"', 'a', 'b', '"', '\x00'})), uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)))
		}
	}
	{
		var p *int8 = (*int8)(unsafe.Pointer(&[3]int8{'a', 'a', '\x00'}))
		var q *int8 = libc.Strstr(p, (*int8)(unsafe.Pointer(&[4]int8{'a', 'a', 'a', '\x00'})))
		if q != nil {
			common.T_printf((*int8)(unsafe.Pointer(&[104]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '2', '6', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', '0', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[5]int8{'"', 'a', 'a', '"', '\x00'})), (*int8)(unsafe.Pointer(&[6]int8{'"', 'a', 'a', 'a', '"', '\x00'})), uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)))
		}
	}
	{
		var p *int8 = (*int8)(unsafe.Pointer(&[5]int8{'a', 'b', 'b', 'a', '\x00'}))
		var q *int8 = libc.Strstr(p, (*int8)(unsafe.Pointer(&[4]int8{'a', 'b', 'a', '\x00'})))
		if q != nil {
			common.T_printf((*int8)(unsafe.Pointer(&[104]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '2', '7', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', '0', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[7]int8{'"', 'a', 'b', 'b', 'a', '"', '\x00'})), (*int8)(unsafe.Pointer(&[6]int8{'"', 'a', 'b', 'a', '"', '\x00'})), uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)))
		}
	}
	{
		var p *int8 = (*int8)(unsafe.Pointer(&[8]int8{'a', 'b', 'c', ' ', 'a', 'b', 'c', '\x00'}))
		var q *int8 = libc.Strstr(p, (*int8)(unsafe.Pointer(&[5]int8{'a', 'b', 'c', 'd', '\x00'})))
		if q != nil {
			common.T_printf((*int8)(unsafe.Pointer(&[104]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '2', '8', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', '0', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[10]int8{'"', 'a', 'b', 'c', ' ', 'a', 'b', 'c', '"', '\x00'})), (*int8)(unsafe.Pointer(&[7]int8{'"', 'a', 'b', 'c', 'd', '"', '\x00'})), uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)))
		}
	}
	{
		var p *int8 = (*int8)(unsafe.Pointer(&[20]int8{'0', '-', '1', '-', '2', '-', '3', '-', '4', '-', '5', '-', '6', '-', '7', '-', '8', '-', '9', '\x00'}))
		var q *int8 = libc.Strstr(p, (*int8)(unsafe.Pointer(&[13]int8{'-', '3', '-', '4', '-', '5', '6', '-', '7', '-', '8', '-', '\x00'})))
		if q != nil {
			common.T_printf((*int8)(unsafe.Pointer(&[104]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '2', '9', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', '0', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[22]int8{'"', '0', '-', '1', '-', '2', '-', '3', '-', '4', '-', '5', '-', '6', '-', '7', '-', '8', '-', '9', '"', '\x00'})), (*int8)(unsafe.Pointer(&[15]int8{'"', '-', '3', '-', '4', '-', '5', '6', '-', '7', '-', '8', '-', '"', '\x00'})), uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)))
		}
	}
	{
		var p *int8 = (*int8)(unsafe.Pointer(&[20]int8{'0', '-', '1', '-', '2', '-', '3', '-', '4', '-', '5', '-', '6', '-', '7', '-', '8', '-', '9', '\x00'}))
		var q *int8 = libc.Strstr(p, (*int8)(unsafe.Pointer(&[14]int8{'-', '3', '-', '4', '-', '5', '+', '6', '-', '7', '-', '8', '-', '\x00'})))
		if q != nil {
			common.T_printf((*int8)(unsafe.Pointer(&[104]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '3', '0', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', '0', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[22]int8{'"', '0', '-', '1', '-', '2', '-', '3', '-', '4', '-', '5', '-', '6', '-', '7', '-', '8', '-', '9', '"', '\x00'})), (*int8)(unsafe.Pointer(&[16]int8{'"', '-', '3', '-', '4', '-', '5', '+', '6', '-', '7', '-', '8', '-', '"', '\x00'})), uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)))
		}
	}
	{
		var p *int8 = (*int8)(unsafe.Pointer(&[12]int8{'_', ' ', '_', ' ', '_', -1, '_', ' ', '_', ' ', '_', '\x00'}))
		var q *int8 = libc.Strstr(p, (*int8)(unsafe.Pointer(&[4]int8{'_', '\u007f', '_', '\x00'})))
		if q != nil {
			common.T_printf((*int8)(unsafe.Pointer(&[104]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '3', '1', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', '0', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[17]int8{'"', '_', ' ', '_', ' ', '_', '\\', 'x', 'f', 'f', '_', ' ', '_', ' ', '_', '"', '\x00'})), (*int8)(unsafe.Pointer(&[9]int8{'"', '_', '\\', 'x', '7', 'f', '_', '"', '\x00'})), uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)))
		}
	}
	{
		var p *int8 = (*int8)(unsafe.Pointer(&[12]int8{'_', ' ', '_', ' ', '_', '\u007f', '_', ' ', '_', ' ', '_', '\x00'}))
		var q *int8 = libc.Strstr(p, (*int8)(unsafe.Pointer(&[4]int8{'_', -1, '_', '\x00'})))
		if q != nil {
			common.T_printf((*int8)(unsafe.Pointer(&[104]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '3', '2', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', '0', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[17]int8{'"', '_', ' ', '_', ' ', '_', '\\', 'x', '7', 'f', '_', ' ', '_', ' ', '_', '"', '\x00'})), (*int8)(unsafe.Pointer(&[9]int8{'"', '_', '\\', 'x', 'f', 'f', '_', '"', '\x00'})), uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)))
		}
	}
	{
		var p *int8 = (*int8)(unsafe.Pointer(&[1]int8{'\x00'}))
		var q *int8 = libc.Strstr(p, (*int8)(unsafe.Pointer(&[1]int8{'\x00'})))
		if uintptr(unsafe.Pointer(q)) == uintptr(unsafe.Pointer(nil)) {
			common.T_printf((*int8)(unsafe.Pointer(&[104]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '3', '4', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '0', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[3]int8{'"', '"', '\x00'})), (*int8)(unsafe.Pointer(&[3]int8{'"', '"', '\x00'})), int32(0))
		} else if int64(uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p))) != int64(0) {
			common.T_printf((*int8)(unsafe.Pointer(&[109]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '3', '4', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[3]int8{'"', '"', '\x00'})), (*int8)(unsafe.Pointer(&[3]int8{'"', '"', '\x00'})), uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)), int32(0))
		}
	}
	{
		var p *int8 = (*int8)(unsafe.Pointer(&[5]int8{'a', 'b', 'c', 'd', '\x00'}))
		var q *int8 = libc.Strstr(p, (*int8)(unsafe.Pointer(&[1]int8{'\x00'})))
		if uintptr(unsafe.Pointer(q)) == uintptr(unsafe.Pointer(nil)) {
			common.T_printf((*int8)(unsafe.Pointer(&[104]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '3', '5', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '0', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[7]int8{'"', 'a', 'b', 'c', 'd', '"', '\x00'})), (*int8)(unsafe.Pointer(&[3]int8{'"', '"', '\x00'})), int32(0))
		} else if int64(uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p))) != int64(0) {
			common.T_printf((*int8)(unsafe.Pointer(&[109]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '3', '5', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[7]int8{'"', 'a', 'b', 'c', 'd', '"', '\x00'})), (*int8)(unsafe.Pointer(&[3]int8{'"', '"', '\x00'})), uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)), int32(0))
		}
	}
	{
		var p *int8 = (*int8)(unsafe.Pointer(&[5]int8{'a', 'b', 'c', 'd', '\x00'}))
		var q *int8 = libc.Strstr(p, (*int8)(unsafe.Pointer(&[2]int8{'a', '\x00'})))
		if uintptr(unsafe.Pointer(q)) == uintptr(unsafe.Pointer(nil)) {
			common.T_printf((*int8)(unsafe.Pointer(&[104]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '3', '6', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '0', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[7]int8{'"', 'a', 'b', 'c', 'd', '"', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'"', 'a', '"', '\x00'})), int32(0))
		} else if int64(uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p))) != int64(0) {
			common.T_printf((*int8)(unsafe.Pointer(&[109]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '3', '6', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[7]int8{'"', 'a', 'b', 'c', 'd', '"', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'"', 'a', '"', '\x00'})), uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)), int32(0))
		}
	}
	{
		var p *int8 = (*int8)(unsafe.Pointer(&[5]int8{'a', 'b', 'c', 'd', '\x00'}))
		var q *int8 = libc.Strstr(p, (*int8)(unsafe.Pointer(&[2]int8{'b', '\x00'})))
		if uintptr(unsafe.Pointer(q)) == uintptr(unsafe.Pointer(nil)) {
			common.T_printf((*int8)(unsafe.Pointer(&[104]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '3', '7', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '0', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[7]int8{'"', 'a', 'b', 'c', 'd', '"', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'"', 'b', '"', '\x00'})), int32(1))
		} else if int64(uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p))) != int64(1) {
			common.T_printf((*int8)(unsafe.Pointer(&[109]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '3', '7', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[7]int8{'"', 'a', 'b', 'c', 'd', '"', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'"', 'b', '"', '\x00'})), uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)), int32(1))
		}
	}
	{
		var p *int8 = (*int8)(unsafe.Pointer(&[5]int8{'a', 'b', 'c', 'd', '\x00'}))
		var q *int8 = libc.Strstr(p, (*int8)(unsafe.Pointer(&[2]int8{'c', '\x00'})))
		if uintptr(unsafe.Pointer(q)) == uintptr(unsafe.Pointer(nil)) {
			common.T_printf((*int8)(unsafe.Pointer(&[104]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '3', '8', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '0', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[7]int8{'"', 'a', 'b', 'c', 'd', '"', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'"', 'c', '"', '\x00'})), int32(2))
		} else if int64(uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p))) != int64(2) {
			common.T_printf((*int8)(unsafe.Pointer(&[109]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '3', '8', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[7]int8{'"', 'a', 'b', 'c', 'd', '"', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'"', 'c', '"', '\x00'})), uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)), int32(2))
		}
	}
	{
		var p *int8 = (*int8)(unsafe.Pointer(&[5]int8{'a', 'b', 'c', 'd', '\x00'}))
		var q *int8 = libc.Strstr(p, (*int8)(unsafe.Pointer(&[2]int8{'d', '\x00'})))
		if uintptr(unsafe.Pointer(q)) == uintptr(unsafe.Pointer(nil)) {
			common.T_printf((*int8)(unsafe.Pointer(&[104]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '3', '9', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '0', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[7]int8{'"', 'a', 'b', 'c', 'd', '"', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'"', 'd', '"', '\x00'})), int32(3))
		} else if int64(uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p))) != int64(3) {
			common.T_printf((*int8)(unsafe.Pointer(&[109]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '3', '9', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[7]int8{'"', 'a', 'b', 'c', 'd', '"', '\x00'})), (*int8)(unsafe.Pointer(&[4]int8{'"', 'd', '"', '\x00'})), uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)), int32(3))
		}
	}
	{
		var p *int8 = (*int8)(unsafe.Pointer(&[5]int8{'a', 'b', 'c', 'd', '\x00'}))
		var q *int8 = libc.Strstr(p, (*int8)(unsafe.Pointer(&[3]int8{'a', 'b', '\x00'})))
		if uintptr(unsafe.Pointer(q)) == uintptr(unsafe.Pointer(nil)) {
			common.T_printf((*int8)(unsafe.Pointer(&[104]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '4', '0', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '0', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[7]int8{'"', 'a', 'b', 'c', 'd', '"', '\x00'})), (*int8)(unsafe.Pointer(&[5]int8{'"', 'a', 'b', '"', '\x00'})), int32(0))
		} else if int64(uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p))) != int64(0) {
			common.T_printf((*int8)(unsafe.Pointer(&[109]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '4', '0', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[7]int8{'"', 'a', 'b', 'c', 'd', '"', '\x00'})), (*int8)(unsafe.Pointer(&[5]int8{'"', 'a', 'b', '"', '\x00'})), uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)), int32(0))
		}
	}
	{
		var p *int8 = (*int8)(unsafe.Pointer(&[5]int8{'a', 'b', 'c', 'd', '\x00'}))
		var q *int8 = libc.Strstr(p, (*int8)(unsafe.Pointer(&[3]int8{'b', 'c', '\x00'})))
		if uintptr(unsafe.Pointer(q)) == uintptr(unsafe.Pointer(nil)) {
			common.T_printf((*int8)(unsafe.Pointer(&[104]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '4', '1', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '0', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[7]int8{'"', 'a', 'b', 'c', 'd', '"', '\x00'})), (*int8)(unsafe.Pointer(&[5]int8{'"', 'b', 'c', '"', '\x00'})), int32(1))
		} else if int64(uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p))) != int64(1) {
			common.T_printf((*int8)(unsafe.Pointer(&[109]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '4', '1', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[7]int8{'"', 'a', 'b', 'c', 'd', '"', '\x00'})), (*int8)(unsafe.Pointer(&[5]int8{'"', 'b', 'c', '"', '\x00'})), uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)), int32(1))
		}
	}
	{
		var p *int8 = (*int8)(unsafe.Pointer(&[5]int8{'a', 'b', 'c', 'd', '\x00'}))
		var q *int8 = libc.Strstr(p, (*int8)(unsafe.Pointer(&[3]int8{'c', 'd', '\x00'})))
		if uintptr(unsafe.Pointer(q)) == uintptr(unsafe.Pointer(nil)) {
			common.T_printf((*int8)(unsafe.Pointer(&[104]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '4', '2', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '0', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[7]int8{'"', 'a', 'b', 'c', 'd', '"', '\x00'})), (*int8)(unsafe.Pointer(&[5]int8{'"', 'c', 'd', '"', '\x00'})), int32(2))
		} else if int64(uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p))) != int64(2) {
			common.T_printf((*int8)(unsafe.Pointer(&[109]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '4', '2', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[7]int8{'"', 'a', 'b', 'c', 'd', '"', '\x00'})), (*int8)(unsafe.Pointer(&[5]int8{'"', 'c', 'd', '"', '\x00'})), uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)), int32(2))
		}
	}
	{
		var p *int8 = (*int8)(unsafe.Pointer(&[6]int8{'a', 'b', 'a', 'b', 'a', '\x00'}))
		var q *int8 = libc.Strstr(p, (*int8)(unsafe.Pointer(&[5]int8{'b', 'a', 'b', 'a', '\x00'})))
		if uintptr(unsafe.Pointer(q)) == uintptr(unsafe.Pointer(nil)) {
			common.T_printf((*int8)(unsafe.Pointer(&[104]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '4', '3', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '0', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[8]int8{'"', 'a', 'b', 'a', 'b', 'a', '"', '\x00'})), (*int8)(unsafe.Pointer(&[7]int8{'"', 'b', 'a', 'b', 'a', '"', '\x00'})), int32(1))
		} else if int64(uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p))) != int64(1) {
			common.T_printf((*int8)(unsafe.Pointer(&[109]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '4', '3', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[8]int8{'"', 'a', 'b', 'a', 'b', 'a', '"', '\x00'})), (*int8)(unsafe.Pointer(&[7]int8{'"', 'b', 'a', 'b', 'a', '"', '\x00'})), uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)), int32(1))
		}
	}
	{
		var p *int8 = (*int8)(unsafe.Pointer(&[7]int8{'a', 'b', 'a', 'b', 'a', 'b', '\x00'}))
		var q *int8 = libc.Strstr(p, (*int8)(unsafe.Pointer(&[6]int8{'b', 'a', 'b', 'a', 'b', '\x00'})))
		if uintptr(unsafe.Pointer(q)) == uintptr(unsafe.Pointer(nil)) {
			common.T_printf((*int8)(unsafe.Pointer(&[104]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '4', '4', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '0', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[9]int8{'"', 'a', 'b', 'a', 'b', 'a', 'b', '"', '\x00'})), (*int8)(unsafe.Pointer(&[8]int8{'"', 'b', 'a', 'b', 'a', 'b', '"', '\x00'})), int32(1))
		} else if int64(uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p))) != int64(1) {
			common.T_printf((*int8)(unsafe.Pointer(&[109]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '4', '4', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[9]int8{'"', 'a', 'b', 'a', 'b', 'a', 'b', '"', '\x00'})), (*int8)(unsafe.Pointer(&[8]int8{'"', 'b', 'a', 'b', 'a', 'b', '"', '\x00'})), uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)), int32(1))
		}
	}
	{
		var p *int8 = (*int8)(unsafe.Pointer(&[8]int8{'a', 'b', 'a', 'b', 'a', 'b', 'a', '\x00'}))
		var q *int8 = libc.Strstr(p, (*int8)(unsafe.Pointer(&[7]int8{'b', 'a', 'b', 'a', 'b', 'a', '\x00'})))
		if uintptr(unsafe.Pointer(q)) == uintptr(unsafe.Pointer(nil)) {
			common.T_printf((*int8)(unsafe.Pointer(&[104]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '4', '5', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '0', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[10]int8{'"', 'a', 'b', 'a', 'b', 'a', 'b', 'a', '"', '\x00'})), (*int8)(unsafe.Pointer(&[9]int8{'"', 'b', 'a', 'b', 'a', 'b', 'a', '"', '\x00'})), int32(1))
		} else if int64(uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p))) != int64(1) {
			common.T_printf((*int8)(unsafe.Pointer(&[109]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '4', '5', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[10]int8{'"', 'a', 'b', 'a', 'b', 'a', 'b', 'a', '"', '\x00'})), (*int8)(unsafe.Pointer(&[9]int8{'"', 'b', 'a', 'b', 'a', 'b', 'a', '"', '\x00'})), uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)), int32(1))
		}
	}
	{
		var p *int8 = (*int8)(unsafe.Pointer(&[9]int8{'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b', '\x00'}))
		var q *int8 = libc.Strstr(p, (*int8)(unsafe.Pointer(&[8]int8{'b', 'a', 'b', 'a', 'b', 'a', 'b', '\x00'})))
		if uintptr(unsafe.Pointer(q)) == uintptr(unsafe.Pointer(nil)) {
			common.T_printf((*int8)(unsafe.Pointer(&[104]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '4', '6', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '0', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[11]int8{'"', 'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b', '"', '\x00'})), (*int8)(unsafe.Pointer(&[10]int8{'"', 'b', 'a', 'b', 'a', 'b', 'a', 'b', '"', '\x00'})), int32(1))
		} else if int64(uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p))) != int64(1) {
			common.T_printf((*int8)(unsafe.Pointer(&[109]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '4', '6', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[11]int8{'"', 'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b', '"', '\x00'})), (*int8)(unsafe.Pointer(&[10]int8{'"', 'b', 'a', 'b', 'a', 'b', 'a', 'b', '"', '\x00'})), uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)), int32(1))
		}
	}
	{
		var p *int8 = (*int8)(unsafe.Pointer(&[10]int8{'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a', '\x00'}))
		var q *int8 = libc.Strstr(p, (*int8)(unsafe.Pointer(&[9]int8{'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a', '\x00'})))
		if uintptr(unsafe.Pointer(q)) == uintptr(unsafe.Pointer(nil)) {
			common.T_printf((*int8)(unsafe.Pointer(&[104]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '4', '7', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '0', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[12]int8{'"', 'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a', '"', '\x00'})), (*int8)(unsafe.Pointer(&[11]int8{'"', 'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a', '"', '\x00'})), int32(1))
		} else if int64(uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p))) != int64(1) {
			common.T_printf((*int8)(unsafe.Pointer(&[109]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '4', '7', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[12]int8{'"', 'a', 'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a', '"', '\x00'})), (*int8)(unsafe.Pointer(&[11]int8{'"', 'b', 'a', 'b', 'a', 'b', 'a', 'b', 'a', '"', '\x00'})), uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)), int32(1))
		}
	}
	{
		var p *int8 = (*int8)(unsafe.Pointer(&[10]int8{'a', 'b', 'b', 'a', 'b', 'a', 'b', 'a', 'b', '\x00'}))
		var q *int8 = libc.Strstr(p, (*int8)(unsafe.Pointer(&[7]int8{'b', 'a', 'b', 'a', 'b', 'a', '\x00'})))
		if uintptr(unsafe.Pointer(q)) == uintptr(unsafe.Pointer(nil)) {
			common.T_printf((*int8)(unsafe.Pointer(&[104]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '4', '8', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '0', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[12]int8{'"', 'a', 'b', 'b', 'a', 'b', 'a', 'b', 'a', 'b', '"', '\x00'})), (*int8)(unsafe.Pointer(&[9]int8{'"', 'b', 'a', 'b', 'a', 'b', 'a', '"', '\x00'})), int32(2))
		} else if int64(uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p))) != int64(2) {
			common.T_printf((*int8)(unsafe.Pointer(&[109]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '4', '8', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[12]int8{'"', 'a', 'b', 'b', 'a', 'b', 'a', 'b', 'a', 'b', '"', '\x00'})), (*int8)(unsafe.Pointer(&[9]int8{'"', 'b', 'a', 'b', 'a', 'b', 'a', '"', '\x00'})), uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)), int32(2))
		}
	}
	{
		var p *int8 = (*int8)(unsafe.Pointer(&[10]int8{'a', 'b', 'b', 'a', 'b', 'a', 'b', 'a', 'b', '\x00'}))
		var q *int8 = libc.Strstr(p, (*int8)(unsafe.Pointer(&[7]int8{'a', 'b', 'a', 'b', 'a', 'b', '\x00'})))
		if uintptr(unsafe.Pointer(q)) == uintptr(unsafe.Pointer(nil)) {
			common.T_printf((*int8)(unsafe.Pointer(&[104]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '4', '9', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '0', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[12]int8{'"', 'a', 'b', 'b', 'a', 'b', 'a', 'b', 'a', 'b', '"', '\x00'})), (*int8)(unsafe.Pointer(&[9]int8{'"', 'a', 'b', 'a', 'b', 'a', 'b', '"', '\x00'})), int32(3))
		} else if int64(uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p))) != int64(3) {
			common.T_printf((*int8)(unsafe.Pointer(&[109]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '4', '9', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[12]int8{'"', 'a', 'b', 'b', 'a', 'b', 'a', 'b', 'a', 'b', '"', '\x00'})), (*int8)(unsafe.Pointer(&[9]int8{'"', 'a', 'b', 'a', 'b', 'a', 'b', '"', '\x00'})), uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)), int32(3))
		}
	}
	{
		var p *int8 = (*int8)(unsafe.Pointer(&[13]int8{'a', 'b', 'a', 'c', 'a', 'b', 'c', 'a', 'b', 'c', 'a', 'b', '\x00'}))
		var q *int8 = libc.Strstr(p, (*int8)(unsafe.Pointer(&[9]int8{'a', 'b', 'c', 'a', 'b', 'c', 'a', 'b', '\x00'})))
		if uintptr(unsafe.Pointer(q)) == uintptr(unsafe.Pointer(nil)) {
			common.T_printf((*int8)(unsafe.Pointer(&[104]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '5', '0', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '0', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[15]int8{'"', 'a', 'b', 'a', 'c', 'a', 'b', 'c', 'a', 'b', 'c', 'a', 'b', '"', '\x00'})), (*int8)(unsafe.Pointer(&[11]int8{'"', 'a', 'b', 'c', 'a', 'b', 'c', 'a', 'b', '"', '\x00'})), int32(4))
		} else if int64(uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p))) != int64(4) {
			common.T_printf((*int8)(unsafe.Pointer(&[109]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '5', '0', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[15]int8{'"', 'a', 'b', 'a', 'c', 'a', 'b', 'c', 'a', 'b', 'c', 'a', 'b', '"', '\x00'})), (*int8)(unsafe.Pointer(&[11]int8{'"', 'a', 'b', 'c', 'a', 'b', 'c', 'a', 'b', '"', '\x00'})), uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)), int32(4))
		}
	}
	{
		var p *int8 = (*int8)(unsafe.Pointer(&[15]int8{'n', 'a', 'n', 'a', 'b', 'a', 'n', 'a', 'b', 'a', 'n', 'a', 'n', 'a', '\x00'}))
		var q *int8 = libc.Strstr(p, (*int8)(unsafe.Pointer(&[4]int8{'a', 'b', 'a', '\x00'})))
		if uintptr(unsafe.Pointer(q)) == uintptr(unsafe.Pointer(nil)) {
			common.T_printf((*int8)(unsafe.Pointer(&[104]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '5', '1', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '0', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[17]int8{'"', 'n', 'a', 'n', 'a', 'b', 'a', 'n', 'a', 'b', 'a', 'n', 'a', 'n', 'a', '"', '\x00'})), (*int8)(unsafe.Pointer(&[6]int8{'"', 'a', 'b', 'a', '"', '\x00'})), int32(3))
		} else if int64(uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p))) != int64(3) {
			common.T_printf((*int8)(unsafe.Pointer(&[109]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '5', '1', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[17]int8{'"', 'n', 'a', 'n', 'a', 'b', 'a', 'n', 'a', 'b', 'a', 'n', 'a', 'n', 'a', '"', '\x00'})), (*int8)(unsafe.Pointer(&[6]int8{'"', 'a', 'b', 'a', '"', '\x00'})), uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)), int32(3))
		}
	}
	{
		var p *int8 = (*int8)(unsafe.Pointer(&[15]int8{'n', 'a', 'n', 'a', 'b', 'a', 'n', 'a', 'b', 'a', 'n', 'a', 'n', 'a', '\x00'}))
		var q *int8 = libc.Strstr(p, (*int8)(unsafe.Pointer(&[4]int8{'b', 'a', 'n', '\x00'})))
		if uintptr(unsafe.Pointer(q)) == uintptr(unsafe.Pointer(nil)) {
			common.T_printf((*int8)(unsafe.Pointer(&[104]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '5', '2', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '0', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[17]int8{'"', 'n', 'a', 'n', 'a', 'b', 'a', 'n', 'a', 'b', 'a', 'n', 'a', 'n', 'a', '"', '\x00'})), (*int8)(unsafe.Pointer(&[6]int8{'"', 'b', 'a', 'n', '"', '\x00'})), int32(4))
		} else if int64(uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p))) != int64(4) {
			common.T_printf((*int8)(unsafe.Pointer(&[109]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '5', '2', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[17]int8{'"', 'n', 'a', 'n', 'a', 'b', 'a', 'n', 'a', 'b', 'a', 'n', 'a', 'n', 'a', '"', '\x00'})), (*int8)(unsafe.Pointer(&[6]int8{'"', 'b', 'a', 'n', '"', '\x00'})), uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)), int32(4))
		}
	}
	{
		var p *int8 = (*int8)(unsafe.Pointer(&[15]int8{'n', 'a', 'n', 'a', 'b', 'a', 'n', 'a', 'b', 'a', 'n', 'a', 'n', 'a', '\x00'}))
		var q *int8 = libc.Strstr(p, (*int8)(unsafe.Pointer(&[5]int8{'a', 'n', 'a', 'b', '\x00'})))
		if uintptr(unsafe.Pointer(q)) == uintptr(unsafe.Pointer(nil)) {
			common.T_printf((*int8)(unsafe.Pointer(&[104]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '5', '3', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '0', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[17]int8{'"', 'n', 'a', 'n', 'a', 'b', 'a', 'n', 'a', 'b', 'a', 'n', 'a', 'n', 'a', '"', '\x00'})), (*int8)(unsafe.Pointer(&[7]int8{'"', 'a', 'n', 'a', 'b', '"', '\x00'})), int32(1))
		} else if int64(uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p))) != int64(1) {
			common.T_printf((*int8)(unsafe.Pointer(&[109]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '5', '3', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[17]int8{'"', 'n', 'a', 'n', 'a', 'b', 'a', 'n', 'a', 'b', 'a', 'n', 'a', 'n', 'a', '"', '\x00'})), (*int8)(unsafe.Pointer(&[7]int8{'"', 'a', 'n', 'a', 'b', '"', '\x00'})), uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)), int32(1))
		}
	}
	{
		var p *int8 = (*int8)(unsafe.Pointer(&[15]int8{'n', 'a', 'n', 'a', 'b', 'a', 'n', 'a', 'b', 'a', 'n', 'a', 'n', 'a', '\x00'}))
		var q *int8 = libc.Strstr(p, (*int8)(unsafe.Pointer(&[7]int8{'b', 'a', 'n', 'a', 'n', 'a', '\x00'})))
		if uintptr(unsafe.Pointer(q)) == uintptr(unsafe.Pointer(nil)) {
			common.T_printf((*int8)(unsafe.Pointer(&[104]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '5', '4', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '0', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[17]int8{'"', 'n', 'a', 'n', 'a', 'b', 'a', 'n', 'a', 'b', 'a', 'n', 'a', 'n', 'a', '"', '\x00'})), (*int8)(unsafe.Pointer(&[9]int8{'"', 'b', 'a', 'n', 'a', 'n', 'a', '"', '\x00'})), int32(8))
		} else if int64(uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p))) != int64(8) {
			common.T_printf((*int8)(unsafe.Pointer(&[109]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '5', '4', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[17]int8{'"', 'n', 'a', 'n', 'a', 'b', 'a', 'n', 'a', 'b', 'a', 'n', 'a', 'n', 'a', '"', '\x00'})), (*int8)(unsafe.Pointer(&[9]int8{'"', 'b', 'a', 'n', 'a', 'n', 'a', '"', '\x00'})), uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)), int32(8))
		}
	}
	{
		var p *int8 = (*int8)(unsafe.Pointer(&[8]int8{'_', ' ', '_', -1, '_', ' ', '_', '\x00'}))
		var q *int8 = libc.Strstr(p, (*int8)(unsafe.Pointer(&[4]int8{'_', -1, '_', '\x00'})))
		if uintptr(unsafe.Pointer(q)) == uintptr(unsafe.Pointer(nil)) {
			common.T_printf((*int8)(unsafe.Pointer(&[104]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '5', '5', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', '0', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[13]int8{'"', '_', ' ', '_', '\\', 'x', 'f', 'f', '_', ' ', '_', '"', '\x00'})), (*int8)(unsafe.Pointer(&[9]int8{'"', '_', '\\', 'x', 'f', 'f', '_', '"', '\x00'})), int32(2))
		} else if int64(uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p))) != int64(2) {
			common.T_printf((*int8)(unsafe.Pointer(&[109]int8{'/', 'U', 's', 'e', 'r', 's', '/', 'v', 'f', 'c', '/', 'g', 'o', 'p', 'l', 'u', 's', '/', 'l', 'i', 'b', 'c', '/', 't', 'e', 's', 't', '/', 's', 'r', 'c', '/', 'f', 'u', 'n', 'c', 't', 'i', 'o', 'n', 'a', 'l', '/', 's', 't', 'r', 'i', 'n', 'g', '_', 's', 't', 'r', 's', 't', 'r', '.', 'c', ':', '5', '5', ':', ' ', 's', 't', 'r', 's', 't', 'r', '(', '%', 's', ',', '%', 's', ')', ' ', 'r', 'e', 't', 'u', 'r', 'n', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', ',', ' ', 'w', 'a', 'n', 't', 'e', 'd', ' ', 's', 't', 'r', '+', '%', 'd', '\n', '\x00'})), (*int8)(unsafe.Pointer(&[13]int8{'"', '_', ' ', '_', '\\', 'x', 'f', 'f', '_', ' ', '_', '"', '\x00'})), (*int8)(unsafe.Pointer(&[9]int8{'"', '_', '\\', 'x', 'f', 'f', '_', '"', '\x00'})), uintptr(unsafe.Pointer(q))-uintptr(unsafe.Pointer(p)), int32(2))
		}
	}
	return common.T_status
}
func TestMain(t *testing.T) {
	if _cgo_ret := _cgo_main(); _cgo_ret != 0 {
		t.Fatal("exit status", _cgo_ret)
	}
}
