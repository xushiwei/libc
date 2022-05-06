package libc

import unsafe "unsafe"

var dummy_file_cgo6 *struct__IO_FILE = nil

func close_file(f *struct__IO_FILE) {
	if !(f != nil) {
		return
	}
	func() int32 {
		if f.lock >= 0 {
			return __lockfile(f)
		} else {
			return 0
		}
	}()
	if uintptr(unsafe.Pointer(f.wpos)) != uintptr(unsafe.Pointer(f.wbase)) {
		f.write(f, nil, uint64(0))
	}
	if uintptr(unsafe.Pointer(f.rpos)) != uintptr(unsafe.Pointer(f.rend)) {
		f.seek(f, int64(uintptr(unsafe.Pointer(f.rpos))-uintptr(unsafe.Pointer(f.rend))), 1)
	}
}
func __stdio_exit() {
	var f *struct__IO_FILE
	for f = *__ofl_lock(); f != nil; f = f.next {
		close_file(f)
	}
	close_file(__stdin_used)
	close_file(__stdout_used)
	close_file(__stderr_used)
}