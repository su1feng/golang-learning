package ispalindrome

import "strings"

func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	var filtered strings.Builder
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			filtered.WriteByte(s[i])
		}
	}
	s = filtered.String()
	n := len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}
	return true
}
