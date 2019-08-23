package aux

import (
	enc "encoding/base64"
	hex "encoding/hex"
)

// HexDecodeBytes decodes given bytes using hex decoder.
func HexDecodeBytes(xs []byte) ([]byte, error) {
	dst := make([]byte, hex.DecodedLen(len(xs)))
	if _, err := hex.Decode(dst, xs); err != nil {
		return nil, err
	}
	return dst, nil
}

// Base64EncodeBytes encodes given bytes using Base64 encoder.
func Base64EncodeBytes(xs []byte) []byte {
	e := enc.StdEncoding
	dst := make([]byte, e.EncodedLen(len(xs)))
	e.Encode(dst, xs)
	return dst
}
