package main

import (
	"fmt"
	"time"
)

func channelMain() {
	c := make(chan string)
	go sender(c)
	go receiver(c)
}

func sender(c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("ping%d", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func receiver(c chan string) {
	for {
		msg := <-c
		fmt.Printf("receiver: %s\n", msg)
	}
}
