package main

import (
	"fmt"
)

func main() {

	//// if
	//if f, err := os.OpenFile("/Users/hujinfeng/test/pom.xml", os.O_RDONLY, 0666); err != nil {
	//	fmt.Printf("%s", err)
	//} else {
	//	var b []byte
	//	n, err := f.Read(b)
	//	if err != nil {
	//		fmt.Printf("%d - %s", n, err)
	//	}
	//	fmt.Println(b)
	//}

	//// for 常见3种样式
	//for i := 0; i < 1; i++ {
	//	fmt.Printf("Now: %d\n", i)
	//	for i > 90 && i < 100 {
	//		i++
	//		for {
	//			fmt.Printf("break: %d\n", i)
	//			break
	//		}
	//	}
	//}
	//
	//// for 遍历
	//// 遍历map
	//mymap := map[int]string{1: "1111", 2: "2222"}
	//for k, v := range mymap {
	//	fmt.Printf("k: %d v: %s\n", k, v)
	//}
	//mymap[100] = "1000000"
	//for k := range mymap {
	//	v := mymap[k]
	//	fmt.Printf("k: %d v: %s\n", k, v)
	//}
	//for _, v := range mymap {
	//	fmt.Printf("v: %s\n", v)
	//}
	//
	//// 遍历字符串
	//for pos, ch := range "hello，胡" {
	//	fmt.Printf("位置: %d, 字符: %#U - %d\n", pos, ch, ch)
	//}

	// switch
	mm := map[int]interface{}{1: "1111", 2: 2222, 3: 3.333, 4: false}
	for _, v := range mm {
		switch t := v.(type) {
		default:
			fmt.Printf("unexpected type %T\n", t) // %T prints whatever type t has
		case bool:
			fmt.Printf("boolean %t\n", t) // t has type bool
		case int:
			fmt.Printf("integer %d\n", t) // t has type int
		case *bool:
			fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
		case *int:
			fmt.Printf("pointer to integer %d\n", *t) // t has type *int
		}
	}
}
