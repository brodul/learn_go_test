package iteration

const repeatCount = 5

func Repeat(char string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += char
	}
	return repeated
}
