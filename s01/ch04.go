package s01

import (
	"bufio"
	"math"
	"os"
)

type Finding = struct {
	LineNo         int
	Src, Decrypted string
}

// FindXoredLine returns line number and
func FindXoredLine(file *os.File) (Finding, error) {
	scanner := bufio.NewScanner(file)
	ret := Finding{LineNo: -1, Src: "", Decrypted: ""}

	var i int
	result := &DecryptResult{Score: math.MaxFloat64}
	for scanner.Scan() {
		i++
		src := scanner.Text()
		found, err := DecryptXored(src)
		if err != nil {
			return ret, err
		}

		if found.Score < result.Score {
			result = found
			ret.LineNo = i
			ret.Src = src
			ret.Decrypted = found.Decrypted
		}
	}
	return ret, nil
}
