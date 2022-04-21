package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

// 读写锁
func lockMain() {
	var counter = &Counter{
		rwLock:  &sync.RWMutex{},
		current: -1,
	}
	// 写程
	go func() {
		for {
			num := rand.Intn(100)
			time.Sleep(time.Duration(num) * time.Millisecond)
			counter.entry()
		}
	}()

	// 读程
	go func() {
		for {
			time.Sleep(3 * time.Second)

			//time := time.Now().Minute()
			time := time.Now().Second()
			record := counter.read(time, 5)
			log.Printf("record: %v", record)
		}
	}()
}

// 计数器 - 按分针或秒针滑动计数 - 初始化60位长度的数组
type Counter struct {
	rwLock  *sync.RWMutex
	current int     // 记录当前时间的序号
	count   [60]int // 初始化计数器 - 按时间0~59分计数请求
}

// 每调用一次本方法，表示一个记录
func (c *Counter) entry() {
	//log.Printf("entring... ...")
	c.rwLock.Lock()
	defer c.rwLock.Unlock()

	//min := time.Now().Minute()
	min := time.Now().Second()
	if c.current == min {
		c.count[min] = c.count[min] + 1
	} else {
		c.current = min
		c.count[min] = 1
	}
	//log.Printf("entried... ...")
}

// 读取当前时刻minute之前的round个记录
func (c *Counter) read(minute int, round int) (out []int) {
	log.Printf("read -->minute time: %d, round: %d", minute, round)
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()

	//cpa := c.count // 数组赋值是传值 // 这里只读数据不修改，所以不需要新变量
	if minute >= round {
		out = c.count[minute-round : minute]
	} else {
		out = c.count[60+minute-round:]
		for _, num := range c.count[:minute] {
			out = append(out, num)
		}
	}
	log.Printf("all record: %v", c.count)
	return
}
