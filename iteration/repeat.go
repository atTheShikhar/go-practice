package iteration

func Repeat(charater string, num int) string {
	var repeated string
	for i := 0; i < num; i++ {
		repeated += charater
	}
	return repeated
}
