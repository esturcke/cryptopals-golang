package cryptopals

// Challenge solver and solutions
type Challenge struct {
	number   int
	solver   func() string
	solution string
}

// Challenges is a list of all solved challenges
var Challenges = []Challenge{
	{1, solve1, "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"},
}
