package challenge13

import (
	"crypto/aes"
	"fmt"
	"strings"

	"github.com/esturcke/cryptopals-golang/bytes"
	"github.com/esturcke/cryptopals-golang/crypt"
)

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
	// Get the ct for a user profile, with "user" in a block by itself
	ctUser := encryptedProfile("a@example.com")

	// Chop the "user" off
	ctUser = ctUser[:len(ctUser)-16]

	// Figure out what the ct for "admin" should be
	ctRole := encryptedProfile("<-buffer->" + "admin" + string([]byte{4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4}))

	// Keep only the admin block (2nd block)
	ctRole = ctRole[16:32]

	// Join the user and role
	ct := append(ctUser, ctRole...)

	// Return the role
	return decode(string(decrypt(ct)))["role"]
}

func decode(s string) map[string]string {
	println(s)
	object := make(map[string]string)
	for _, substring := range strings.Split(s, "&") {
		keyValue := strings.Split(substring, "=")
		if len(keyValue) != 2 {
			panic("Invalid string")
		}
		object[keyValue[0]] = keyValue[1]
	}
	return object
}

func profileFor(email string) string {
	return fmt.Sprintf("email=%s&uid=10&role=user", sanitizeEmail(email))
}

func sanitizeEmail(email string) string {
	return strings.Replace(strings.Replace(email, "&", "", -1), "=", "", -1)
}

var block, _ = aes.NewCipher(bytes.Random(16))

func encrypt(pt []byte) []byte {
	return crypt.EncryptEcb(block, bytes.PadPkcs7(pt, block.BlockSize()))
}

func decrypt(ct []byte) []byte {
	return bytes.StripPkcs7(crypt.DecryptEcb(block, ct))
}

func encryptedProfile(email string) []byte {
	return encrypt([]byte(profileFor(email)))
}
