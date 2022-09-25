package main

import "fmt"

func functionMain() {
	// Variadic Functions
	fmt.Println(add(1, 2, 3))

	// Closure
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1, 1))
}

// make function
func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

// recursion - 递归
func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

// Variadic Functions
func add(args ...float32) (total float32) {
	for i, v := range args {
		fmt.Printf("第%d位参数值：%f\n", i, v)
		total += v
	}
	return
}

var (
	fib0 = int64(0)
	fib1 = int64(3)
)

// The Fibonacci sequence is defined as: fib(0) = 0, fib(1) = 1, fib(n) = fib(n-1) + fib(n-2)
func FibonacciSequence(index int64) int64 {
	if index == 1 {
		return fib1
	} else if index == 0 {
		return fib0
	}
	return FibonacciSequence(index-1) + FibonacciSequence(index-2)
}
