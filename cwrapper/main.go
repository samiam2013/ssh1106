package cwrapper

// #cgo CFLAGS: -g -Wall
// #include <stdlib.h>
// #include "ssh_c_lib.h"
import "C"
import (
	"fmt"
	"unsafe"
)

func PrintA() {
	i2cDev := C.CString("/dev/i2c-1")
	defer C.free(unsafe.Pointer(i2cDev))
	initRetCode := C.lcd_init(i2cDev)
	fmt.Println("init return code: ", initRetCode)
	C.lcd_clear();
	var charVal rune = 'A';
	var position int = 2;
	cCharVal := C.char(charVal);
	cPosVal := C.int(position);
	retCode := C.lcd_printc(cCharVal, cPosVal);
	fmt.Println("return code: ", retCode)
	C.lcd_clear();
}
