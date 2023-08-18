package funcs

import (
	"math/rand"
	"strings"
)

var (
	symbols        = "!@#$%^&*()_+[]{}|;':,.<>?/\\`~"
	numbers        = "0123456789"
	lowerAlphabets = "abcdefghijklmnopqrstuvwxyz"
	upperAlphabets = strings.ToUpper(lowerAlphabets)

	allChars = symbols + numbers + lowerAlphabets + upperAlphabets
)

func PassGen(l int) string {
	password := make([]byte, l)

	for i := 0; i < l; i++ {
		randIndex := rand.Intn(len(allChars))
		password[i] = allChars[randIndex]
	}
	return string(password)
}

func Atoi(s string) int {
	n := 0
	for _, ch := range s {
		n = n*10 + int(ch-'0')
	}
	return n
}
