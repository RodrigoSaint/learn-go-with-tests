package iteration

import "strings"

func Repeat(word string, limit int) string {
	repeated := ""
	for i := 0; i < limit; i++ {
		repeated += word
	}
	return repeated
}

func RepeatBuilder(word string, limit int) string {
	return strings.Repeat(word, limit)
}
