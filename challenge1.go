package cryptopals

import (
	"encoding/base64"
	"encoding/hex"
)

func solve1() string {
	return toBase64(fromHex("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"))
}

func fromHex(s string) []byte {
	data, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

func toHex(data []byte) string {
	return hex.EncodeToString(data)
}

func toBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}
