package test

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestPi(t *testing.T) {
	start := time.Now().UnixNano()
	count := 10000000
	half := count / 2
	base := float64(count)
	points := make(map[float64]float64, count)
	for i := 0; i < count; i++ {
		rand.Seed(time.Now().UnixNano())
		x := float64(rand.Intn(count)-half) / base
		y := float64(rand.Intn(count)-half) / base
		points[x] = y
	}
	var in int64
	rr := 0.5
	for x, y := range points {
		r := math.Sqrt(x*x + y*y)
		if r <= rr {
			in++
		}
	}
	pi := float64(in) * 4 / base
	fmt.Printf("%d:%v", time.Now().UnixNano()-start, pi)
}
