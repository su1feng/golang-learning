package ispalindrome

import (
	"strings"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"simple palindrome", "racecar", true},
		{"with spaces", "race car", true},
		{"mixed case", "RaceCar", true},
		{"not palindrome", "hello", false},
		{"empty string", "", true},
		{"single character", "a", true},
		{"long palindrome", "A man a plan a canal Panama", true},
		{"numbers", "12321", true},
		{"mixed alphanumeric", "A1B2B1A", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsPalindrome(tt.input)
			if got != tt.want {
				t.Errorf("IsPalindrome(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestIsPalindrome_Subtests(t *testing.T) {
	t.Run("simple palindrome", func(t *testing.T) {
		if !IsPalindrome("racecar") {
			t.Error("racecar should be palindrome")
		}
	})

	t.Run("not palindrome", func(t *testing.T) {
		if IsPalindrome("hello") {
			t.Error("hello should not be palindrome")
		}
	})

	t.Run("case insensitive", func(t *testing.T) {
		if !IsPalindrome("Madam") {
			t.Error("Madam should be palindrome (case insensitive)")
		}
	})
}

func BenchmarkIsPalindrome(b *testing.B) {
	testStrings := []string{
		"racecar",
		"A man a plan a canal Panama",
		"1234567890",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, s := range testStrings {
			_ = IsPalindrome(s)
		}
	}
}

func BenchmarkIsPalindrome_LongString(b *testing.B) {
	longString := "A" + strings.Repeat("a", 1000) + "A"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = IsPalindrome(longString)
	}
}
