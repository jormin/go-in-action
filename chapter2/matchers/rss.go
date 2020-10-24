package matchers

import (
	"encoding/xml"
	"fmt"
	"github.com/jormin/go-in-action/chapter2/search"
	"gitlab.wcxst.com/jormin/go-tools/log"
	"net/http"
	"regexp"
)

type (
	rssDocument struct {
		XMLName xml.Name `xml:"rss"`
		Channel channel  `xml:"channel"`
	}
	channel struct {
		XMLName       xml.Name `xml:"channel"`
		Title         string   `xml:"title"`
		Link          string   `xml:"link"`
		Description   string   `xml:"description"`
		Language      string   `xml:"language"`
		Copyright     string   `xml:"copyright"`
		Generator     string   `xml:"generator"`
		LastBuildDate string   `xml:"lastBuildDate"`
		Image         image    `xml:"image"`
		Item          []item   `xml:"item"`
	}
	image struct {
		XMLName xml.Name `xml:"image"`
		URL     string   `xml:"url"`
		Title   string   `xml:"title"`
		Link    string   `xml:"link"`
	}
	item struct {
		XMLName     xml.Name `xml:"item"`
		Title       string   `xml:"title"`
		Description string   `xml:"description"`
		PubDate     string   `xml:"pubDate"`
		Link        string   `xml:"link"`
		Guid        string   `xml:"guid"`
	}
)

// Rss Rss 匹配器
type Rss struct {
}

// Search 搜索
func (r Rss) Search(feed *search.Feed, searchTerm string) ([]*search.Result, error) {
	log.Info("Search [%s] use feed type [%s] on site [%s] with uri [%s]", searchTerm, feed.Type, feed.Name, feed.URI)
	var results []*search.Result
	// 获取 Uri 数据
	doc, err := r.Retrieve(feed)
	if err != nil {
		log.Error("Search [%s] use feed type [%s] on site [%s] with uri [%s] error: %+v", searchTerm, feed.Type, feed.Name, feed.URI, err)
		return nil, err
	}
	for _, item := range doc.Channel.Item {
		matched, _ := regexp.MatchString(searchTerm, item.Title)
		if matched {
			results = append(results, &search.Result{
				Field:   "title",
				Content: item.Title,
				Feed:    feed,
			})
		}
		matched, _ = regexp.MatchString(searchTerm, item.Description)
		if matched {
			results = append(results, &search.Result{
				Field:   "description",
				Content: item.Description,
				Feed:    feed,
			})
		}
	}
	return results, nil
}

// Retrieve 获取并解析 rss 文档
func (r Rss) Retrieve(feed *search.Feed) (*rssDocument, error) {
	// 检测要读取的 url 是否为空
	if feed.URI == "" {
		return nil, nil
	}
	// 读取 url 内容
	req, err := http.NewRequest(http.MethodGet, feed.URI, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.80 Safari/537.36")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	// 关闭 http 响应
	defer resp.Body.Close()
	// 判断响应状态码是否为200
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Http Response Error: %d\n", resp.StatusCode)
	}
	// xml 解析响应内容
	document := &rssDocument{}
	err = xml.NewDecoder(resp.Body).Decode(document)
	return document, err
}
