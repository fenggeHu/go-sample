package main

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	channelMain()
	//
	time.Sleep(10 * time.Second)
}

// 下面receiver和selectChan都接收了c，但1条信息只消费1次
func channelMain() {
	c := make(chan string, 10)
	c2 := make(chan string, 20)
	c3 := make(chan string, 20)
	//go ponger(c)
	go receiver(c)
	go sender("tag-c", c)
	go selectChan(c, c2, c3)
	go sender("tag-c2", c2)
	go sender("tag-c3", c3)
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
		time.Sleep(5 * time.Second)
	}
}

// 定时发送
func sender(tag string, c chan string) {
	for i := 0; i < 2; i++ {
		c <- fmt.Sprintf("[%s]_%d", tag, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func receiver(c chan string) {
	for {
		msg := <-c
		fmt.Printf("receiver: %s\n", msg)
	}
}

// Go has a special statement called select which works like a switch but for channels
func selectChan(c1 chan string, c2 chan string, c3 chan string) {
	for {
		select {
		case m1 := <-c1:
			fmt.Printf("m1 <-c1: %s\n", m1)
		case <-c1: // 共同消费c1，也是有效的
			fmt.Printf("<-c1: %s\n", time.Now())
		case m2 := <-c2:
			fmt.Printf("m2 <-c2: %s\n", m2)
		case m3 := <-c3:
			fmt.Printf("m3 <-c3: %s\n", m3)
		case m := <-time.After(time.Second):
			fmt.Printf("timeout: %s\n", m)
			//default:		// The default case happens immediately if none of the channels are ready。没有业务逻辑的时候不要
			//	fmt.Println("default")
		}
	}
}
