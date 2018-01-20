# Penguin
[![Release][1]][2]

[1]: https://img.shields.io/badge/release-v0.2-brightgreen.svg
[2]: https://github.com/ntian2/Penguin/releases

```text
 _______                                           __           
|       \                                         |  \          
| $$$$$$$\  ______   _______    ______   __    __  \$$ _______  
| $$__/ $$ /      \ |       \  /      \ |  \  |  \|  \|       \ 
| $$    $$|  $$$$$$\| $$$$$$$\|  $$$$$$\| $$  | $$| $$| $$$$$$$\
| $$$$$$$ | $$    $$| $$  | $$| $$  | $$| $$  | $$| $$| $$  | $$
| $$      | $$$$$$$$| $$  | $$| $$__| $$| $$__/ $$| $$| $$  | $$
| $$       \$$     \| $$  | $$ \$$    $$ \$$    $$| $$| $$  | $$
 \$$        \$$$$$$$ \$$   \$$ _\$$$$$$$  \$$$$$$  \$$ \$$   \$$
                              |  \__| $$                        
                               \$$    $$                        
                                \$$$$$$                         
```
Penguin 是一款小巧的 Markdown 静态博客生成器 它使用 `Golang` 开发

# Install
从此处 [releases](https://github.com/ntian2/Penguin/releases) 下载

# Usage
```text
Usage:

pengiue command [args...]

	初始化博客文件夹
	penguin init

	新建 markdown 文件
	penguin new filename

	编译博客
	penguin build

	清理博客
	penguin clean

	打开本地服务器 [http://localhost:12345]
	penguin http
```
## penguin init
> penguin 会自动生成 source 和 public 文件夹 以及 config.yaml
```
# 站点信息
title: 网站名称
description: 自我描述
github: github 地址 可为空
weibo: weibo 地址 可为空
# Logo 与 头像 放置于 theme/asserts/images 文件目录下
# 其中 source 为 markdown博文源文件 public 为生成的页面
```
## License
[Apache License 2.0](https://github.com/ntian2/Penguin/blob/master/LICENSE)