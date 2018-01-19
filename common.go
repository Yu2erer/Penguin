package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"time"
)

var (
	tplHomePagePath = "theme/index.html"
	tplBlogPagePath = "theme/blog.html"
	publicPath      = "public"
	sourcePath      = "source"
	themePath       = "theme"
	liConnect       = "<li class=\"{{ .ConnectWay }}\"><a href=\"{{ .Address }}\"><i class=\"fa fa-{{ .ConnectWay }}\"></i></a></li>"
	liListItem      = "<li class=\"article_list_item\"><span class=\"article_list_link\"><a href=\"{{ .Address }}\"><h3 class=\"mar-t-z\">{{ .Title }}</h3></a><samll class=\"publish_time\">{{ .Date }}</samll></span><small class=\"tag {{ .TagColor }}\">{{ .TagTitle }}</small></li>"
)

var (
	date     = time.Now().Format("2006.01.02")
	markdown = `---
title: Your article title
date: ` + date + `
tag-title: Your article tag title
tag-color: Your article tag color - red blue yellow green
---`
)

var pengiueStr = `                                         __                     
|  \                    
______    ______   _______    ______   \$$ __    __   ______  
/      \  /      \ |       \  /      \ |  \|  \  |  \ /      \ 
|  $$$$$$\|  $$$$$$\| $$$$$$$\|  $$$$$$\| $$| $$  | $$|  $$$$$$\
| $$  | $$| $$    $$| $$  | $$| $$  | $$| $$| $$  | $$| $$    $$
| $$__/ $$| $$$$$$$$| $$  | $$| $$__| $$| $$| $$__/ $$| $$$$$$$$
| $$    $$ \$$     \| $$  | $$ \$$    $$| $$ \$$    $$ \$$     \
| $$$$$$$   \$$$$$$$ \$$   \$$ _\$$$$$$$ \$$  \$$$$$$   \$$$$$$$
| $$                          |  \__| $$                        
| $$                           \$$    $$                        
\$$                            \$$$$$$     
    
`

const HELP = `
Usage:

pengiue command [args...]

	初始化博客文件夹
	pengiue init

	新建 markdown 文件
	pengiue new filename

	编译博客
	pengiue build

	清理博客
	pengiue clean

	打开本地服务器 [http://localhost:12345]
	pengiue http

`

func createMarkDown(filename string) {
	file := path.Join(sourcePath, filename+".md")
	_, err := os.Stat(file)
	if !os.IsNotExist(err) {
		fmt.Println(file + " 已存在")
		os.Exit(1)
	}
	_, err = os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	_ = ioutil.WriteFile(file, []byte(markdown), 0666)
}
func new() {
	_, err := os.Stat(sourcePath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(sourcePath, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	_, err = os.Stat(publicPath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(publicPath, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}
func _http() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(publicPath))))
	http.Handle("/theme/", http.StripPrefix("/theme/", http.FileServer(http.Dir(themePath))))
	if err := http.ListenAndServe(":12345", nil); err != nil {
		log.Fatalln(err)
	}
}
func clean() {
	dir_list, err := ioutil.ReadDir(sourcePath)
	if err != nil {
		fmt.Println("Source 文件夹不存在, 请先 init")
		os.Exit(1)
	}
	for _, dir := range dir_list {
		err = os.Remove(sourcePath + "/" + dir.Name())
		if err != nil {
			fmt.Println(dir.Name() + " 删除失败")
		}
	}
}
func checkFile() {
	if _, err := os.Stat(themePath); os.IsNotExist(err) {
		fmt.Println("Theme 模板文件夹丢失")
		os.Exit(1)
	}
	if _, err := os.Stat(sourcePath); os.IsNotExist(err) {
		fmt.Println("请先 init")
		os.Exit(1)
	}
	if _, err := os.Stat(publicPath); os.IsNotExist(err) {
		fmt.Println("请先 init")
		os.Exit(1)
	}
}
func printUsage() {
	fmt.Println(HELP)
}
