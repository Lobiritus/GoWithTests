package iteration

// Repeats characters as many times as specified by the second argument.
func Repeat(character string, count int) string {
	var repeated string
	for i := 0; i < count; i++ {
		repeated += character
	}
	return repeated
}
