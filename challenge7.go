package cryptopals

import (
	"crypto/aes"

	"github.com/esturcke/cryptopals-golang/bytes"
)

/*

AES in ECB mode

See https://cryptopals.com/sets/1/challenges/7

The Base64-encoded content in this file (https://cryptopals.com/static/challenge-data/7.txt) has been encrypted via AES-128 in ECB mode under the key

	YELLOW SUBMARINE

(case-sensitive; exactly 16 characters; I like "YELLOW SUBMARINE" because it's exactly 16 bytes long, and now you do too).

Decrypt it. You know the key, after all.

Easiest way: use OpenSSL::Cipher and give it AES-128-ECB as the cipher.

*/
func solve7() string {
	ct := bytes.FromBase64(string(readFile("data/7.txt")))
	block, error := aes.NewCipher([]byte("YELLOW SUBMARINE"))
	if error != nil {
		panic(error)
	}

	pt := make([]byte, len(ct))
	for i := 0; i < len(ct); i += 16 {
		block.Decrypt(pt[i:i+16], ct[i:i+16])
	}

	return string(bytes.StripPkcs7(pt))
}
