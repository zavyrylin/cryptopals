package util

import (
	"testing"
)

var base64Cases = []struct {
	in, out string
}{
	{"",
		""},
	{"49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d",
		"SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"},
}

func TestBase64(t *testing.T) {
	for _, test := range base64Cases {
		res := HexToBase64(test.in)
		if res != test.out {
			t.Errorf("HexToBase64(%q) = %q, want %q", test.in, res, test.out)
		}
	}
}
