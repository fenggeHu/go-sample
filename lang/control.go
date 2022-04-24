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
	var x = []int{1, 3, 5, 7, 13}
	var total float64 = 0
	for i, value := range x {
		fmt.Printf("x第%d个element是%d\n", i+1, value)
		total += float64(value)
	}
	fmt.Printf("把x所有元素平均得到%f\n", total/float64(len(x)))

	for i := 1; i < 100; i++ {
		if i%10 == 0 {
			fmt.Printf("%d, its multiple of 10\n", i)
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
