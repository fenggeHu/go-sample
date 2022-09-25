package main

import (
	"log"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	//now := time.Now()
	//yday := now.YearDay()
	//log.Printf("\n%d", yday)

	date := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	log.Printf("\n%v", date.IsZero()) // false
	log.Printf("\nsec: %v", date.Second())
	log.Printf("\nnsec: %v", date.Nanosecond())
	log.Printf("\nTimeFromUnixZero: %v", time.Unix(0, 0))
}
