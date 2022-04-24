package main

import (
	"fmt"
	"time"
)

// Control Structures: for/if/switch
func controlMain() {
	forIfMe()
	switchMe()
}

func forIfMe() {
	for i := 0; i < 100; i++ {
		if i%10 == 0 {
			fmt.Printf("%d, its multiple of 10", i)
		}
	}
}

func switchMe() {
	fmt.Println("switch example: ")
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
