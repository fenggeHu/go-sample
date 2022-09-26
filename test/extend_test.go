package main

import (
	"fmt"
	"testing"
)

func TestExtends(t *testing.T) {
	// error
	//p := &Parent
	//m := &Me{Parent{Man{"me"}}}
}

func (m *Man) Relation(m2 *Man) {
	fmt.Printf("m1:%s - m2:%s", m.Name, m2.Name)
}

type Man struct {
	Name string
}

type Parent struct {
	Man
}

type Me struct {
	Parent
}

type Son struct {
	Me
}

type GrandSon struct {
	Son
}
