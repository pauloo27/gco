package main

import (
	"flag"
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
		// TODO: init
		return
	}
	// TODO
}
