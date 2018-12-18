package challenge8

import (
	"github.com/esturcke/cryptopals-golang/bytes"
	"github.com/esturcke/cryptopals-golang/io"
)

/*

Detect AES in ECB mode

See https://cryptopals.com/sets/1/challenges/8

In this file (https://cryptopals.com/static/challenge-data/8.txt) are a bunch of hex-encoded ciphertexts.

One of them has been encrypted with ECB.

Detect it.

Remember that the problem with ECB is that it is stateless and deterministic; the same 16 byte plaintext block will always produce the same 16 byte ciphertext.

*/
func Solve() string {
	for _, line := range io.ReadHexLines("data/8.txt") {
		if hasRepeats(line) {
			return bytes.ToHex(line)
		}
	}
	return ""
}

func hasRepeats(line []byte) bool {
	found := make(map[string]bool)
	for i := 0; i < len(line)-16; i += 16 {
		chunk := string(line[i : i+16])
		if found[chunk] {
			return true
		}
		found[chunk] = true
	}
	return false
}
