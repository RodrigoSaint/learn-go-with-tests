package iteration

func Repeat(word string, limit int) string {
	repeated := ""
	for i := 0; i < limit; i++ {
		repeated += word
	}
	return repeated
}
