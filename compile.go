package main

import (
	"strings"
	"os"
	"path"
	"log"
	"io/ioutil"
	"html/template"
)

type HomePage struct {
	Title string
	Description string
	Connect template.HTML
}
var tplHomePage string

func prepareExampleHomePage() HomePage {
	github := "https://github.com"
	weibo := "https://weibo.com"
	connectWays := map[string]string {
		"github": github,
		"weibo": weibo,
	}
	var connect template.HTML
	for connectWay, connectAddress := range connectWays {
		if connectAddress != "" {
			way := strings.Replace(liConnect, "{{ .Address }}", connectAddress, 1)
			way = strings.Replace(way, "{{ .ConnectWay }}", connectWay, -1)
			connect += template.HTML(way + "\n")
		}
	}
	return HomePage{"我的博客Title", "我的博客描述", connect}
}
func preparetplHomePage() {
	byte, err := ioutil.ReadFile(tplHomePagePath)
	if err != nil {
		log.Fatalln("文件丢失")
	}
	tplHomePage = string(byte)
}
func compile() {
	preparetplHomePage()
	exampleHomePage := prepareExampleHomePage()
	t, err := template.New("tplHomePage").Parse(tplHomePage)
	filePath := path.Join(publicPath, "index.html")
	f, err := os.Create(filePath)
	if err != nil {
		log.Fatalln("创建文件失败")
	}
	t.Execute(f, exampleHomePage)
}