package main

import "fmt"

func functionMain() {
	// Variadic Functions
	fmt.Println(add(1, 2, 3))
}

// Variadic Functions
func add(args ...float32) (total float32) {
	for i, v := range args {
		fmt.Printf("第%d位参数值：%f\n", i, v)
		total += v
	}
	return
}
