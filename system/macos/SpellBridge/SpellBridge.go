package spellbridge

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#include <stdlib.h>

const char* CheckSpelling(const char* word);
*/
import (
	"C"
	"unsafe"
)

func SpellCheck(word string) (string, bool) {
	cword := C.CString(word)
	defer C.free(unsafe.Pointer(cword))

	result := C.CheckSpelling(cword)
	if result == nil {
		return "", false // No Error
	}

	suggestion := C.GoString(result)
	C.free(unsafe.Pointer(result))

	return suggestion, true
}
