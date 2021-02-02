package main

import (
	"Osiris-pwm/crypt"
	gui "Osiris-pwm/gui"
	_ "fmt"
)

func main() {
	gui.Call()
	println(crypt.GenerateKey())
}
