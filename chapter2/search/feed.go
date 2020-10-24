package search

import (
	"encoding/json"
	"os"
)

// Feed 搜索数据源
type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

var dataFile = "data/data.json"

// RetrieveFeeds 获取Feed数据
func RetrieveFeeds() ([]*Feed, error) {
	// 打开文件
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}
	// 关闭文件
	defer file.Close()
	// 解析数据
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)
	return feeds, err
}
