package search

// Result 匹配结果
type Result struct {
	Field   string
	Content string
	Feed    *Feed
}

// Matcher 匹配器接口
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

// 匹配
func Match(matcher Matcher, feed *Feed, searchItem string, results chan<- *Result) {
	// 使用匹配器进行搜索
	sr, err := matcher.Search(feed, searchItem)
	if err != nil {
		return
	}
	// 将搜索结果依次写入通道
	for _, result := range sr {
		results <- result
	}
}
