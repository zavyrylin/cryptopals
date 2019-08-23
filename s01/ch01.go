package s01

import "cryptopals/aux"

func HexToBase64(src []byte) ([]byte, error) {
	dst1, err := aux.HexDecodeBytes(src)
	if err != nil {
		return nil, err
	}
	return aux.Base64EncodeBytes(dst1), nil
}

func HexToBase64String(s string) (string, error) {
	ret, err := HexToBase64([]byte(s))
	return string(ret), err
}
