package cryptopals

import (
	"github.com/esturcke/cryptopals-golang/bytes"
)

func solve1() string {
	return bytes.ToBase64(bytes.FromHex("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"))
}
