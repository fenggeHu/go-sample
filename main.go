package main

import (
	"flag"
	"fmt"
	"go-sample/test"
)

var (
	vers    *bool
	help    *bool
	conf    *string
	testing *string
)

// function init run before main
func init() {
	vers = flag.Bool("v", false, "display the version.")
	help = flag.Bool("h", false, "print this help.")
	conf = flag.String("f", "", "specify configuration file.")
	testing = flag.String("t", "", "test configuration.")
	flag.Parse()

	fmt.Println(*vers, *help, *conf, *testing)
}

func main() {
	test.Pi()
}
