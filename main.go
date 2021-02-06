package main

import (
	"Osiris-pwm/gui"
	_ "fmt"
	"os"
)

func main() {
	//create the data and the gui folders if they don't already exist
	os.Mkdir("data", 0777)
	os.Mkdir("gui", 0777)
	//call the main gui and run it
	gui.Call()
}
