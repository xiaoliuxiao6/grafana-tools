package main

import (
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"io"
	"io/fs"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"time"
)

var (
	bakDir string
)

// 备份仪表板
func backupDashboards() {

	// 如果没有指定备份目录就自动生成
	if len(bakDir) == 0 {
		bakDir = fmt.Sprintf("grafana-dashboard-bak-%v", time.Now().Format("20060102-150405"))
	}

	// 创建备份目录
	if err := os.Mkdir(bakDir, 0600); err != nil {
		if !errors.Is(err, fs.ErrExist) {
			log.Fatalf("创建备份目录失败: %v", err)
		}
	}

	// 搜索所有面板并备份
	DashboardsFolderser := searchAll()
	iNum := 0
	for _, DashboardsFolder := range DashboardsFolderser {
		//if DashboardsFolder.Type == "dash-db" && DashboardsFolder.Title == "xxxzzz" {
		if DashboardsFolder.Type == "dash-db" {
			iNum++
			log.Printf("正在备份第 %v 个 Dashboard: %v", iNum, DashboardsFolder.Title)
			getDashboardByUID(DashboardsFolder)
		}
	}
}

// 通过 UID 获取仪表板信息
func getDashboardByUID(DashboardsFolder DashboardsFolders) {
	endpoint := "/api/dashboards/uid/"
	// 解析 URL
	u, err := url.Parse(grafanaURL)
	if err != nil {
		log.Fatal(err)
	}
	u.Path = path.Join(u.Path, endpoint, DashboardsFolder.UID)

	// 发送请求
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, u.String(), nil)

	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", grafanaAPI))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	value := gjson.Get(string(body), "dashboard").String()

	if len(DashboardsFolder.FolderTitle) == 0 {
		DashboardsFolder.FolderTitle = "general"
	}

	// 将内容写入文件
	fileName := fmt.Sprintf("%v-%v.json", DashboardsFolder.FolderTitle, DashboardsFolder.Title)
	filePath := path.Join(bakDir, fileName)

	if err := ioutil.WriteFile(filePath, []byte(value), 0644); err != nil {
		log.Println("写入文件失败: %v", err)
	}
}

// DashboardsFolders 搜索所有仪表板（返回所有自定义文件夹、仪表板）
type DashboardsFolders struct {
	ID          int    `json:"id"`
	UID         string `json:"uid"`
	Title       string `json:"title"`
	URI         string `json:"uri"`
	URL         string `json:"url"`
	Slug        string `json:"slug"`
	Type        string `json:"type"`
	IsStarred   bool   `json:"isStarred"`
	FolderID    int    `json:"folderId,omitempty"`
	FolderUID   string `json:"folderUid,omitempty"`
	FolderTitle string `json:"folderTitle,omitempty"`
	FolderURL   string `json:"folderUrl,omitempty"`
}

// 搜索所有仪表板（返回所有自定义文件夹、仪表板）
func searchAll() []DashboardsFolders {
	endpoint := "/api/search/"

	// 解析 URL
	u, err := url.Parse(grafanaURL)
	if err != nil {
		log.Fatal(err)
	}
	u.Path = path.Join(u.Path, endpoint)

	// 发送请求
	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, u.String(), nil)

	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", grafanaAPI))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// 解析到结构体
	var DashboardsFolderser []DashboardsFolders
	if err := json.Unmarshal(body, &DashboardsFolderser); err != nil {
		log.Fatalf("解析失败: %v", err)
	}

	return DashboardsFolderser
}
