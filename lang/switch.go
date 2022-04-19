package main

import (
	"fmt"
	"time"
)

func switchMain() {
	today := time.Now().Weekday()
	fmt.Println(today)
	fmt.Println(today - 2)
	switch time.Wednesday {
	case today:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	}
}
