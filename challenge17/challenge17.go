package challenge17

import (
	"crypto/aes"
	"fmt"
	"math/rand"
	"time"

	"github.com/esturcke/cryptopals-golang/bytes"
	"github.com/esturcke/cryptopals-golang/crypt"
)

/*Solve challenge 17

The CBC padding oracle

See https://cryptopals.com/sets/3/challenges/17

This is the best-known attack on modern block-cipher cryptography.

Combine your padding code and your CBC code to write two functions.

The first function should select at random one of the following 10 strings:

	MDAwMDAwTm93IHRoYXQgdGhlIHBhcnR5IGlzIGp1bXBpbmc=
	MDAwMDAxV2l0aCB0aGUgYmFzcyBraWNrZWQgaW4gYW5kIHRoZSBWZWdhJ3MgYXJlIHB1bXBpbic=
	MDAwMDAyUXVpY2sgdG8gdGhlIHBvaW50LCB0byB0aGUgcG9pbnQsIG5vIGZha2luZw==
	MDAwMDAzQ29va2luZyBNQydzIGxpa2UgYSBwb3VuZCBvZiBiYWNvbg==
	MDAwMDA0QnVybmluZyAnZW0sIGlmIHlvdSBhaW4ndCBxdWljayBhbmQgbmltYmxl
	MDAwMDA1SSBnbyBjcmF6eSB3aGVuIEkgaGVhciBhIGN5bWJhbA==
	MDAwMDA2QW5kIGEgaGlnaCBoYXQgd2l0aCBhIHNvdXBlZCB1cCB0ZW1wbw==
	MDAwMDA3SSdtIG9uIGEgcm9sbCwgaXQncyB0aW1lIHRvIGdvIHNvbG8=
	MDAwMDA4b2xsaW4nIGluIG15IGZpdmUgcG9pbnQgb2g=
	MDAwMDA5aXRoIG15IHJhZy10b3AgZG93biBzbyBteSBoYWlyIGNhbiBibG93

.. generate a random AES key (which it should save for all future encryptions), pad the string out to the 16-byte AES block size and CBC-encrypt it under that key, providing the caller the ciphertext and IV.

The second function should consume the ciphertext produced by the first function, decrypt it, check its padding, and return true or false depending on whether the padding is valid.

What you're doing here:
This pair of functions approximates AES-CBC encryption as its deployed serverside in web applications; the second function models the server's consumption of an encrypted session token, as if it was a cookie.

It turns out that it's possible to decrypt the ciphertexts provided by the first function.

The decryption here depends on a side-channel leak by the decryption function. The leak is the error message that the padding is valid or not.

You can find 100 web pages on how this attack works, so I won't re-explain it. What I'll say is this:

The fundamental insight behind this attack is that the byte 01h is valid padding, and occur in 1/256 trials of "randomized" plaintexts produced by decrypting a tampered ciphertext.

02h in isolation is not valid padding.

02h 02h is valid padding, but is much less likely to occur randomly than 01h.

03h 03h 03h is even less likely.

So you can assume that if you corrupt a decryption AND it had valid padding, you know what that padding byte is.

It is easy to get tripped up on the fact that CBC plaintexts are "padded". Padding oracles have nothing to do with the actual padding on a CBC plaintext. It's an attack that targets a specific bit of code that handles decryption. You can mount a padding oracle on any CBC block, whether it's padded or not.

*/
func Solve() string {
	rand.Seed(time.Now().UTC().UnixNano())
	iv, ct, pt := randEncryptedString()
	if !bytes.Match(decrypt(iv, ct), pt) {
		panic(fmt.Sprintf("poo\n%v\n%v", decrypt(iv, ct), pt))
	}
	return "yay"
}

func decrypt(iv, ct []byte) []byte {
	bs := block.BlockSize()
	pt := make([]byte, len(ct))
	for i := len(ct); i >= bs; i -= bs {
		// Sometimes we have multiple match, first pick the first match
		ptBlock := decryptLastBlock(bytes.Join(iv, ct[:i]), true)

		// If the second to last byte is 2, we probably are decrypting to 16, 15...2, x
		// In this case, use the last match
		if ptBlock[bs-2] == 2 {
			ptBlock = decryptLastBlock(bytes.Join(iv, ct[:i]), false)
		}
		copy(pt[i-bs:i], ptBlock)
	}
	return pt
}

func decryptLastBlock(originalIvAndCt []byte, pickFirst bool) []byte {
	bs := block.BlockSize()
	pt := make([]byte, bs)
	ctLen := len(originalIvAndCt)
	ivAndCt := make([]byte, ctLen)
	copy(ivAndCt, originalIvAndCt)

	// We manipulate the second to last block (which might be the IV)
	tamperedCtBlock := ivAndCt[ctLen-2*bs : ctLen-bs]
	originalCtBlock := originalIvAndCt[ctLen-2*bs : ctLen-bs]

	for i := bs - 1; i >= 0; i-- {
		// Adjust CT so the resulting PT will be the new padding
		pad := byte(bs - i)
		for j := i + 1; j < bs; j++ {
			tamperedCtBlock[j] = pt[j] ^ originalCtBlock[j] ^ pad
		}

		// Remember the original
		originalByte := originalCtBlock[i]
		pt[i] = pad
		for c := 0; c < 256; c++ {
			if byte(c) != originalByte {
				tamperedCtBlock[i] = byte(c)
				if isValidPadding(ivAndCt[:bs], ivAndCt[bs:]) {
					pt[i] = originalByte ^ byte(c) ^ pad
					if pickFirst {
						break
					}
				}
			}
		}
	}
	return pt
}

var block, _ = aes.NewCipher(bytes.Random(16))

func isValidPadding(iv, ct []byte) (isValid bool) {
	isValid = true
	defer func() {
		if r := recover(); r != nil {
			isValid = false
		}
	}()
	bytes.StripPkcs7(crypt.DecryptCbc(block, ct, iv))
	return
}

func randEncryptedString() (iv, ct, pt []byte) {
	pt = bytes.PadPkcs7([]byte(randString()), block.BlockSize())
	iv = bytes.Random(16)
	ct = crypt.EncryptCbc(block, pt, iv)
	return
}

func randString() string {
	strings := []string{
		"MDAwMDAwTm93IHRoYXQgdGhlIHBhcnR5IGlzIGp1bXBpbmc=",
		"MDAwMDAxV2l0aCB0aGUgYmFzcyBraWNrZWQgaW4gYW5kIHRoZSBWZWdhJ3MgYXJlIHB1bXBpbic=",
		"MDAwMDAyUXVpY2sgdG8gdGhlIHBvaW50LCB0byB0aGUgcG9pbnQsIG5vIGZha2luZw==",
		"MDAwMDAzQ29va2luZyBNQydzIGxpa2UgYSBwb3VuZCBvZiBiYWNvbg==",
		"MDAwMDA0QnVybmluZyAnZW0sIGlmIHlvdSBhaW4ndCBxdWljayBhbmQgbmltYmxl",
		"MDAwMDA1SSBnbyBjcmF6eSB3aGVuIEkgaGVhciBhIGN5bWJhbA==",
		"MDAwMDA2QW5kIGEgaGlnaCBoYXQgd2l0aCBhIHNvdXBlZCB1cCB0ZW1wbw==",
		"MDAwMDA3SSdtIG9uIGEgcm9sbCwgaXQncyB0aW1lIHRvIGdvIHNvbG8=",
		"MDAwMDA4b2xsaW4nIGluIG15IGZpdmUgcG9pbnQgb2g=",
		"MDAwMDA5aXRoIG15IHJhZy10b3AgZG93biBzbyBteSBoYWlyIGNhbiBibG93",
	}
	return strings[rand.Intn(len(strings))]
}
