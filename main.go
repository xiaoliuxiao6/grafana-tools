package main

import (
	kingpin "gopkg.in/alecthomas/kingpin.v2"
	"log"
	"os"
)

var (
	grafanaURL string
	grafanaAPI string

	app    = kingpin.New("chat", "Grafana 管理工具")
	bakcup = app.Command("bak", "备份所有 Dashboard,（可选）指定备份路径，默认为当前目录")
	dir    = bakcup.Arg("dir", "（可选）指定备份路径，默认为当前目录").String()
)

func init() {
	grafanaURL = os.Getenv("grafanaURL")
	if len(grafanaURL) == 0 {
		log.Fatalf("无法获取 Grafana URL，请指定 grafanaURL 变量")
	}

	grafanaAPI = os.Getenv("grafanaAPI")
	if len(grafanaAPI) == 0 {
		log.Fatalf("无法获取 Grafana API，请指定 grafanaAPI 变量")
	}
}

func main() {

	//kingpin.CommandLine.HelpFlag.Short('h') // 启用 -h 短参数
	app.Version("0.1.0")    // 定义版本号
	app.HelpFlag.Short('h') // 启用 -h 短参数

	//kingpin.Parse()                         // 解析
	switch kingpin.MustParse(app.Parse(os.Args[1:])) { // 解析 {
	case bakcup.FullCommand():
		if len(*dir) > 0 {
			bakDir = *dir
		}
		backupDashboards()
	}

	// 如果没有指定任何信息的话默认打印帮助信息
	//kingpin.Usage()
}
