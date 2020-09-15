package iteration

// Repeat takes string and integer and returns it repeated n times based on the integer.
func Repeat(character string, repeat int) string {
	var repeated string
	for i := 0; i < repeat; i++ {
		repeated += character
	}
	return repeated
}
