package main

import (
	"bytes"
	"fmt"
	"strings"
)

func strMain() {
	InOutPut()
}

func InOutPut() {
	str := "hello world"
	// To read or write to a []byte or a string you can use the Buffer struct found in the bytes package:
	var buf bytes.Buffer
	buf.Write([]byte(str))
	fmt.Printf("%s, len: %d\n", buf.String(), buf.Len())

	// A Buffer doesn't have to be initialized and supports both the Reader and Writer interfaces.
	//You can convert it into a []byte by calling buf.Bytes().
	//If you only need to read from a string you can also use the strings.NewReader function
	//which is more efficient than using a buffer.
	reader := strings.NewReader(str)
	fmt.Printf("len: %d\n", reader.Len())
}

// string常见操作
func str() {
	fmt.Println(
		// true
		strings.Contains("test", "es"),

		// 2
		strings.Count("test", "t"),

		// true
		strings.HasPrefix("test", "te"),

		// true
		strings.HasSuffix("test", "st"),

		// 1
		strings.Index("test", "e"),

		// "a-b"
		strings.Join([]string{"a", "b"}, "-"),

		// == "aaaaa"
		strings.Repeat("a", 5),

		// "bbaa"
		strings.Replace("aaaa", "a", "b", 2),

		// []string{"a","b","c","d","e"}
		strings.Split("a-b-c-d-e", "-"),

		// "test"
		strings.ToLower("TEST"),

		// "TEST"
		strings.ToUpper("test"),
	)
}
