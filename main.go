package main

import (
	"flag"
	"fmt"
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
			fmt.Println("缺少文件名")
			printUsage()
			os.Exit(1)
		}
		name := args[1]
		fmt.Println("生成的文件名:" + name)
		createMarkDown(name)
	case "compile":
		// 编译
		compile()
	case "clean":
		// 清空 source 里的文件
		clean()
	default:
		printUsage()
		os.Exit(1)
	}
}
