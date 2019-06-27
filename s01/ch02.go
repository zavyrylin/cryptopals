package s01

import (
	hex "encoding/hex"
	"errors"
)

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
