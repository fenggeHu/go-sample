package main

import (
	"go-sample/index"
	"log"
	"time"
)

var (
	vers    *bool
	help    *bool
	conf    *string
	testing *string
)

// function init run before main
func init() {

	//vers = flag.Bool("v", false, "display the version.")
	//help = flag.Bool("h", false, "print this help.")
	//conf = flag.String("f", "", "specify configuration file.")
	//testing = flag.String("t", "", "test configuration.")
	//flag.Parse()
	//
	//fmt.Println(*vers, *help, *conf, *testing)
}

func main() {
	//test.Pi()
	baseDir := "/Users/max/test"
	start := time.Now().UnixNano()
	files := index.Scan(baseDir)
	for _, f := range files {
		video := index.VideoInfo(f, baseDir)

		log.Println(video)
	}
	end := time.Now().UnixNano()

	log.Printf("%d, %d", (end-start)/1000000, len(files))
}
