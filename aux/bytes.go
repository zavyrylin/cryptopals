package aux

import (
	"encoding/base64"
	"encoding/hex"
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
	e := base64.StdEncoding
	dst := make([]byte, e.EncodedLen(len(xs)))
	e.Encode(dst, xs)
	return dst
}

// Base64DecodeBytes decodes given bytes using Base64 decoder.
func Base64DecodeBytes(xs []byte) []byte {
	e := base64.StdEncoding
	dst := make([]byte, e.DecodedLen(len(xs)))
	e.Decode(dst, xs)
	return dst
}
