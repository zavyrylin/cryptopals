package util

import (
	"bytes"
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
		input := []byte(test.in)
		ret := HexToBase64(input)
		expected := []byte(test.out)
		if false == bytes.Equal(expected, ret) {
			t.Errorf(
				"HexToBase64(%q) = %q, want %q",
				string(test.in),
				string(ret),
				test.out)
		}
	}
}
