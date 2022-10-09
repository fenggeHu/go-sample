package test

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type point struct {
	x float64
	y float64
}

func Pi() {
	start := time.Now().UnixNano()
	count := 1000000
	half := count / 2
	base := float64(count)
	points := make([]point, count)
	for i := 0; i < count; i++ {
		rand.Seed(time.Now().UnixNano())
		x := float64(rand.Intn(count)-half) / base
		y := float64(rand.Intn(count)-half) / base
		points[i].x = x
		points[i].y = y
	}
	var in int64
	rr := 0.5
	for _, p := range points {
		r := math.Sqrt(p.x*p.x + p.y*p.y)
		if r <= rr {
			in++
		}
	}
	pi := float64(in) * 4.0 / base
	fmt.Printf("%dms - pi:%v", (time.Now().UnixNano()-start)/1000000, pi)
}
