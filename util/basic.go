package util

import (
	enc "encoding/base64"
	hex "encoding/hex"
)

func HexToBase64(s string) string {
	src := []byte(s)
	dst := make([]byte, hex.DecodedLen(len(src)))
	_, err := hex.Decode(dst, src)
	if err != nil {
		panic(err)
	}
	return enc.StdEncoding.EncodeToString(dst)
}
