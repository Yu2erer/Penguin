package main

import (
	"flag"
	"os"
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
		new()
	case "new":
		// 新建 MarkDown
		if len(args) != 2 {
			printUsage()
			os.Exit(1)
		}
		name := args[1]
		createMarkDown(name)
	case "build":
		// 编译
		build()
	case "http":
		// 运行本地服务器
		_http()
	case "clean":
		// 清空 source 里的文件
		clean()
	case "help":
		printUsage()
	default:
		printUsage()
		os.Exit(1)
	}
}
