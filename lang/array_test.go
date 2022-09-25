package main

import (
	"log"
	"strconv"
	"testing"
)

func TestArray(t *testing.T) {
	a := []string{"h", "e", "l"}
	set(a...)
	log.Printf("%s", a)
}

// 可变参数是指针传递
func set(sa ...string) {
	sa[0] = strconv.Itoa(len(sa))
}
