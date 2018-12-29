package challenge14

import (
	"crypto/aes"
	"math/rand"
	"time"

	"github.com/esturcke/cryptopals-golang/bytes"
	"github.com/esturcke/cryptopals-golang/crypt"
)

/*Solve challenge 14

Byte-at-a-time ECB decryption (Harder)

See https://cryptopals.com/sets/2/challenges/14

Take your oracle function from #12. Now generate a random count of random bytes and prepend this string to every plaintext. You are now doing:

	AES-128-ECB(random-prefix || attacker-controlled || target-bytes, random-key)

Same goal: decrypt the target-bytes.

*/
func Solve() string {
	// Get something sort of random

	// 1. Get block size
	blockSize := findBlockSize()

	// 2. Detect ECB
	ct := encryptionOracle(make([]byte, blockSize*3))
	isEcb, _ := hasRepeatedBlock(ct, blockSize)
	if !isEcb {
		panic("Not ECB!")
	}

	// 3-6. Decrypt

	// Find the length of the prefix and the buffer length to
	// add to align the start of the chosen plaintext
	prefixLength := findPrefixLength(blockSize)
	prefixBuffer := (blockSize - prefixLength%blockSize) % blockSize
	prefixBlocks := (prefixLength + prefixBuffer) / blockSize

	// Find the length of the secret plus padding
	secretLength := len(encryptionOracle([]byte{})) - prefixLength
	secretMessage := make([]byte, secretLength)

	// Put the first unknown byte of the secret at the
	// end of a block to get a target for the block, then loop through all
	// characters to find the match
	for i := range secretMessage {
		offsetLength := prefixBuffer + blockSize - 1 - i%blockSize // cycled from blockSize - 1 to 0
		offset := make([]byte, offsetLength)
		block := prefixBlocks + i/blockSize // block we want to discover the last byte of
		target := encryptionOracle(offset)[block*blockSize : (block+1)*blockSize]

		// Look for the first unknown byte, placed at the last place in a block
		for c := 0; c < 256; c++ {
			message := bytes.Join(offset, secretMessage[:i], []byte{byte(c)}) // join block alignment offset with what we know and then next guess
			block := (prefixLength + offsetLength + i) / blockSize
			guess := encryptionOracle(message)[block*blockSize : (block+1)*blockSize]

			if bytes.Match(target, guess) {
				if c > 1 {
					secretMessage[i] = byte(c)
				}
				break
			}
		}
	}

	return string(bytes.StripByte(secretMessage, 0))
}

func findBlockSize() int {
	length1 := len(encryptionOracle([]byte{}))
	length2 := length1

	// Find jump in length
	for i := 0; length2 == length1; i++ {
		length2 = len(encryptionOracle(make([]byte, i)))
	}
	return length2 - length1
}

func hasRepeatedBlock(ct []byte, blockSize int) (bool, int) {
	for i := 0; i < len(ct)-2*blockSize; i += blockSize {
		if bytes.Match(ct[i:i+blockSize], ct[i+blockSize:i+2*blockSize]) {
			return true, i
		}
	}
	return false, 0
}

func findPrefixLength(blockSize int) int {
	for i := 2 * blockSize; i < 3*blockSize; i++ {
		if hasRepeat, pos := hasRepeatedBlock(encryptionOracle(make([]byte, i)), blockSize); hasRepeat {
			return pos + 2*blockSize - i
		}
	}
	panic("Failed to find the prefix length")
}

var block, _ = aes.NewCipher(bytes.Random(16))
var prefix = randomPrefix()

func randomPrefix() []byte {
	rand.Seed(time.Now().UnixNano())
	return bytes.Random(rand.Intn(60))
}

func encryptionOracle(pt []byte) []byte {
	suffix := bytes.FromBase64("Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK")
	pt = bytes.Join(prefix, pt, suffix)
	return crypt.EncryptEcb(block, bytes.PadPkcs7(pt, block.BlockSize()))
}
