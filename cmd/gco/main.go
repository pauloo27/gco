package main

import (
	"flag"

	"github.com/Pauloo27/gco/internal/mode"
)

var isHelp, isInitGlobal, isInit, isRetry, isSkip bool

func init() {
	flag.BoolVar(&isHelp, "help", false, "show command help")
	flag.BoolVar(&isInit, "init", false, "create a project config")
	flag.BoolVar(&isRetry, "retry", false, "retry a commit from .gobkp")
	flag.BoolVar(&isSkip, "skip", false, "skip pre/post hooks")
	flag.BoolVar(&isInitGlobal, "init-global", false, "create a global config")
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
	if isInitGlobal {
		mode.InitGlobal()
		return
	}
	if isRetry {
		mode.Retry()
		return
	}
	mode.Commit(isSkip)
}
