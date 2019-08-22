package s01

// Let's pretend to look more scientific ;)
//
// To qualify a given text as an English text, we have to somehow calculate
// so called 'score' for that text and correlate it with some numerical bounds. Scores that fall
// inside those bounds will say us that the text is in English with a certain probability.
//
// It feels natural to exploit letter frequencies in our 'Englishness' analysis. To be scientific
// to some extent but not too much ;) we can assume that frequencies of different letters are
// completely independent of each other, therefore we got 26 independent standard normal random
// variables. We will use any of statistical tests for that.
//
// Here we use chi-squared test (https://en.wikipedia.org/wiki/Chi-squared_test) for analysis.
// Our null hypothesis will be the following: a given text is not an English text.
// Chi-square statistics is constructed through sum of squared errors. An error for each letter
// is a difference between observed count this letter appears in the text and 'theoretical' count
// this letter expected to appear.
// The lower this sum is the more probable that the text is in English.
// It that's simple ;)

import (
	"math"
	"unicode"
)

var (
	// Frequency values are taken from
	// https://en.wikipedia.org/wiki/Letter_frequency#Relative_frequencies_of_letters_in_the_English_language
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
		'w': 0.02360,
		'f': 0.02228,
		'g': 0.02015,
		'y': 0.01974,
		'p': 0.01929,
		'b': 0.01492,
		'v': 0.00978,
		'k': 0.00772,
		'j': 0.00153,
		'x': 0.00150,
		'q': 0.00095,
		'z': 0.00074,
	}
)

func xorRunesAgainstRune(rs []rune, r rune) []rune {
	for i, r0 := range rs {
		rs[i] = r0 ^ r
	}
	return rs
}

// XorStringAgainstRune applies xor operation of every rune in a given
// string against a given rune.
func XorStringAgainstRune(s string, r rune) string {
	rs := []rune(s)
	ret := xorRunesAgainstRune(rs, r)
	return string(ret)
}

// EnglishScore calculates a chi-square (score) statistics of a given text.
func EnglishScore(s string) float64 {
	counts := make(map[rune]float64, len(engRuneFreqs))
	var l float64

	// Count english letters in s.
	for _, r0 := range s {
		if !unicode.IsLetter(r0) {
			continue
		}
		l++
		r := unicode.ToLower(r0)
		f := engRuneFreqs[r]
		if f > 0 {
			counts[r] += 1.0
		}
	}

	// The score for a text that contains no letters.
	if l == 0.0 {
		return math.Inf(1.0)
	}

	// Calculate chi square statistics of a given string.
	score := 0.0
	for k, e := range engRuneFreqs {
		observed := counts[k]
		expected := l * e
		score += math.Pow(observed-expected, 2) / expected
	}
	return score
}
