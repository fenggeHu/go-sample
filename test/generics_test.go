package test

import (
	"fmt"
	"testing"
)

type Number interface {
	int | int64 | float64
}

func sum[V Number](n1, n2 V) V {
	return n1 + n2
}

func TestGenerics(t *testing.T) {
	s1 := sum(12, 23)
	fmt.Println(s1)

	s2 := sum(12.11, 23.23)
	fmt.Println(s2)
}
