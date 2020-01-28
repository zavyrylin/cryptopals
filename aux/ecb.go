package aux

import (
	"crypto/cipher"
)

type ecb struct {
	b         cipher.Block
	blockSize int
}

// NewECBDecrypter returns a BlockMode which decrypts in electronic
// codeblock mode (ECB), using the given Block.
func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
	return &ecb{
		b:         b,
		blockSize: b.BlockSize(),
	}
}

func (x *ecb) BlockSize() int { return x.blockSize }

func (x *ecb) CryptBlocks(dst, src []byte) {
	if len(src)%x.blockSize != 0 {
		panic("crypto/cipher: input not full blocks")
	}
	if len(dst) < len(src) {
		panic("crypto/cipher: output smaller than input")
	}
	for len(src) > 0 {
		x.b.Decrypt(dst, src[:x.blockSize])
		src = src[x.blockSize:]
		dst = dst[x.blockSize:]
	}
}
