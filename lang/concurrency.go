package main

//import (
//	"fmt"
//	"math/rand"
//	"os"
//	"time"
//)
//
//var quantity = 1000
//
//func concurrencyMain() {
//	go func() {
//		for {
//			num := get()
//			//fmt.Printf("库存量：%d\n", num)
//			if num < 0 {
//				fmt.Errorf("超卖%d\n", num)
//				os.Exit(-1)
//			}
//			time.Sleep(10 * time.Millisecond)
//		}
//	}()
//	for i := 0; i < 10; i++ {
//		go func() {
//			for {
//				num := rand.Intn(5)
//				if num <= 0 {
//					fmt.Errorf("rand int err: %d\n", num)
//				}
//				if !out(num) {
//					//fmt.Printf("库存不足%d\n", num)
//					time.Sleep(5 * time.Millisecond)
//				}
//			}
//		}()
//	}
//}
//
//// 查询数量
//func get() int {
//	return quantity
//}
//
//// 扣减数量
//func out(n int) bool {
//	if quantity >= n {
//		quantity = quantity - n
//		return true
//	} else {
//		return false
//	}
//}
