package s01

import (
	enc "encoding/base64"
	hex "encoding/hex"
)

// general functions

func unhex(xs []byte) ([]byte, error) {
	dst := make([]byte, hex.DecodedLen(len(xs)))
	if _, err := hex.Decode(dst, xs); err != nil {
		return nil, err
	}
	return dst, nil
}

func base64(xs []byte) []byte {
	e := enc.StdEncoding
	dst := make([]byte, e.EncodedLen(len(xs)))
	e.Encode(dst, xs)
	return dst
}

func HexToBase64(src []byte) ([]byte, error) {
	dst1, err := unhex(src)
	if err != nil {
		return nil, err
	}
	return base64(dst1), nil
}

func HexToBase64String(s string) (string, error) {
	ret, err := HexToBase64([]byte(s))
	return string(ret), err
}
