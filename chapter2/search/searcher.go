package search

import (
	"gitlab.wcxst.com/jormin/go-tools/log"
	"sync"
)

type Searcher struct {
	Matchers map[string]Matcher
}

// Run 搜索方法入口
func (s *Searcher) Run(searchTerm string) {
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal("Retrieve feeds error: %+v", err)
	}

	// 存放结果的无缓冲通道
	results := make(chan *Result)
	// 使用 WaitGroup 来跟踪 goroutine
	wg := sync.WaitGroup{}
	wg.Add(len(feeds))
	// 遍历搜索数据源
	for _, feed := range feeds {
		// 为数据源指定匹配器
		matcher, exists := s.Matchers[feed.Type]
		if !exists {
			// 如果注册的匹配器集合中没有特定的匹配器，则使用默认的匹配器
			matcher = s.Matchers[DefaultMatcher]
		}
		// 开启一个 goroutine 来处理匹配
		go func(matcher Matcher, feed *Feed) {
			// 匹配
			Match(matcher, feed, searchTerm, results)
			// 匹配完成后完成这个 goroutine
			wg.Done()
		}(matcher, feed)
	}
	// 启动一个监听匹配结果的 goroutine
	go func() {
		// 监听所有 goroutine
		wg.Wait()
		// 所有 goroutine 都完成后，关闭结果通道
		close(results)
	}()
	// 显示匹配结果
	Display(results)
}

// Register 注册匹配器
func (s *Searcher) Register(feedType string, matcher Matcher) {
	if _, exists := s.Matchers[feedType]; exists {
		log.Fatal("Matcher %s already registered", feedType)
	}
	log.Info("Register matcher %s successful", feedType)
	s.Matchers[feedType] = matcher
}
