package iteration

func Repeat(s string, count int) string {
	repeated := ""
	for i := 0; i < count; i++ {
		repeated += s
	}
	return repeated
}
