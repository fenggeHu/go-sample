package test

import (
	"fmt"
	"math"
	"testing"
)

func TestSum(t *testing.T) {
	i := 100
	fmt.Println(float64(i))

	a := 1000.12
	p := 50.0
	fmt.Println(math.Floor(a / p))
}
