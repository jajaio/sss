package nil

import (
	"strconv"
	"time"
)

const Version string = "\033[1;35mV S-1.0.0"

func GenerateID() string {
	t := time.Now()
	//set variable 't' as Time's "now" function
	year, month, day := t.Date()
	//Parces year, month and day from "t" using the date function
	formatMonth := "SET"
	//Declares a filler variable, nothing to see here.
	//below checks the month and sets formatMonth to it's abbreviation.
	switch month {
	case time.January:
		formatMonth = "JAN"
	case time.February:
		formatMonth = "FEB"
	case time.March:
		formatMonth = "MAR"
	case time.April:
		formatMonth = "APR"
	case time.May:
		formatMonth = "MAY"
	case time.June:
		formatMonth = "JUN"
	case time.July:
		formatMonth = "JUL"
	case time.August:
		formatMonth = "AUG"
	case time.September:
		formatMonth = "SEP"
	case time.October:
		formatMonth = "OCT"
	case time.November:
		formatMonth = "NOV"
	case time.December:
		formatMonth = "DEC"
	default:
		formatMonth = "NIL"
	}
	hour, min, _ := t.Clock()
	//Parces hour, minuite from "t" using the clock function
	strday := strconv.Itoa(day)
	stryear := strconv.Itoa(year)
	strhour := strconv.Itoa(hour)
	strmin := strconv.Itoa(min)
	//converts all of the above to strings for formatting

	uuid := strday + "-" + formatMonth + "-" + stryear + "~" + strhour + ":" + strmin
	//format's all of the variables into a UUID var and returns it
	return uuid
}
