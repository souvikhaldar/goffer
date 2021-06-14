package utils

func ReverseStr(input string) string {
	var out []rune
	l := len(input)

	for i, _ := range input {
		out = append(out, rune(input[l-i-1]))
	}
	return string(out)
}
