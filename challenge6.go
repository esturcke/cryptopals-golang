package cryptopals

import (
	"io/ioutil"

	"github.com/esturcke/cryptopals-golang/bytes"
	"github.com/esturcke/cryptopals-golang/crack"
	"github.com/esturcke/cryptopals-golang/io"
)

/*

# Break repeating-key XOR

*[Set 1 / Challenge 6](http://cryptopals.com/sets/1/challenges/6/)*

There's a file here (http://cryptopals.com/static/challenge-data/6.txt).
It's been base64'd after being encrypted with repeating-key XOR.

Decrypt it.

*/
func solve6() string {
	ct := bytes.FromBase64(string(io.Read("data/6.txt")))
	keyLength := guessKeyLength(ct)

	// construct the key
	key := make([]byte, keyLength)
	for i, row := range bytes.CycledSplit(ct, keyLength) {
		key[i] = crack.GuessXorKey(row)
	}
	return string(bytes.CycledXor(ct, key))
}

func guessKeyLength(ct []byte) (length int) {
	var min = 100.
	for l := 19; l <= 40; l++ {
		if d := sampledEditDistance(ct, l); d < min {
			min, length = d, l
		}
	}
	return length
}

func sampledEditDistance(ct []byte, l int) float64 {
	sum := 0
	for d := 0; d < 40; d++ {
		sum += bytes.EditDistance(ct[d:l+d], ct[l+d:2*l+d])
	}
	return float64(sum) / float64(41*l)
}

func readFile(file string) []byte {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return bytes
}
