package main

import (
	l "development/modules/UITest"
	e "development/modules/email"
	i "development/modules/other"
	"fmt"
	c "github.com/skilstak/go/colors"
	"strings"
)

func main() {
	l.Output()
	fmt.Println()
	done := false
	for done == false {
		function, _ := i.Input("~", "sss")
		function = strings.ToLower(function)
		switch function {
		case "a":
			e.Startup()
			done = true
		case "b":

			fmt.Println(c.CL + c.B1 + "Copy and paste this into your browser")
			fmt.Println()
			fmt.Println(c.O + "https://sss.surge.sh")
			fmt.Println()
			fmt.Println()
			done = true
		case "c":
			done = true
		default:
			fmt.Println(c.B00 + "Sorry, try entering one of the letters in red")
		}
	}
}
