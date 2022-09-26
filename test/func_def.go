package main

import "errors"

type Who struct {
	Name    string
	Execute func()
}

// Go 函数的返回值或结果 “形参” 可被命名，并作为常规变量使用，就像传入的形参一样。
// 命名后，一旦该函数开始执行，它们就会被初始化为与其类型相应的零值； 若该函数执行了一条不带实参的 return 语句，则结果形参的当前值将被返回。
func (w *Who) Say(word string) (message string, err error) {
	if w == nil || w.Name == "" {
		err = errors.New("nobody is here")
	}
	message = w.Name + " say: " + word
	return
}
