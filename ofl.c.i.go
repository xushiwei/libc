package libc

import unsafe "unsafe"

var ofl_head_cgo21 *struct__IO_FILE
var ofl_lock_cgo22 [1]int32
var __stdio_ofl_lockptr *int32 = (*int32)(unsafe.Pointer(&ofl_lock_cgo22))

func __ofl_lock() **struct__IO_FILE {
	__lock((*int32)(unsafe.Pointer(&ofl_lock_cgo22)))
	return &ofl_head_cgo21
}
func __ofl_unlock() {
	__unlock((*int32)(unsafe.Pointer(&ofl_lock_cgo22)))
}