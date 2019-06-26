package util

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

var xorTests = []struct {
	in1, in2, exp string
	ok            bool
}{
	{"", "", "", true},
	{"123", "", "", false},
	{"1234", "123", "", false},
	{"", "123", "", false},
	{"", "1234", "", false},
	{
		"1c0111001f010100061a024b53535009181c",
		"686974207468652062756c6c277320657965",
		"746865206b696420646f6e277420706c6179",
		true,
	},
}

func TestXorString(t *testing.T) {
	for _, test := range xorTests {
		got, err := XorString(test.in1, test.in2)
		if err != nil && test.ok {
			t.Errorf("Xor(%q, %q) -> err = %q", test.in1, test.in2, err)
			continue
		}
		if got != test.exp {
			t.Errorf("Xor(%q, %q) = %q, want %q",
				test.in1, test.in2, string(got), test.exp)
		}
	}
}

var xorStringTests = []struct {
	in  string
	r   rune
	out string
}{
	{"", 'a', ""},
	{"abcd", 'a', "\x00\x03\x02\x05"},
	{"abcd", 'x', "\x19\x1a\x1b\x1c"},
}

func TestXorStringAgainstRune(t *testing.T) {
	for _, test := range xorStringTests {
		ret := XorStringAgainstRune(test.in, test.r)
		if ret != test.out {
			t.Errorf("XorStringAgainstRune(%q, %q) = %q, want %q",
				test.in, test.r, ret, test.out)
		}
	}
}
