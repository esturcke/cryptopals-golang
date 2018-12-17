package crack

import (
	"github.com/esturcke/cryptopals-golang/bytes"
	"github.com/esturcke/cryptopals-golang/english"
)

// GuessXorKey finds a 1 byte xor key that produces
// a string that matches an English character distribution.
func GuessXorKey(ct []byte) (bestKey byte) {
	var topScore float64
	key := make([]byte, len(ct))
	for b := 0; b <= 255; b++ {
		for i := range ct {
			key[i] = byte(b)
		}
		guess := bytes.Xor(ct, key)
		score := english.LikeEnglish(guess)
		if score > topScore {
			bestKey = byte(b)
			topScore = score
		}
	}
	return
}
