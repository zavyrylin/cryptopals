package s01

// Let's pretend to look more scientific ;)
//
// To qualify a given text as an English text, we have to somehow calculate
// so called 'score' for that text and correlate it with some numerical bounds. Scores that fall
// inside those bounds will say us that the text is in English with a certain probability.
//
// It feels natural to exploit letter frequencies in our 'Englishness' analysis. To be scientific
// to some extent but not too much ;) we can assume that frequencies of different letters are
// completely independent of each other, therefore we kind of got 27 (26+1 for space character) independent
// standard normal random variables. We can use any of existing statistical tests for that.
//
// Here we use chi-squared test (https://en.wikipedia.org/wiki/Chi-squared_test) for analysis.
// Our null hypothesis will be the following: there is no distinction between a calculated character frequency
// and theoretical letter frequency in English text.
// Chi-square statistics is constructed through sum of squared errors. An error for each letter
// is a difference between observed count this letter appears in the text and 'theoretical' count
// this letter expected to appear.
// The lower this sum is the more probable that the text is in English.
// It that's simple ;)

import (
	"cryptopals/aux"
	"math"
	"strings"
	"unicode"
	"unicode/utf8"
)

var (
	// Frequency values are taken from
	// https://web.archive.org/web/20170918020907/http://www.data-compression.com/english.html
	engRuneFreqs = map[rune]float64{
		'a': 0.0651738,
		'b': 0.0124248,
		'c': 0.0217339,
		'd': 0.0349835,
		'e': 0.1041442,
		'f': 0.0197881,
		'g': 0.0158610,
		'h': 0.0492888,
		'i': 0.0558094,
		'j': 0.0009033,
		'k': 0.0050529,
		'l': 0.0331490,
		'm': 0.0202124,
		'n': 0.0564513,
		'o': 0.0596302,
		'p': 0.0137645,
		'q': 0.0008606,
		'r': 0.0497563,
		's': 0.0515760,
		't': 0.0729357,
		'u': 0.0225134,
		'v': 0.0082903,
		'w': 0.0171272,
		'x': 0.0013692,
		'y': 0.0145984,
		'z': 0.0007836,
		' ': 0.1918182,
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

func frequencyAnalysis(s string) map[rune]float64 {
	fa := make(map[rune]float64, len(engRuneFreqs)+1)
	s0 := strings.ToLower(s)
	for _, r := range s0 {
		if engRuneFreqs[r] > 0 {
			fa[r] += 1.0
		} else {
			fa[unicode.MaxRune] += 1.0
		}
		// cnt := fa[r]
		// if cnt > 0 {
		// 	fa[r] = cnt + 1.0
		// } else {
		// 	fa[r] = 1.0
		// }
	}
	return fa
}

func chiSquare(fa map[rune]float64, len float64) float64 {
	cs := 0.0
	for r, observed := range fa {
		expected := engRuneFreqs[r]
		if expected > 0 {
			expected *= len
		} else {
			// A slight customization of Chi-Square.
			// If expected count is zero, the statistic will be +Inf which is useless.
			// Furthermore, if we leave out all non-English letters for calculation, we may get
			// surprisingly low chi-square statistic for gibberish strings -- even lower,
			// than usual English text might give.
			// By assigning insignificantly small value to expected count for all non-letter runes
			// we get noticeably higher values of chi-square for non-text strings or for mixed strings
			// consist of letter-like symbols and other bytes.
			// This value should not be too small, otherwise even punctuation runes will give the value
			// indistinguishable from that of what string of random bytes gives.
			expected = 0.01
		}
		cs += math.Pow(observed-expected, 2) / expected
	}
	return cs
}

// EnglishScore calculates a chi-square statistic for a given text.
// The lower the result the more probable that s is in English.
func EnglishScore(s string) float64 {
	fa := frequencyAnalysis(s)
	if len(fa) == 0 {
		return math.MaxFloat64
	}
	l := float64(utf8.RuneCountInString(s))
	chi := chiSquare(fa, l)
	return chi
}

func xorBytesAgainstByte(bs []byte, b byte) []byte {
	ret := make([]byte, len(bs))
	for i, b0 := range bs {
		ret[i] = b0 ^ b
	}
	return ret
}

type DecryptResult = struct {
	Score     float64
	Decrypted string
}

func DecryptXored(s string) (*DecryptResult, error) {
	var bs []byte
	var err error
	if bs, err = aux.HexDecodeBytes([]byte(s)); err != nil {
		return nil, err
	}

	var found byte
	min := math.MaxFloat64
	for i := 0; i <= 255; i++ {
		b := byte(i)
		xored := xorBytesAgainstByte(bs, b)
		score := EnglishScore(string(xored))

		if score < min {
			min = score
			found = b
		}
	}
	src := xorBytesAgainstByte(bs, found)
	decrypted := string(src)
	ret := DecryptResult{Score: min, Decrypted: decrypted}
	return &ret, nil
}
