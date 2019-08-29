package s01

import (
	"encoding/hex"
)

// modifies src bytearray
func xorRepeatedly(src, xor []byte) {
	l := len(xor)
	for i, b := range src {
		src[i] = b ^ xor[i%l]
	}
}

// EncryptStringWithRepeatedXor performs repeatable xor encryption of src
// using given xor string. Returns HEX-ed string.
func EncryptStringWithRepeatedXor(src, xor string) string {
	bxor := []byte(xor)
	bs := []byte(src)
	xorRepeatedly(bs, bxor)
	ret := hex.EncodeToString(bs)
	return ret
}
