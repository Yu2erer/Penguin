package main

import (
	"os"
	"flag"
	"fmt"
)
var (
	args []string
)
func main() {
	flag.Parse()
	args = flag.Args()
	if len(args) == 0 || len(args) > 3 {
		printUsage()
		os.Exit(1)
	}
	switch args[0] {
	case "init":
		// 初始化
	case "new":
		// 新建 MarkDown
	case "complie":
		// 编译
	}
}

func printUsage() {
	fmt.Println("Print Help")
}
