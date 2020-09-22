package iteration

func IterateText(text string, times int) string {
	var repeated string
	for i := 0; i < times; i++ {
		repeated += text
	}
	return repeated
}
