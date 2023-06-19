package pointerAndErrorHandling

import "strings"

func Caesar(offset int, input string) string {
	geser := ""
	text := strings.ToLower(input)
	for _, char := range text {
		geser += string(rune(offset)%26 + char)
	}
	return geser
}
