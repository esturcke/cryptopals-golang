package challenge10

import (
	"crypto/aes"

	"github.com/esturcke/cryptopals-golang/bytes"
	"github.com/esturcke/cryptopals-golang/crypt"
	"github.com/esturcke/cryptopals-golang/io"
)

/*Solve challenge 10

Implement CBC mode

See https://cryptopals.com/sets/2/challenges/10

CBC mode is a block cipher mode that allows us to encrypt irregularly-sized messages, despite the fact that a block cipher natively only transforms individual blocks.

In CBC mode, each ciphertext block is added to the next plaintext block before the next call to the cipher core.

The first plaintext block, which has no associated previous ciphertext block, is added to a "fake 0th ciphertext block" called the initialization vector, or IV.

Implement CBC mode by hand by taking the ECB function you wrote earlier, making it encrypt instead of decrypt (verify this by decrypting whatever you encrypt to test), and using your XOR function from the previous exercise to combine them.

The file here (https://cryptopals.com/static/challenge-data/10.txt) is intelligible (somewhat) when CBC decrypted against "YELLOW SUBMARINE" with an IV of all ASCII 0 (\x00\x00\x00 &c)

*/
func Solve() string {
	block, err := aes.NewCipher([]byte("YELLOW SUBMARINE"))
	if err != nil {
		panic(err)
	}
	ct := io.ReadBase64("data/10.txt")
	iv := make([]byte, 16)
	return string(bytes.StripPkcs7(crypt.DecryptCbc(block, ct, iv)))
}
