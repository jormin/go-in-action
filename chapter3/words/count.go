package words

import "strings"

// CountWords 统计字符数
func CountWords(text string) int {
	return len(strings.Fields(text))
}
