// cprint.go
package main

import (
	"C"
	"unsafe"
)

func main() {
	cstr := C.CString("Hello World")
	C.puts(cstr)
	C.free(unsafe.Pointer(cstr))
}
