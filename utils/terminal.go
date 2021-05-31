package utils

import (
	"log"
	"os"
	"os/exec"
	"runtime"
)

func DebugErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func CheckErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func CallClear() {
	clear := make(map[string]func()) //Initialize it
	clear["darwin"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	//if we defined a clear func for that platform:
	if ok {
		value() //we execute it
	} else {
		panic("Your platform is unsupported! I can't clear terminal screen :(") //unsupported platform
	}
}

func ResizeWindow() {
	resize := make(map[string]func())
	resize["darwin"] = func() {
		cmd := exec.Command("resize -s 35 120")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	resize["linux"] = func() {
		cmd := exec.Command("resize -s 35 120")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	resize["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "mode con:cols=130 lines=35") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	value, ok := resize[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	//if we defined a clear func for that platform:
	if ok {
		value() //we execute it
	} else {
		panic("Your platform is unsupported! I can't resize terminal screen :(") //unsupported platform
	}
}
