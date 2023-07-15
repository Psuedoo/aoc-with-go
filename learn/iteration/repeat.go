package iteration

func Repeat(char string, repeatCount int) (repeated string) {
	for i := 0; i < repeatCount; i++ {
		repeated += char
	}
	return
}
