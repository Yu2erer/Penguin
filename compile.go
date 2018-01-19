package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type HomePage struct {
	Title       string
	Description string
	Connect     template.HTML
	ArticleListItem 	template.HTML
}
type Article struct {
	Title    string
	Date     string
	TagTitle string
	TagColor string
	MdBody   string
	Filename string
}

var tplHomePage string

func prepareExampleHomePage() HomePage {
	github := "https://github.com"
	weibo := "https://weibo.com"
	connectWays := map[string]string{
		"github": github,
		"weibo":  weibo,
	}
	var connect template.HTML
	for connectWay, connectAddress := range connectWays {
		if connectAddress != "" {
			way := strings.Replace(liConnect, "{{ .Address }}", connectAddress, 1)
			way = strings.Replace(way, "{{ .ConnectWay }}", connectWay, -1)
			connect += template.HTML(way + "\n")
		}
	}
	ArticleListItem := prepareHomePageArticle()
	return HomePage{"我的博客Title", "我的博客描述", connect, ArticleListItem}
}
func prepareHomePageArticle() template.HTML {
	var listItem template.HTML
	articlelist := articleList()
	for _, article := range articlelist {
		art := strings.Replace(liListItem, "{{ .Title }}", article.Title, 1)
		art = strings.Replace(art, "{{ .Date }}", article.Date, 1)
		art = strings.Replace(art, "{{ .TagTitle }}", article.TagTitle, 1)
		art = strings.Replace(art, "{{ .TagColor }}", article.TagColor, 1)
		listItem += template.HTML(art + "\n")
	}
	return listItem
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
func articleList() []Article {
	markdownlist := markdownList()
	articlelist := []Article{}
	for _, md := range markdownlist {
		file, _ := ioutil.ReadFile(md)
		lines := strings.Split(string(file), "\n")
		title := string(lines[1])
		title = string([]rune(title)[7:])
		date := string(lines[2])
		date = string([]rune(date)[6:])
		tagtitle := string(lines[3])
		tagtitle = string([]rune(tagtitle)[10:])
		tagcolor := string(lines[4])
		tagcolor = string([]rune(tagcolor)[10:])
		filenameWithSuffix := path.Base(md)        //获取文件名带后缀
		fileSuffix := path.Ext(filenameWithSuffix) //获取文件后缀
		filename := strings.TrimSuffix(filenameWithSuffix, fileSuffix)
		mdbody := strings.Join(lines[6:len(lines)], "\n")
		articlelist = append(articlelist, Article{title, date, tagtitle, tagcolor, mdbody, filename})
	}
	return articlelist
}
func markdownList() (markdownlist []string) {
	filepath.Walk(sourcePath, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			log.Fatalln(err)
		}
		if strings.HasSuffix(f.Name(), ".md") {
			markdownlist = append(markdownlist, path)
		}
		return nil
	})
	return markdownlist
}
