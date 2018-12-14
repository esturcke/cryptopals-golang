package cryptopals

/*

# Fixed XOR

Write a function that takes two equal-length buffers and produces their XOR combination.

If your function works properly, then when you feed it the string:

```
1c0111001f010100061a024b53535009181c
```

... after hex decoding, and when XOR'd against:

```
686974207468652062756c6c277320657965
```

... should produce:

```
746865206b696420646f6e277420706c6179
```

*/
func solve2() string {
	return toHex(xor(
		fromHex("1c0111001f010100061a024b53535009181c"),
		fromHex("686974207468652062756c6c277320657965"),
	))
}

func xor(a []byte, b []byte) []byte {
	if len(a) != len(b) {
		panic("xor only works on equal length byte slices")
	}
	c := make([]byte, len(a))
	for i := range a {
		c[i] = a[i] ^ b[i]
	}
	return c
}
