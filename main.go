package main

import (
	"flag"
	"wenkang365t.cn/sysproxy/system"
)

func main() {
	var name string
	var value string
	flag.StringVar(&name, "n", "", "控制名")
	flag.StringVar(&value, "l", "", "pac url")
	flag.Parse()
	if name == "open" {
		system.TurnOnSystemProxy(value)
	} else if name == "close" {
		system.TurnOffSystemProxy()
	}
}
