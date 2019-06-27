package s01

import (
	"testing"
)

func TestEnglishScore(t *testing.T) {
	tests := []struct {
		s   string
		thr float64
	}{
		{"English!", 0.05},
		{"English, motherfucker, do you speak it?", 0.9},
		{"Английский, ублюдок, говоришь на нём?", 0.0},
		{"Inglés, hijo de puta, ¿lo hablas?", 0.0},
		{"英语，妈妈，你会说吗?", 0.0},
	}
	for _, test := range tests {
		score := EnglishScore(test.s)
		if score < test.thr {
			t.Errorf("EnglishScore(%q) = %v, want > %v",
				test.s, score, test.thr)
		}
	}
}
