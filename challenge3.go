package cryptopals

import (
	"github.com/esturcke/cryptopals-golang/english"
)

/*

# Single-byte XOR cipher

The hex encoded string:

```
1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736
```

... has been XOR'd against a single character. Find the key, decrypt the message.

You can do this by hand. But don't: write code to do it for you.

How? Devise some method for "scoring" a piece of English plaintext. Character frequency is a good metric. Evaluate each output and choose the one with the best score.
```

*/
func solve3() string {
	ct := fromHex("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	var topScore float64
	var pt []byte
	key := make([]byte, len(ct))
	for b := 0; b <= 255; b++ {
		for i := range ct {
			key[i] = byte(b)
		}
		guess := xor(ct, key)
		score := english.LikeEnglish(guess)
		if score > topScore {
			pt = guess
			topScore = score
		}
	}
	return string(pt)
}
