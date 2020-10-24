package main

import (
	"github.com/jormin/go-in-action/chapter2/matchers"
	_ "github.com/jormin/go-in-action/chapter2/matchers"
	"github.com/jormin/go-in-action/chapter2/search"
	"gitlab.wcxst.com/jormin/go-tools/log"
)

func init() {
	// 设置日志前缀
	log.SetPrefix("Go In Action")
}

func main() {
	// 搜索器
	searcher := search.Searcher{
		Matchers: map[string]search.Matcher{},
	}
	searcher.Register(search.DefaultMatcher, matchers.Default{})
	searcher.Register(search.RssMatcher, matchers.Rss{})
	// 搜索
	searcher.Run("film")
}
