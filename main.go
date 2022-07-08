package main

import (
	"time"

	"github.com/samiam2013/ssh1106/cwrapper"
)

func main() {
	message := "Hello, world!"
	lcd := cwrapper.NewLCD("/dev/i2c-1")
	lcd.LCDInit()
	for i := 0; i < len(message); i++ {
		lcd.PrintAtRowCol(rune(message[i]), 1, i+1)
	}
	time.Sleep(time.Second * 5)
	lcd.Clear()
}
