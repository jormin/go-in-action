package matchers

import (
	"github.com/jormin/go-in-action/chapter2/search"
)

// Default 默认匹配器
type Default struct {
}

// Search 搜索
func (d Default) Search(feed *search.Feed, searchTerm string) ([]*search.Result, error) {
	return nil, nil
}
