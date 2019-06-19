package util

import (
	enc "encoding/base64"
	hex "encoding/hex"
)

func HexToBase64(src []byte) []byte {
	// un-hex
	dst1 := make([]byte, hex.DecodedLen(len(src)))
	_, err := hex.Decode(dst1, src)
	if err != nil {
		panic(err)
	}

	// base64
	e := enc.StdEncoding
	dst2 := make([]byte, e.EncodedLen(len(dst1)))
	e.Encode(dst2, dst1)

	return dst2
}
