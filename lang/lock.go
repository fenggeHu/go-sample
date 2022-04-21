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
	go func() {
		for {
			num := rand.Intn(100)
			time.Sleep(time.Duration(num) * time.Millisecond)
			counter.entry()
		}
	}()

	go func() {
		for {
			time.Sleep(2 * time.Second)

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
	current int     //
	count   [60]int // 初始化计数器 - 按时间0~59分计数请求
}

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

func (c *Counter) read(minute int, round int) (out []int) {
	log.Printf("read -->minute time: %d, round: %d", minute, round)
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()
	if minute >= round {
		out = c.count[minute-round : minute]
		//copy(c.count[minute-round:minute], out)
	} else {
		out = c.count[:minute]
		for i := range c.count[58+minute-round:] {
			out = append(out, i)
		}
	}
	log.Printf("all record: %v", c.count)
	return
}
