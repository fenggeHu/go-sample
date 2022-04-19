package main

import (
	"log"
	"strings"
	"time"
)

func containsMain() {
	str := "65539.hk,63312.hk,56495.hk,52648.hk,09988.hk,56335.hk,69461.hk,62833.hk,00700.hk,03690.hk,01810.hk,03968.hk,06862.hk,02318.hk"
	array := strings.Split(str, ",")

	start := time.Now().UnixMilli()
	for i := 0; i < 1000000; i++ {
		strings.Contains(str, "52648.hk,")
		strings.Contains(str, "03968.hk,")
		strings.Contains(str, "02648.hk,")

		//index1 := strings.Index(str, "52648.hk,")
		//index2 := strings.Index(str, "03968.hk,")
		//index3 := strings.Index(str, "02648.hk,")
		//log.Printf("%d - %d - %d", index1, index2, index3)
	}
	log.Printf("strings.Index: %d", time.Now().UnixMilli()-start)

	start = time.Now().UnixMilli()
	for i := 0; i < 1000000; i++ {
		var sucess1 bool
		for _, v := range array {
			sucess1 = strings.EqualFold(v, "52648.hk")
			if sucess1 {
				break
			}
		}
		var sucess2 bool
		for _, v := range array {
			sucess2 = strings.EqualFold(v, "03968.hk")
			if sucess2 {
				break
			}
		}
		var sucess3 bool
		for _, v := range array {
			sucess3 = strings.EqualFold(v, "02648.hk")
			if sucess3 {
				break
			}
		}
		//log.Printf("%t - %v - %v", sucess1, sucess2, sucess3)
	}
	log.Printf("strings.Index: %d", time.Now().UnixMilli()-start)
}
