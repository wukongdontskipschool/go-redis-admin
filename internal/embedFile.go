package internal

import (
	"embed"
	"io/fs"
)

// 静态文件前置目录
var preDir string = "web/weAdmin/"

//go:embed web
var F embed.FS

// 静态文件打开目录
type EmbedDir string

func (dir EmbedDir) Open(name string) (fs.File, error) {
	name = preDir + string(dir) + "/" + name
	return F.Open(name)
}

func GetFs() embed.FS {
	return F
}

func GetIndexHtml() string {
	return preDir + "index.html"
}

func GetLoginHtml() string {
	return preDir + "login.html"
}
