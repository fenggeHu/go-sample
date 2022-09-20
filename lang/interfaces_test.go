package main

import (
	"fmt"
	"testing"
)

type Closable interface {
	// Close release all resources used by this object, including goroutines.
	Close() error
}

func Close(obj interface{}) error {
	if c, ok := obj.(Closable); ok {
		return c.Close()
	}
	return nil
}

func TestClose(t *testing.T) {
	obj := "hello"
	err := Close(obj)
	fmt.Println(err)
}
