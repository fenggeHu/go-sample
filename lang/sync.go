package main

import "sync"

var mymap = &sync.Map{}

func syncMap() {

}

func put(k string, v interface{}) {
	mymap.Store(k, v)
}

func remove(k string) {
	mymap.Delete(k)
}

func get(k string) interface{} {
	v, _ := mymap.Load(k)
	return v
}
