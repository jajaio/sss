package nil

import (
	t "time"
)

//Good god what have I gotten myself into?

func TestForBannedUser(username string) bool {
	isBanned := false
	switch username {
	case user1.name:
		if now.After(user1.until) == true {
			isBanned = true
		} else {
			isBanned = false
		}
	case user2.name:
		if now.After(user2.until) == true {
			isBanned = true
		} else {
			isBanned = false
		}
	default:
		isBanned = false
	}
	return isBanned
}

const Version string = "\033[1;35mV S-1.0.0"

type Banned struct {
	name  string
	until t.Time
}

var now = t.Now()

//this is the part where I declare the banned users.

var user1 = Banned{
	name:  "sshhh",
	until: t.Date(2010, t.January, 05, 9, 13, 30, 0, t.Local),
}
var user2 = Banned{
	name:  "you",
	until: t.Date(9999, t.December, 59, 23, 59, 59, 0, t.Local),
}
