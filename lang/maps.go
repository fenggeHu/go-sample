package main

import "fmt"

func mapsMain() {
	m := make(map[int]interface{}, 16)
	m[0] = `zero`
	m[1] = "壹"
	m[2] = 2
	m[3] = "三"
	m[4] = "四"
	m[5] = 5

	for k, v := range m {
		fmt.Printf("k: %d, v: %v\n", k, v)
	}

	delete(m, 2)
	fmt.Println("deleted m[2]")

	for k, v := range m {
		fmt.Printf("k: %d, v: %v\n", k, v)
	}
}
