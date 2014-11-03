package string

import "strings"

func Reverse(s string) string {
	b := []rune(s)
	for i := 0; i < len(b)/2; i++ {
		j := len(b) - i - 1
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func IsPangram(s string) string {
	b := []byte(strings.ToLower(s))

	var chars map[byte]bool
	chars = make(map[byte]bool)

	for i := 0; i < len(b); i++ {
		chars[b[i]] = true
	}

	if len(chars) == 27 {
		return "pangram"
	} else {
		return "not pangram"
	}

}
