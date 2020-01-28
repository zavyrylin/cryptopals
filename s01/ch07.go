package s01

import (
	"crypto/aes"
	"cryptopals/aux"
)

// AESDecryptECB decrypts ciphertext using AES-CBC with key
func AESDecryptECB(ciphertext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	out := make([]byte, len(ciphertext))
	mode := aux.NewECBDecrypter(block)
	mode.CryptBlocks(out, ciphertext)

	return out, nil
}
