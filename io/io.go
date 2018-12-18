package io

import (
	"bufio"
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
