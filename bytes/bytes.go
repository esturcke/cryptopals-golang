package bytes

import (
	"encoding/base64"
	"encoding/hex"
)

// BitCount returns the count of bits in the byte slice
func BitCount(bytes []byte) (count int) {
	for _, b := range bytes {
		count += bitCounts[b]
	}
	return
}

// EditDistance returns the Hamming distance between the byte slides
func EditDistance(a, b []byte) (distance int) {
	return BitCount(Xor(a, b))
}

// Xor returns a byte slice that is the Xor of the two byte slices
func Xor(a, b []byte) []byte {
	if len(a) != len(b) {
		panic("Xor only works on equal length byte slices")
	}
	c := make([]byte, len(a))
	for i := range a {
		c[i] = a[i] ^ b[i]
	}
	return c
}

// CycledXor returns the Xor where the second byte slice is cycled
func CycledXor(a, b []byte) []byte {
	var c = make([]byte, len(a))
	for i := range a {
		c[i] = a[i] ^ b[i%len(b)]
	}
	return c
}

// FromBase64 decodes a base 64 encoded string to bytes
func FromBase64(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

// ToBase64 encodes bytes as base 64 string
func ToBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// FromHex decodes a hex string to bytes
func FromHex(s string) []byte {
	data, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

// ToHex encodes bytes to a hex string
func ToHex(data []byte) string {
	return hex.EncodeToString(data)
}

func toBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// CycledSplit splits the bytes into
func CycledSplit(data []byte, n int) [][]byte {
	rows := make([][]byte, n)
	for i := range rows {
		rows[i] = make([]byte, 0, len(data)/n+1)
	}
	for i, d := range data {
		rows[i%n] = append(rows[i%n], d)
	}
	return rows
}