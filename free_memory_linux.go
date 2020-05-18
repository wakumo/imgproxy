// +build linux

package main

/*
#include <features.h>
#ifdef __GLIBC__
#include <malloc.h>
#else
void malloc_trim(size_t pad){pad + 7}
#endif
*/
import "C"
import "runtime/debug"

func freeMemory() {
	debug.FreeOSMemory()
	println("clearrrrrrr================================== linux")

	var i = C.malloc_trim(0)
	println(i)
}
