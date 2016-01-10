package nil

import (
	id "development/modules/TimeID"
	e "development/modules/email"
	ban "development/modules/users"
	"fmt"
	c "github.com/skilstak/go/colors"
)

const version string = c.V + "V S-1.2.0"

func Output() {
	fmt.Print(c.Clear)
	fmt.Println(c.B3 + "╔═════════════════════════════════════════════════════╗")
	fmt.Println(c.B3 + "║ " + c.Y + "SkilStak Support System, By: " + c.G + "whitman-colm & tsnlc04 " + c.B3 + "║")
	fmt.Println(c.B3 + "║ " + c.O + "This is run by the following modules:               " + c.B3 + "║")
	fmt.Println(c.B3 + "║                                                     ║")
	fmt.Println(c.B3 + "║ " + c.V + "SSS Main  |  " + c.V + "V S-2.0.2                              " + c.B3 + "║")
	fmt.Println(c.B3 + "║ " + c.V + "SSS Email |  " + e.Version + "                              " + c.B3 + "║")
	fmt.Println(c.B3 + "║ " + c.V + "SSS UUID  |  " + id.Version + "                              " + c.B3 + "║")
	fmt.Println(c.B3 + "║ " + c.V + "SSS UI    |  " + version + "                              " + c.B3 + "║")
	fmt.Println(c.B3 + "║ " + c.V + "SSS Ban   |  " + ban.Version + "                              " + c.B3 + "║")
	fmt.Println(c.B3 + "╚═════════════════════════════════════════════════════╝")
	fmt.Println()
	fmt.Println(c.M + "How may we help you today?")
	fmt.Println(c.R + "[A]" + c.B + " I need programming help!")
	fmt.Println(c.R + "[B]" + c.B + " Check the FAQ.")
	fmt.Println(c.R + "[C]" + c.B + " Just stoppin' by.")
}

func One() {
	fmt.Println(c.B3 + "╔══════════════════════════════════════════════════════════════════════════════════════════════════════════╗")
	fmt.Println(c.B3 + "║ " + c.B00 + "What is the priority of the ticket? You will be banned to prevent spamming." + c.B3 + "                              ║")
	fmt.Println(c.B3 + "║ " + c.B00 + "You will be banned from S3 if you have failed to check the FAQ or do not honestly prioritize the ticket." + c.B3 + " ║")
	fmt.Println(c.B3 + "║                                                                                                          ║")
	fmt.Println(c.B3 + "║ " + c.G + "Low Priority    " + c.X + "┃┃ " + c.B1 + " No ban" + c.X + "                ┃┃  " + c.B1 + "Example: I have a small bug in my program." + c.B3 + "                 ║")
	fmt.Println(c.B3 + "║ " + c.Y + "Medium Priority " + c.X + "┃┃ " + c.B1 + " One day ban" + c.X + "           ┃┃  " + c.B1 + "Example: My side project can't be setup." + c.B3 + "                   ║")
	fmt.Println(c.B3 + "║ " + c.O + "High Priority   " + c.X + "┃┃ " + c.B1 + " Two day ban" + c.X + "           ┃┃  " + c.B1 + "Example: I'm presenting to my class and it broke!" + c.B3 + "          ║")
	fmt.Println(c.B3 + "║ " + c.R + "Urgent Priority " + c.X + "┃┃ " + c.B1 + " Three to Five day ban" + c.X + " ┃┃  " + c.B1 + "Example: The server blew up!" + c.B3 + "                               ║")
	fmt.Println(c.B3 + "╚══════════════════════════════════════════════════════════════════════════════════════════════════════════╝")
}
