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

func RemoveSpaces(cont string) string {
	return strings.Replace(cont, " ", "", -1)
}