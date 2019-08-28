package s01

import (
	"os"
	"testing"
)

var samplesDir = os.Getenv("SAMPLES_DIR")

func openSample(name string) (*os.File, error) {
	if file, err := os.Open("data/" + name); err == nil {
		return file, nil
	} else if os.IsNotExist(err) {
		if file, err := os.Open("../data/" + name); err == nil {
			return file, nil
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}

func TestFindXoredLine(t *testing.T) {
	want := "Now that the party is jumping\n"
	file, _ := openSample("s01e04.txt")
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
