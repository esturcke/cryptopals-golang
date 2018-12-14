package cryptopals

import (
	"encoding/base64"
	"encoding/hex"
	"log"
)

func solve1() string {
	hexString := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	data, err := hex.DecodeString(hexString)
	if err != nil {
		log.Fatal(err)
	}
	return base64.StdEncoding.EncodeToString(data)
}
