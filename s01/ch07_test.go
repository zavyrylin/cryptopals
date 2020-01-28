package s01

import (
	"cryptopals/cptest"
	"strings"
	"testing"
)

func TestAESDecryptECB(t *testing.T) {
	input := cptest.FixtureContentsBase64("s01e07.txt")
	want := "I'm back and I'm ringin' the bell"
	got, err := AESDecryptECB(input, []byte("YELLOW SUBMARINE"))

	if err != nil {
		t.Errorf("AESDecryptECB() error = %v", err)
		return
	}

	if strings.Index(string(got), want) != 0 {
		t.Errorf("AESDecryptECB() = %v, want %v", string(got), string(want))
	}
}
