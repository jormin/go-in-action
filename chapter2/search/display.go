package search

import "gitlab.wcxst.com/jormin/go-tools/log"

// Display 展示搜索结果
func Display(results chan *Result) {
	for result := range results {
		log.Info("Got Result: Site: %s, URI: %s, Field: %s, Content: %s", result.Feed.Name, result.Feed.URI, result.Field, result.Content)
	}
}
