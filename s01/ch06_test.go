package s01

import (
	s "strings"
	"testing"
)

func TestHammingDistance(t *testing.T) {
	tests := []struct {
		s1, s2 string
		want   int
	}{
		{"a", "a", 0},
		{"a", "b", 2},
		{"a", "c", 1},
		{"a", "ab", -1},
		{"ab", "ab", 0},
		{"ab", "ac", 1},
		{"ab", "ad", 2},
		{"ab", "ae", 3},
		{"ab", "ad", 2},
		{"abc", "bca", 4},
		{s.Repeat("a", 100), s.Repeat("b", 100), 200},
		{s.Repeat("ab", 100), s.Repeat("ae", 100), 300},
		{"this is a test", "wokka wokka!!!", 37},
	}

	for _, tt := range tests {
		if got := HammingDistance(tt.s1, tt.s2); got != tt.want {
			t.Errorf("HammingDistance(%q, %q) = %v, want %v", tt.s1, tt.s2, got, tt.want)
		}
	}
}
