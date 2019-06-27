package s01

import (
	"testing"
)

var base64Tests = []struct {
	in, out string
	ok      bool
}{
	{"abcdefgh", "", false},
	{"01020", "", false},
	{"", "", true},
	{
		"49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d",
		"SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t",
		true,
	},
}

func TestHexToBase64String(t *testing.T) {
	for _, test := range base64Tests {
		got, err := HexToBase64String(test.in)
		if err != nil && test.ok {
			t.Errorf("HexToBase64String(%q) -> err = %q", test.in, err)
			continue
		}
		if got != test.out {
			t.Errorf("HexToBase64(%q) = %q, want %q", test.in, got, test.out)
		}
	}
}
