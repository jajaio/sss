package main

import (
	"fmt"
	c "github.com/skilstak/go/colors"
)

func main() {
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
