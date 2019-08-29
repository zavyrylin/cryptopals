package s01

import (
	"math/bits"
)

// HammingDistance calculates Hamming distance between to strings.
// Note: s1 and s2 must be of the same length, otherwise function returns -1.
// See https://en.wikipedia.org/wiki/Hamming_distance
func HammingDistance(s1, s2 string) int {
	if len(s1) != len(s2) {
		return -1
	}
	bs1 := []byte(s1)
	bs2 := []byte(s2)
	var dist int
	for idx, b1 := range bs1 {
		diff := b1 ^ bs2[idx]
		dist += bits.OnesCount8(diff)
	}
	return dist
}
