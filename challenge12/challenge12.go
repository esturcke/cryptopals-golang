package challenge12

import (
	"crypto/aes"

	"github.com/esturcke/cryptopals-golang/bytes"
	"github.com/esturcke/cryptopals-golang/crypt"
)

/*Solve challenge 12

Byte-at-a-time ECB decryption (Simple)

See https://cryptopals.com/sets/2/challenges/12

Copy your oracle function to a new function that encrypts buffers under ECB mode using a consistent but unknown key (for instance, assign a single random key, once, to a global variable).

Now take that same function and have it append to the plaintext, BEFORE ENCRYPTING, the following string:

  Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkg
  aGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBq
  dXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUg
  YnkK

Base64 decode the string before appending it. Do not base64 decode the string by hand; make your code do it. The point is that you don't know its contents.

What you have now is a function that produces:

	AES-128-ECB(your-string || unknown-string, random-key)

It turns out: you can decrypt "unknown-string" with repeated calls to the oracle function!

Here's roughly how:

1. Feed identical bytes of your-string to the function 1 at a time --- start with 1 byte ("A"), then "AA", then "AAA" and so on. Discover the block size of the cipher. You know it, but do this step anyway.
2. Detect that the function is using ECB. You already know, but do this step anyways.
3. Knowing the block size, craft an input block that is exactly 1 byte short (for instance, if the block size is 8 bytes, make "AAAAAAA"). Think about what the oracle function is going to put in that last byte position.
4. Make a dictionary of every possible last byte by feeding different strings to the oracle; for instance, "AAAAAAAA", "AAAAAAAB", "AAAAAAAC", remembering the first block of each invocation.
5. Match the output of the one-byte-short input to one of the entries in your dictionary. You've now discovered the first byte of unknown-string.
6. Repeat for the next byte.

*/
func Solve() string {
	// 1. Get block size
	blockSize := findBlockSize()

	// 2. Detect ECB
	ct := encryptionOracle(make([]byte, blockSize*2))
	if !bytes.Match(ct[:blockSize], ct[blockSize:2*blockSize]) {
		panic("Not ECB!")
	}

	// 3-6. Decrypt

	// Find the length of the secret plus padding
	secretLength := len(encryptionOracle([]byte{}))
	secretMessage := make([]byte, secretLength)

	// The secretLength should be a multiple of the block size
	if secretLength%blockSize != 0 {
		panic("Secret is not a multiple of the block size")
	}

	// Put the first unknown byte of the secret at the
	// end of a block to get a target for the block, then loop through all
	// characters to find the match
	for i := range secretMessage {
		offsetLength := blockSize - 1 - i%blockSize // cycled from blockSize - 1 to 0
		offset := make([]byte, offsetLength)
		block := i / blockSize // block we want to discover the last byte of
		target := encryptionOracle(offset)[block*blockSize : (block+1)*blockSize]

		// Look for the first unknown byte, placed at the last place in a block
		for c := 0; c < 256; c++ {
			message := bytes.Join(offset, secretMessage[:i], []byte{byte(c)}) // join block alignment offset with what we know and then next guess
			block := (offsetLength + i) / blockSize
			guess := encryptionOracle(message)[block*blockSize : (block+1)*blockSize]
			if bytes.Match(target, guess) {
				if c > 1 { // Once we hit a 1 we hit the padding
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

var block, _ = aes.NewCipher(bytes.Random(16))

func encryptionOracle(pt []byte) []byte {
	suffix := bytes.FromBase64("Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK")
	pt = bytes.Join(pt, suffix)
	return crypt.EncryptEcb(block, bytes.PadPkcs7(pt, block.BlockSize()))
}
