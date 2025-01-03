package iteration

func Repeat(word string) string {
	repeated := ""
	for i := 0; i < 4; i++ {
		repeated += word
	}
	return repeated
}
