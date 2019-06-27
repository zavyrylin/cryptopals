package s01

import (
	"testing"
)

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
