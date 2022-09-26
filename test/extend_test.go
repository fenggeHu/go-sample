package test

import (
	"fmt"
	"testing"
)

func TestExtends(t *testing.T) {
	// error
	//p := &Parent
	//m := &Me{Parent{Man{"me"}}}
	m := &GrandSon{}
	m.Say("Win")
}

func (m *Man) Relation(m2 *Man) {
	fmt.Printf("m1:%s - m2:%s", m.Name, m2.Name)
}

type Man struct {
	Name string
	// 定义一个空方法，子类实现这个方法。 达到类似Java的抽象类的效果
	Greeting func()
}
type IMan interface {
	Say(message string)
}

func (m *Man) Say(message string) {
	m.Greeting = func() {
		fmt.Println("Man said, Hello")
	}
	m.Greeting()
	fmt.Printf("Man: %s\n", message)
}

type Parent struct {
	Man
}

// 实现了Man的抽象方法Greeting
func (m *Parent) Greeting() {
	fmt.Printf("Hello\n")
}

func (m *Parent) Say(message string) {
	fmt.Printf("Parent: %s\n", message)
}

type Me struct {
	Parent
}

func (m *Me) Say(message string) {
	fmt.Printf("Me: %s\n", message)
}

type Son struct {
	Me
}

//func (m *Son) Say(message string) {
//	fmt.Printf("Son: %s\n", message)
//}

type GrandSon struct {
	Son
}

func (m *GrandSon) Say(message string) {
	m.Son.Say(message)
	m.Greeting()
	fmt.Printf("GrandSon: %s\n", message)
}
