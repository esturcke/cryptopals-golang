package io

import (
	"bufio"
	"io/ioutil"
	"os"

	"github.com/esturcke/cryptopals-golang/bytes"
)

// ReadHexLines reads lines from a file
func ReadHexLines(path string) (lines [][]byte) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, bytes.FromHex(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return
}

// ReadBase64 read a base 64 encoded file into a byte slice
func ReadBase64(path string) []byte {
	return bytes.FromBase64(string(Read(path)))
}

// Read read a file
func Read(path string) []byte {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return bytes
}
