package libc

import unsafe "unsafe"

var strings_cgo34_strsignal [671]int8 = [671]int8{'U', 'n', 'k', 'n', 'o', 'w', 'n', ' ', 's', 'i', 'g', 'n', 'a', 'l', '\x00', 'H', 'a', 'n', 'g', 'u', 'p', '\x00', 'I', 'n', 't', 'e', 'r', 'r', 'u', 'p', 't', '\x00', 'Q', 'u', 'i', 't', '\x00', 'I', 'l', 'l', 'e', 'g', 'a', 'l', ' ', 'i', 'n', 's', 't', 'r', 'u', 'c', 't', 'i', 'o', 'n', '\x00', 'T', 'r', 'a', 'c', 'e', '/', 'b', 'r', 'e', 'a', 'k', 'p', 'o', 'i', 'n', 't', ' ', 't', 'r', 'a', 'p', '\x00', 'A', 'b', 'o', 'r', 't', 'e', 'd', '\x00', 'B', 'u', 's', ' ', 'e', 'r', 'r', 'o', 'r', '\x00', 'A', 'r', 'i', 't', 'h', 'm', 'e', 't', 'i', 'c', ' ', 'e', 'x', 'c', 'e', 'p', 't', 'i', 'o', 'n', '\x00', 'K', 'i', 'l', 'l', 'e', 'd', '\x00', 'U', 's', 'e', 'r', ' ', 'd', 'e', 'f', 'i', 'n', 'e', 'd', ' ', 's', 'i', 'g', 'n', 'a', 'l', ' ', '1', '\x00', 'S', 'e', 'g', 'm', 'e', 'n', 't', 'a', 't', 'i', 'o', 'n', ' ', 'f', 'a', 'u', 'l', 't', '\x00', 'U', 's', 'e', 'r', ' ', 'd', 'e', 'f', 'i', 'n', 'e', 'd', ' ', 's', 'i', 'g', 'n', 'a', 'l', ' ', '2', '\x00', 'B', 'r', 'o', 'k', 'e', 'n', ' ', 'p', 'i', 'p', 'e', '\x00', 'A', 'l', 'a', 'r', 'm', ' ', 'c', 'l', 'o', 'c', 'k', '\x00', 'T', 'e', 'r', 'm', 'i', 'n', 'a', 't', 'e', 'd', '\x00', 'S', 't', 'a', 'c', 'k', ' ', 'f', 'a', 'u', 'l', 't', '\x00', 'C', 'h', 'i', 'l', 'd', ' ', 'p', 'r', 'o', 'c', 'e', 's', 's', ' ', 's', 't', 'a', 't', 'u', 's', '\x00', 'C', 'o', 'n', 't', 'i', 'n', 'u', 'e', 'd', '\x00', 'S', 't', 'o', 'p', 'p', 'e', 'd', ' ', '(', 's', 'i', 'g', 'n', 'a', 'l', ')', '\x00', 'S', 't', 'o', 'p', 'p', 'e', 'd', '\x00', 'S', 't', 'o', 'p', 'p', 'e', 'd', ' ', '(', 't', 't', 'y', ' ', 'i', 'n', 'p', 'u', 't', ')', '\x00', 'S', 't', 'o', 'p', 'p', 'e', 'd', ' ', '(', 't', 't', 'y', ' ', 'o', 'u', 't', 'p', 'u', 't', ')', '\x00', 'U', 'r', 'g', 'e', 'n', 't', ' ', 'I', '/', 'O', ' ', 'c', 'o', 'n', 'd', 'i', 't', 'i', 'o', 'n', '\x00', 'C', 'P', 'U', ' ', 't', 'i', 'm', 'e', ' ', 'l', 'i', 'm', 'i', 't', ' ', 'e', 'x', 'c', 'e', 'e', 'd', 'e', 'd', '\x00', 'F', 'i', 'l', 'e', ' ', 's', 'i', 'z', 'e', ' ', 'l', 'i', 'm', 'i', 't', ' ', 'e', 'x', 'c', 'e', 'e', 'd', 'e', 'd', '\x00', 'V', 'i', 'r', 't', 'u', 'a', 'l', ' ', 't', 'i', 'm', 'e', 'r', ' ', 'e', 'x', 'p', 'i', 'r', 'e', 'd', '\x00', 'P', 'r', 'o', 'f', 'i', 'l', 'i', 'n', 'g', ' ', 't', 'i', 'm', 'e', 'r', ' ', 'e', 'x', 'p', 'i', 'r', 'e', 'd', '\x00', 'W', 'i', 'n', 'd', 'o', 'w', ' ', 'c', 'h', 'a', 'n', 'g', 'e', 'd', '\x00', 'I', '/', 'O', ' ', 'p', 'o', 's', 's', 'i', 'b', 'l', 'e', '\x00', 'P', 'o', 'w', 'e', 'r', ' ', 'f', 'a', 'i', 'l', 'u', 'r', 'e', '\x00', 'B', 'a', 'd', ' ', 's', 'y', 's', 't', 'e', 'm', ' ', 'c', 'a', 'l', 'l', '\x00', 'R', 'T', '3', '2', '\x00', 'R', 'T', '3', '3', '\x00', 'R', 'T', '3', '4', '\x00', 'R', 'T', '3', '5', '\x00', 'R', 'T', '3', '6', '\x00', 'R', 'T', '3', '7', '\x00', 'R', 'T', '3', '8', '\x00', 'R', 'T', '3', '9', '\x00', 'R', 'T', '4', '0', '\x00', 'R', 'T', '4', '1', '\x00', 'R', 'T', '4', '2', '\x00', 'R', 'T', '4', '3', '\x00', 'R', 'T', '4', '4', '\x00', 'R', 'T', '4', '5', '\x00', 'R', 'T', '4', '6', '\x00', 'R', 'T', '4', '7', '\x00', 'R', 'T', '4', '8', '\x00', 'R', 'T', '4', '9', '\x00', 'R', 'T', '5', '0', '\x00', 'R', 'T', '5', '1', '\x00', 'R', 'T', '5', '2', '\x00', 'R', 'T', '5', '3', '\x00', 'R', 'T', '5', '4', '\x00', 'R', 'T', '5', '5', '\x00', 'R', 'T', '5', '6', '\x00', 'R', 'T', '5', '7', '\x00', 'R', 'T', '5', '8', '\x00', 'R', 'T', '5', '9', '\x00', 'R', 'T', '6', '0', '\x00', 'R', 'T', '6', '1', '\x00', 'R', 'T', '6', '2', '\x00', 'R', 'T', '6', '3', '\x00', 'R', 'T', '6', '4', '\x00'}

func Strsignal(signum int32) *int8 {
	var s *int8 = (*int8)(unsafe.Pointer(&strings_cgo34_strsignal))
	signum = signum
	if uint32(signum)-uint32(1) >= uint32(64) {
		signum = int32(0)
	}
	for ; func() (_cgo_ret int32) {
		_cgo_addr := &signum
		_cgo_ret = *_cgo_addr
		*_cgo_addr--
		return
	}() != 0; *(*uintptr)(unsafe.Pointer(&s))++ {
		for ; *s != 0; *(*uintptr)(unsafe.Pointer(&s))++ {
		}
	}
	return (*int8)(unsafe.Pointer(__lctrans_cur(s)))
}
