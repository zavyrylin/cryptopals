package util

import (
	enc "encoding/base64"
	hex "encoding/hex"
	"errors"
)

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

// s01ch01
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

func xor(xs, ys []byte) ([]byte, error) {
	xlen := len(xs)
	if xlen != len(ys) {
		return nil, errors.New("input arrays' lengths are not equal")
	}
	ret := make([]byte, xlen)
	for i, x := range xs {
		ret[i] = x ^ ys[i]
	}
	return ret, nil
}

// s01ch02
func XorString(s1, s2 string) (string, error) {
	var xs, ys []byte
	var err error
	if xs, err = hex.DecodeString(s1); err != nil {
		return "", err
	}
	if ys, err = hex.DecodeString(s2); err != nil {
		return "", err
	}
	var zs []byte
	if zs, err = xor(xs, ys); err != nil {
		return "", err
	}
	return hex.EncodeToString(zs), nil
}
