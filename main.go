package main

import (
	"flag"

	"github.com/mapleafgo/sysproxy"
)

func main() {
	var name string
	var value string
	flag.StringVar(&name, "n", "", "控制名")
	flag.StringVar(&value, "l", "", "pac url")
	flag.Parse()
	if name == "open" {
		sysproxy.TurnOnSystemProxy(value)
	} else if name == "close" {
		sysproxy.TurnOffSystemProxy()
	}
}
