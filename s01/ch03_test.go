package s01

import (
	"math"
	"testing"
)

func TestEnglishScore(t *testing.T) {
	tests := []struct {
		s   string
		thr float64
	}{
		{"English!", 100.0},
		{"English, motherfucker, do you speak it?", 100.0},
		{"Английский, ублюдок, говоришь на нём?", 100.0},
		{"Inglés, hijo de puta, ¿lo hablas?", 100.0},
		{"英语，妈妈，你会说吗?", 100.0},
	}
	for _, test := range tests {
		score := EnglishScore(test.s)
		if score > test.thr || math.IsInf(score, 1.0) {
			t.Errorf("EnglishScore(%q) = %v, want > %v",
				test.s, score, test.thr)
		}
	}
}
