package util

import (
	"math"
	"unicode"
)

// see https://en.wikipedia.org/wiki/Letter_frequency#Relative_frequencies_of_letters_in_the_English_language
var (
	engRuneFreqs = map[rune]float64{
		'e': 0.12702,
		't': 0.09056,
		'a': 0.08167,
		'o': 0.07507,
		'i': 0.06966,
		'n': 0.06749,
		's': 0.06327,
		'h': 0.06094,
		'r': 0.05987,
		'd': 0.04253,
		'l': 0.04025,
		'c': 0.02782,
		'u': 0.02758,
		'm': 0.02406,
	}
)

func EnglishScore(s string) float64 {
	freq := make(map[rune]float64, len(engRuneFreqs))
	var l float64

	for _, r0 := range s {
		if !unicode.IsLetter(r0) {
			continue
		}
		l++
		r := unicode.ToLower(r0)
		f := engRuneFreqs[r]
		if f > 0 {
			freq[r] += 1.0
		}
	}

	if l == 0 {
		return 0.0
	}

	// calculate frequency for each rune and sum of their frequencies
	var score float64
	for k, v := range freq {
		w := engRuneFreqs[k]
		diff := w - (v / l)
		score += diff * diff
	}

	return 1.0 - math.Sqrt(score)/float64(len(engRuneFreqs))
}
