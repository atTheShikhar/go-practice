package iteration

func Repeat(charater string) string {
	var repeated string
	for i := 0; i < 5; i++ {
		repeated = repeated + charater
	}
	return repeated
}
