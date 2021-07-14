package main

import (
	"flag"

	"github.com/Pauloo27/gco/mode"
)

var isHelp, isInit, isRetry, isSkip bool

func init() {
	flag.BoolVar(&isHelp, "help", false, "show command help")
	flag.BoolVar(&isInit, "init", false, "create a project config")
	flag.BoolVar(&isRetry, "retry", false, "retry a commit from .gobkp")
	flag.BoolVar(&isSkip, "skip", false, "skip pre/post hooks")
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
	if isRetry {
		mode.Retry()
		return
	}
	mode.Commit(isSkip)
}
