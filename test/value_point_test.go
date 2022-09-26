package test

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

// 传值/指针(引用)的区别其实与Java里概念是类似，各语言又有自己的特点
func TestName(t *testing.T) {
	v := Value{"张三", 18}
	v1 := val(v)
	v2 := point(&v)

	assert.NotEqual(t, v1, v2)
}

// 传值/指针区别
type Value struct {
	Name string
	Age  int
}

func val(v Value) Value {
	v.Name += ".val"
	v.Age += 1
	return v
}

func point(v *Value) Value {
	v.Name += ".point"
	v.Age += -1
	return *v
}
