package main

import (
	"flag"
	"fmt"
)

func main() {
	parseArgs()

	//index := 100
	//start := time.Now().UnixMilli()
	//num := FibonacciSequence(int64(index))
	//fmt.Printf("time over: %d, fib(%d)=%d", time.Now().UnixMilli()-start, index, num)

	//controlMain()
	//strMain()
	//mapsMain()
	//functionMain()
	//channelMain()

	//concurrencyMain()

	//lockMain()

	//var input string
	//fmt.Scanln(&input)
}

// program arguments format：
// 单横杆或双横杆+参数名 + 等号或空格 + 值，格式可以混用
// -name=hu	/ -name hu
// --name=hu / --name hu
func parseArgs() {
	// 2种赋值方式
	var count int
	flag.IntVar(&count, "count", 10, "count")
	name := flag.String("name", "", "Input your name")
	flag.Parse() // 【必须调用】从 arguments 中解析注册的 flag

	fmt.Printf("my name is %s\n", *name)
	for i := 0; i < count; i++ {
		fmt.Println(i)
	}
}
