package cryptopals

import (
	"bufio"
	"os"

	"github.com/esturcke/cryptopals-golang/bytes"
	"github.com/esturcke/cryptopals-golang/english"
)

/*

# Detect single-character XOR

The hex encoded string:


One of the 60-character strings in this [file](https://cryptopals.com/static/challenge-data/4.txt)
has been encrypted by single-character XOR.

Find it.
```

*/
func solve4() string {
	var topScore float64
	var pt string
	for _, ct := range getCts() {
		score, guess := decodeByteXor(ct)
		if score > topScore {
			topScore, pt = score, guess
		}
	}
	return pt
}

func decodeByteXor(ct []byte) (float64, string) {
	var topScore float64
	var pt []byte
	key := make([]byte, len(ct))
	for b := 0; b <= 255; b++ {
		for i := range ct {
			key[i] = byte(b)
		}
		guess := bytes.Xor(ct, key)
		score := english.LikeEnglish(guess)
		if score > topScore {
			pt = guess
			topScore = score
		}
	}
	return topScore, string(pt)
}

func getCts() [][]byte {
	file, err := os.Open("data/4.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var cts [][]byte
	for scanner.Scan() {
		cts = append(cts, bytes.FromHex(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return cts
}
