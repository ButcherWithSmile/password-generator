// This package imports the necessary libraries for generating passwords and converting strings to integers

package funcs

import (
	"math/rand"
	"strings"
)

// These variables store the different types of characters that can be used to generate passwords
var (
	symbols        = "!@#$%^&*()_+[]{}|;':,.<>?/\\`~"
	numbers        = "0123456789"
	lowerAlphabets = "abcdefghijklmnopqrstuvwxyz"
	upperAlphabets = strings.ToUpper(lowerAlphabets)

	allChars = symbols + numbers + lowerAlphabets + upperAlphabets
)

// This function generates a random password of the specified length
// It does this by randomly selecting characters from the allChars variable
func PassGen(l int) string {
	password := make([]byte, l)

	for i := 0; i < l; i++ {
		randIndex := rand.Intn(len(allChars))
		password[i] = allChars[randIndex]
	}
	return string(password)
}

// This function converts a string to an integer
// It does this by iterating over the characters in the string 
// multiplying the current value of n by 10 and adding the ASCII value of the current character
func Atoi(s string) int {
	n := 0
	for _, ch := range s {
		n = n*10 + int(ch-'0')
	}
	return n
}
