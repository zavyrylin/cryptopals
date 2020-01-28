package s01

import (
	"cryptopals/cptest"
	"testing"
)

func TestFindXoredLine(t *testing.T) {
	want := "Now that the party is jumping\n"
	file := cptest.FixtureFile("s01e04.txt")
	defer file.Close()
	got, err := FindXoredLine(file)
	if err != nil {
		t.Errorf("FindXoredLine() error = %v", err)
		return
	}
	if got.Decrypted != want {
		t.Errorf("FindXoredLine() -> %+q, want %+q", got.Decrypted, want)
	}
}
