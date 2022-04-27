package main

import (
	"log"
	"os"
)

var basePath = "/Users/max/.qlib/qlib_data/cn_data"

func main() {
	features := basePath + "/features/sh600004/close.day.bin"
	f, err := os.Open(features)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	//dense := mat.Dense{}
	//n, err := dense.UnmarshalBinaryFrom(f)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//log.Println(n)

	//r, _ := gonpy.NewFileReader(features)
	//data, _ := r.GetFloat64()
	//
	//log.Println(data)

	//var m [6]byte	//var magic [6]byte
	//binary.Read(f, binary.LittleEndian, m)
	//for len(m) >0 {
	//	fmt.Printf("data = %v\n", m)
	//	binary.Read(f, binary.LittleEndian, m)
	//}

	//var m []float64
	//err = npyio.Read(f, &m)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//r, err := npyio.NewReader(f)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Printf("npy-header: %v\n", r.Header)
	//shape := r.Header.Descr.Shape
	//raw := make([]float64, shape[0]*shape[1])
	//
	//err = r.Read(&raw)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//m := mat.NewDense(shape[0], shape[1], raw)
	//fmt.Printf("data = %v\n", mat.Formatted(m, mat.Prefix("       ")))
}
