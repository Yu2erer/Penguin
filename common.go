package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"
)

var (
	homePageCounts  = 2
	tplHomePagePath = "theme/index.html"
	publicPath      = "public"
	sourcePath      = "source"
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

func printUsage() {
	fmt.Println("Print Help")
}
