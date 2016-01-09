package main

import (
	l "development/modules/UITest"
	e "development/modules/email"
	i "development/modules/other"
	ban "development/modules/users"
	"fmt"
	c "github.com/skilstak/go/colors"
	u "github.com/whitman-colm/go-1/utils/other"
	get "os/exec"
	usr "os/user"
	"strings"
)

func main() {
	usrdata, err := usr.Current()
	u.QuitAtError(err)
	user := string(usrdata.Username)
	get.Command("go get development/modules/users")
	//u.QuitAtError(err)
	l.Output()
	fmt.Println()
	done := false
	isBanned := ban.TestForBannedUser(user)
	if isBanned != false {
		fmt.Println(c.CL + c.B01 + "Sorry, but you have been banned from using S3 for some reason or another...")
		u.Go(1)
		u.Bye()
	}
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
