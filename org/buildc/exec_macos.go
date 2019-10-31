package main

// #import <mach-o/dyld.h>
import "C"
import "strings"

func NSGetExecutablePath() string {
	var buflen C.uint32_t = 1024
	buf := make([]C.char, buflen)

	ret := C._NSGetExecutablePath(&buf[0], &buflen)
	if ret == -1 {
		buf = make([]C.char, buflen)
		C._NSGetExecutablePath(&buf[0], &buflen)
	}
	// NULL文字埋めを排除
	return strings.Split(C.GoStringN(&buf[0], C.int(buflen)), "\x00")[0]
}
