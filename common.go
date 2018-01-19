package main


var (
	tplHomePagePath = "theme/index.html"
	publicPath = "public"
	liConnect = "<li class=\"{{ .ConnectWay }}\"><a href=\"{{ .Address }}\"><i class=\"fa fa-{{ .ConnectWay }}\"></i></a></li>"
)