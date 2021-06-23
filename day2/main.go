package main

import (
	"fmt"
	"time"
)

func main() {
	testSwitch()
}

func testIf() {
	i := 7
	if i == 7 {
		fmt.Println("7")
	}
}

func testIfElse() {
	points := 6000
	if points < 5000 {
		fmt.Println("Silver")
	} else if points < 10000 {
		fmt.Println("Gold")
	} else {
		fmt.Println("Platinum")
	}
}

func testSwitch() {
	weekday := time.Now().Weekday()
	switch weekday {
	case 1:
		fmt.Println("Dushanba")
	case 2:
		fmt.Println("Seshanba")
	case 3:
		fmt.Println("Chorshanba")
	case 4:
		fmt.Println("Payshanba")
	case 5:
		fmt.Println("Juma")
	case 6:
		fmt.Println("Shanba")
	case 7:
		fmt.Println("Yakshanba")
	default:
		fmt.Println("Xato")
	}
}
