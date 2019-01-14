package crypt

import (
	"crypto/cipher"
	"encoding/binary"

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

// EncryptCbc encrypt using CBC
func EncryptCbc(block cipher.Block, pt, iv []byte) []byte {
	blockSize := block.BlockSize()
	if blockSize != len(iv) {
		panic("IV expected to match block size")
	}
	if len(pt)%blockSize != 0 {
		panic("Data expected to be a multiple of the block size")
	}

	ct := make([]byte, len(pt))
	carry := iv
	for i := 0; i < len(ct); i += blockSize {
		block.Encrypt(ct[i:i+16], bytes.Xor(pt[i:i+16], carry))
		carry = ct[i : i+16]
	}
	return ct
}

// DecryptEcb decrypt using ECB
func DecryptEcb(block cipher.Block, ct []byte) []byte {
	blockSize := block.BlockSize()
	if len(ct)%blockSize != 0 {
		panic("Data expected to be a multiple of the block size")
	}

	pt := make([]byte, len(ct))
	for i := 0; i < len(ct); i += 16 {
		block.Decrypt(pt[i:i+16], ct[i:i+16])
	}
	return pt
}

// EncryptEcb encrypt using ECB
func EncryptEcb(block cipher.Block, pt []byte) []byte {
	blockSize := block.BlockSize()
	if len(pt)%blockSize != 0 {
		panic("Data expected to be a multiple of the block size")
	}

	ct := make([]byte, len(pt))
	for i := 0; i < len(pt); i += 16 {
		block.Encrypt(ct[i:i+16], pt[i:i+16])
	}
	return ct
}

// EncryptCtr encrypt using CTR
func EncryptCtr(block cipher.Block, pt, nonce []byte) []byte {
	blockSize := block.BlockSize()
	if blockSize != len(nonce)+8 {
		panic("Nonce expected leave room for 8 byte counter")
	}

	ct := make([]byte, len(pt))
	for i := 0; i < len(ct); i += blockSize {
		key := make([]byte, blockSize)
		counter := bytes.Join(nonce, getCrtCount(i/blockSize))
		block.Encrypt(key, counter)
		n := min(blockSize, len(pt)-i)
		copy(ct[i:i+n], bytes.Xor(key[:n], pt[i:i+n]))
	}
	return ct
}

// DecryptCtr decrypt using CTR
func DecryptCtr(block cipher.Block, ct, nonce []byte) []byte {
	blockSize := block.BlockSize()
	if blockSize != len(nonce)+8 {
		panic("Nonce expected leave room for 8 byte counter")
	}

	pt := make([]byte, len(ct))
	for i := 0; i < len(pt); i += blockSize {
		key := make([]byte, blockSize)
		counter := bytes.Join(nonce, getCrtCount(i/blockSize))
		block.Encrypt(key, counter)
		n := min(blockSize, len(ct)-i)
		copy(pt[i:i+n], bytes.Xor(key[:n], ct[i:i+n]))
	}
	return pt
}

func getCrtCount(block int) []byte {
	count := make([]byte, 8)
	binary.LittleEndian.PutUint64(count, uint64(block))
	return count
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
