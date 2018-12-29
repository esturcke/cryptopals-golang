package challenge15

import "github.com/esturcke/cryptopals-golang/bytes"

/*Solve challenge 15

PKCS#7 padding validation

See https://cryptopals.com/sets/2/challenges/15

Write a function that takes a plaintext, determines if it has valid PKCS#7 padding, and strips the padding off.

The string:

	"ICE ICE BABY\x04\x04\x04\x04"

... has valid padding, and produces the result "ICE ICE BABY".

The string:

	"ICE ICE BABY\x05\x05\x05\x05"

... does not have valid padding, nor does:

	"ICE ICE BABY\x01\x02\x03\x04"

If you are writing in a language with exceptions, like Python or Ruby, make your function throw an exception on bad padding.

Crypto nerds know where we're going with this. Bear with us.
*/
func Solve() string {
	if string(bytes.StripPkcs7([]byte("ICE ICE BABY\x04\x04\x04\x04"))) != "ICE ICE BABY" {
		panic("Fail")
	}
	return "done"
}
