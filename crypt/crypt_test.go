package crypt

import (
	"crypto/aes"
	"fmt"
	"testing"

	"github.com/esturcke/cryptopals-golang/bytes"
)

func TestCbc(t *testing.T) {
	var tests = []string{
		"",
		"1234567890",
		"hello",
		"<-- something really long -->",
	}

	for _, test := range tests {
		block, _ := aes.NewCipher(bytes.Random(16))
		iv := bytes.Random(16)
		t.Run(fmt.Sprintf("CBC encrypt/decrypt %s", test), func(t *testing.T) {
			ct := EncryptCbc(block, bytes.PadPkcs7([]byte(test), block.BlockSize()), iv)
			pt := bytes.StripPkcs7(DecryptCbc(block, ct, iv))
			if string(pt) != test {
				t.Errorf("Expected %s, but got %s", test, pt)
			}
		})
	}
}
