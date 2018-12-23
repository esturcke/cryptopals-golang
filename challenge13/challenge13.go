package challenge13

/*Solve challenge 13

ECB cut-and-paste

See https://cryptopals.com/sets/2/challenges/13

Write a k=v parsing routine, as if for a structured cookie. The routine should take:

	foo=bar&baz=qux&zap=zazzle

... and produce:

	{
	  foo: 'bar',
	  baz: 'qux',
	  zap: 'zazzle'
	}

(you know, the object; I don't care if you convert it to JSON).

Now write a function that encodes a user profile in that format, given an email address. You should have something like:

	profile_for("foo@bar.com")

... and it should produce:

	{
	  email: 'foo@bar.com',
	  uid: 10,
	  role: 'user'
	}

... encoded as:

	email=foo@bar.com&uid=10&role=user

Your "profile_for" function should not allow encoding metacharacters (& and =). Eat them, quote them, whatever you want to do, but don't let people set their email address to "foo@bar.com&role=admin".

Now, two more easy functions. Generate a random AES key, then:

A. Encrypt the encoded user profile under the key; "provide" that to the "attacker".
B. Decrypt the encoded user profile and parse it.

Using only the user input to profile_for() (as an oracle to generate "valid" ciphertexts) and the ciphertexts themselves, make a role=admin profile.

*/
func Solve() string {
	return "xxx"
}

/*
func Solve() string {
	// 1. Get block size
	blockSize := findBlockSize()

	// 2. Detect ECB
	ct := encryptionOracle(make([]byte, blockSize*2))
	if !bytes.Match(ct[:blockSize], ct[blockSize:2*blockSize]) {
		panic("Not ECB!")
	}

	// 3-6. Decrypt

	// Find the length of the secret plus padding
	secretLength := len(encryptionOracle([]byte{}))
	secretMessage := make([]byte, secretLength)

	// The secretLength should be a multiple of the block size
	if secretLength%blockSize != 0 {
		panic("Secret is not a multiple of the block size")
	}

	// Put the first unknown byte of the secret at the
	// end of a block to get a target for the block, then loop through all
	// characters to find the match
	for i := range secretMessage {
		offsetLength := blockSize - 1 - i%blockSize // cycled from blockSize - 1 to 0
		offset := make([]byte, offsetLength)
		block := i / blockSize // block we want to discover the last byte of
		target := encryptionOracle(offset)[block*blockSize : (block+1)*blockSize]

		// Look for the first unknown byte, placed at the last place in a block
		for c := 0; c < 256; c++ {
			message := bytes.Join(offset, secretMessage[:i], []byte{byte(c)}) // join block alignment offset with what we know and then next guess
			block := (offsetLength + i) / blockSize
			guess := encryptionOracle(message)[block*blockSize : (block+1)*blockSize]
			if bytes.Match(target, guess) {
				secretMessage[i] = byte(c)
				break
			}
		}
	}

	return string(bytes.StripPkcs7(secretMessage))
}

func findBlockSize() int {
	length1 := len(encryptionOracle([]byte{}))
	length2 := length1

	// Find jump in length
	for i := 0; length2 == length1; i++ {
		length2 = len(encryptionOracle(make([]byte, i)))
	}
	return length2 - length1
}

var block, _ = aes.NewCipher(bytes.Random(16))

func encryptionOracle(pt []byte) []byte {
	suffix := bytes.FromBase64("Um9sbGluJyBpbiBteSA1LjAKV2l0aCBteSByYWctdG9wIGRvd24gc28gbXkgaGFpciBjYW4gYmxvdwpUaGUgZ2lybGllcyBvbiBzdGFuZGJ5IHdhdmluZyBqdXN0IHRvIHNheSBoaQpEaWQgeW91IHN0b3A/IE5vLCBJIGp1c3QgZHJvdmUgYnkK")
	pt = bytes.Join(pt, suffix)
	return crypt.EncryptEcb(block, bytes.PadPkcs7(pt, block.BlockSize()))
}

*/
