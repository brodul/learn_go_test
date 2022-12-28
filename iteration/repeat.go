package iteration

func Repeat(char string) string {
	repeated := ""
	for i := 0; i < 5; i++ {
		repeated = repeated + char
	}
	return repeated
}
