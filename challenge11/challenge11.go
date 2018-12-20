package challenge11

import (
	"crypto/aes"
	"math/rand"

	"github.com/esturcke/cryptopals-golang/bytes"
	"github.com/esturcke/cryptopals-golang/crypt"
)

/*Solve challenge 11

An ECB/CBC detection oracle

See https://cryptopals.com/sets/2/challenges/11

Now that you have ECB and CBC working:

Write a function to generate a random AES key; that's just 16 random bytes.

Write a function that encrypts data under an unknown key --- that is, a function that generates a random key and encrypts under it.

The function should look like:

	encryption_oracle(your-input)
	=> [MEANINGLESS JIBBER JABBER]

	Under the hood, have the function append 5-10 bytes (count chosen randomly) before the plaintext and 5-10 bytes after the plaintext.

Now, have the function choose to encrypt under ECB 1/2 the time, and under CBC the other half (just use random IVs each time for CBC). Use rand(2) to decide which to use.

Detect the block cipher mode the function is using each time. You should end up with a piece of code that, pointed at a block box that might be encrypting ECB or CBC, tells you which one is happening.

*/
func Solve() string {
	for i := 0; i < 100; i++ {
		ct, isCbc := encryptionOracle(make([]byte, 48))

		// Check if the 2nd and 3rd block match
		if match(ct[16:32], ct[32:48]) && isCbc {
			panic("Looks like ECB to me!!!")
		}
	}
	return "guessed it"
}

func match(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func encryptionOracle(pt []byte) ([]byte, bool) {
	prefix := bytes.Random(rand.Intn(6) + 5)
	suffix := bytes.Random(rand.Intn(6) + 5)
	pt = bytes.Join(prefix, pt, suffix)

	block, err := aes.NewCipher(bytes.Random(16))
	if err != nil {
		panic(err)
	}
	if rand.Intn(2) > 0 {
		// Use CBC
		iv := bytes.Random(16)
		return crypt.EncryptCbc(block, bytes.PadPkcs7(pt, block.BlockSize()), iv), true
	}

	// Use ECB
	return crypt.EncryptEcb(block, bytes.PadPkcs7(pt, block.BlockSize())), false

}
