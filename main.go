package main

import (
	"time"

	"github.com/samiam2013/ssh1106/cwrapper"
)

func main() {
	message := "0123456789ABCD"
	lcd := cwrapper.NewLCD("/dev/i2c-1", 0x3c)
	lcd.LCDInit()
	for i := 0; i < 8; i++ {
		for j := 0; j < len(message); j++ {
			lcd.PrintAtRowCol(rune(message[j]), i, j+2)
		}
	}
	time.Sleep(time.Second * 5)
	lcd.Clear()
	lcd.Close()
}
