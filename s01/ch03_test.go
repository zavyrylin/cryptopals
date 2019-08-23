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

func TestTryDecryptXored(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{
			"cryptopals s01ch03",
			"1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736",
			"Cooking MC's like a pound of bacon",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TryDecryptXored(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("TryDecryptXored() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TryDecryptXored() = %q, want %q", got, tt.want)
			}
		})
	}
}
