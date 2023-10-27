package iteration

func Repeated(char string, repeat int) string {
	var repeated string
	for i := 0; i < repeat; i++ {
		repeated += "a"
	}
	return repeated
}
