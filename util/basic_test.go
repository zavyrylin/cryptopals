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
		got, err := HexToBase64([]byte(test.in))

		if err != nil {
			t.Errorf("HexToBase64(%q) -> err = %q", test.in, err)
			continue
		}

		if bytes.Equal(got, []byte(test.out)) == false {
			t.Errorf("HexToBase64(%q) = %q, want %q", test.in, string(got), test.out)
		}
	}
}
