package cwrapper

// #cgo CFLAGS: -g -Wall
// #include <stdlib.h>
// #include "ssh_c_lib.h"
import "C"
import (
	"fmt"
	"unsafe"
)

// LCD has a fixed width and height and adjustable left and right margins
type LCD struct {
	i2cDev      string
	cols        int
	rows        int
	leftMargin  int
	rightMargin int
}

// NewLCD returns a new LCD struct with default values
func NewLCD(i2cDev string) LCD {
	return LCD{
		i2cDev:      i2cDev,
		cols:        16,
		rows:        8,
		leftMargin:  1,
		rightMargin: 0,
	}
}

// LCDInit initializes the LCD
func (l LCD) LCDInit() error {
	i2cpath := C.CString(l.i2cDev)
	defer C.free(unsafe.Pointer(i2cpath))
	if C.lcd_init(i2cpath) != 0 {
		return fmt.Errorf("could not initiale LCD, non-zero return code from C.lcd_init()")
	}

	return nil
}

// PrintAtPos prints a value at a position using the C.lcd_printc function
func (l LCD) PrintAtRowCol(value rune, row, col int) error {
	var position int
	if row > l.rows {
		return fmt.Errorf("row value %d is greater than LCD rows %d", row, l.rows)
	} else if col > (l.cols - (l.leftMargin + l.rightMargin)) {
		return fmt.Errorf("col value %d is greater than LCD cols %d minus margins (%d, %d)",
			col, l.cols, l.leftMargin, l.rightMargin)
	} else {
		// compute the actual position
		position = (row-1)*l.cols + col - l.leftMargin
	}

	if C.lcd_printc(C.char(value), C.int(position)) != 0 {
		return fmt.Errorf("Error printing value %c at position %d", value, position)
	}
	return nil
}

// Clear empties the screen with C.lcd_clear()
func (l LCD) Clear() error {
	if C.lcd_clear() != 0 {
		return fmt.Errorf("call to C.lcd_clear returned non-zero code")
	}
	return nil
}
