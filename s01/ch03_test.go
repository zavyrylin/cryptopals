package s01

import (
	"testing"
)

func TestEnglishScore(t *testing.T) {
	tests := []struct {
		s   string
		thr float64
	}{
		{"English!", 10.0},
		{"English, motherfucker, do you speak it?", 10.0},
		{"Английский, ублюдок, говоришь на нём?", 100.0},
		{"Inglés, hijo de puta, ¿lo hablas?", 20.0},
		{"英语，妈妈，你会说吗?", 100.0},
	}
	for _, test := range tests {
		score := EnglishScore(test.s)
		if score < test.thr {
			t.Errorf("EnglishScore(%q) = %v, want > %v",
				test.s, score, test.thr)
		}
	}
}

func TestCompareEnglishScore(t *testing.T) {
	tests := []struct {
		lesserScoreText, greaterScoreText string
	}{
		{"Cooking MC's like a pound of bacon", "\x06**.,+\"e\b\x06b6e),. e$e5*0+!e*#e'$&*+"},
		{"Cooking MC's like a pound of bacon", "cOOKING\x00mc'S\x00LIKE\x00A\x00POUND\x00OF\x00BACON"},
	}

	for _, test := range tests {
		ls := EnglishScore(test.lesserScoreText)
		gs := EnglishScore(test.greaterScoreText)
		if ls > gs {
			t.Errorf("EnglishScore(%+q)=%v < EnglishScore(%+q)=%v, want the opposite",
				test.lesserScoreText, ls,
				test.greaterScoreText, gs)
		}
	}
}

func TestDecryptXored(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			"1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736",
			"Cooking MC's like a pound of bacon",
		},
	}
	for _, tt := range tests {
		got, err := DecryptXored(tt.input)
		if err != nil {
			t.Errorf("TryDecryptXored() error = %v", err)
			return
		}
		if got.Decrypted != tt.want {
			t.Errorf("TryDecryptXored() = %q, want %q", got.Decrypted, tt.want)
		}
	}
}
