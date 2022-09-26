package test

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// channel 传对象还是指针跟方法参数一样。如是否希望被修改，视具体场景而定，和java类似
type ChanDO struct {
	Id   string
	Name string
	Age  int
}

func send(cd chan *ChanDO, d *ChanDO) {
	cd <- d
	fmt.Printf("send data: %v\n", d)
}

// 不能缺少for
func receive(cd chan *ChanDO) {
	for { // for循环接收
		v, ok := <-cd // 建立了通道-阻塞线程，如果没有for则在接收一次后关闭了通道
		fmt.Printf("[%v]Receive Data: %v\n", ok, v)
		time.Sleep(2 * time.Second)
	}
}

func TestChan(t *testing.T) {
	ch := make(chan *ChanDO, 1)
	defer close(ch)
	go receive(ch) // 因为channel阻塞了线程，所以必须go线程

	for i := 0; i < 5; i++ {
		cd := &ChanDO{"id_" + strconv.Itoa(i), "name_" + strconv.Itoa(i), 20 + i}
		go send(ch, cd) // 同步或异步 - 如果要保证顺序就使用同步
		//time.Sleep(100 * time.Millisecond)
	}
	time.Sleep(10 * time.Second)
}

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
