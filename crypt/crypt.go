package crypt

import (
	"crypto/cipher"

	"github.com/esturcke/cryptopals-golang/bytes"
)

// DecryptCbc decrypt using CBC
func DecryptCbc(block cipher.Block, ct, iv []byte) []byte {
	blockSize := block.BlockSize()
	if blockSize != len(iv) {
		panic("IV expected to match block size")
	}
	if len(ct)%blockSize != 0 {
		panic("Data expected to be a multiple of the block size")
	}

	pt := make([]byte, len(ct))
	carry := iv
	for i := 0; i < len(ct); i += blockSize {
		block.Decrypt(pt[i:i+16], ct[i:i+16])
		copy(pt[i:i+16], bytes.Xor(pt[i:i+16], carry))
		carry = ct[i : i+16]
	}
	return pt
}
