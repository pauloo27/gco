package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Pauloo27/gommit/mode"
)

var isHelp, isInit bool

func init() {
	flag.BoolVar(&isHelp, "help", false, "show command help")
	flag.BoolVar(&isInit, "init", false, "create a project config")
}

func main() {
	flag.Parse()
	if isHelp {
		flag.Usage()
		return
	}
	if isInit {
		mode.Init()
		return
	}
	// TODO
	fmt.Println("not implemented yet")
	os.Exit(-1)
}
