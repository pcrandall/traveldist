package main

import (
	"fmt"
	"log"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/pcrandall/figlet4go"
)

func checkErr(err error, str string) {
	if err != nil {
		if str != "" {
			log.Println(str, err)
		} else {
			log.Println(err)
		}
	}
}

func callClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func callResize() {
	value, ok := resize[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                           //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func printHeader(str string) {
	callClear()
	t := time.Now()
	y := t.Year()
	y -= 2000 // ill be suprised if im using this in 2100
	padding := ""
	signature := "pcrandall '" + strconv.Itoa(y)
	paddingLen := 0

	ascii := figlet4go.NewAsciiRender()
	// change the font color
	// uncomment to activate colors
	colors := [...]color.Attribute{
		color.FgWhite,
		// color.FgMagenta,
		// color.FgYellow,
		// color.FgCyan,
		// color.FgBlue,
		// color.FgHiGreen,
		// color.FgGreen,
	}
	options := figlet4go.NewRenderOptions()
	options.FontColor = make([]color.Attribute, len(str))
	for i := range options.FontColor {
		options.FontColor[i] = colors[i%len(colors)]

	}

	// you can add more fonts like this if you want ascii.LoadFont("./fonts/bigMoneyNE.flf")
	renderStr, _ := ascii.RenderOpts(str, options)

	// calculate the correct padding for the signature 11 is the font height
	var last, longestRow int
	for i := 0; i < len(renderStr)-1; i++ {
		if renderStr[i] == 10 {
			curlongest := i - last
			last = i
			if curlongest > longestRow {
				longestRow = curlongest
			}
		}
	}

	//check if even or odd, add some more padding
	if longestRow%2 == 1 {
		longestRow /= 2
	} else {
		longestRow = (longestRow / 2) + 4
	}

	// TODO fix the padding calc
	// paddingLen = len(renderStr)/(11*2) - len(signature)
	paddingLen = longestRow - len(signature)
	for i := 0; i <= paddingLen; i++ {
		padding += " "
	}
	// remove the last three blank rows, all uppercase chars have a height of 8, the font height for default font is 11
	fmt.Println(renderStr[:len(renderStr)-len(renderStr)/11*3-1])
	// print the signature
	fmt.Printf("%s%s\n", padding, signature)
}

func cleanString(str string) string {
	// return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
	return strings.TrimSpace(str)
}
