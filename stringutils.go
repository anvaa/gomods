package stringutils

import (
	"strings"
)

func CountWords(cont string) int {
	return strings.Count(cont, " ") + 1
}

func CountChars(cont string) int {
	return len(cont)
}
