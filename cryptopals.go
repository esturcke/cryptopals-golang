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
	{2, solve2, "746865206b696420646f6e277420706c6179"},
	{3, solve3, "Cooking MC's like a pound of bacon"},
	{4, solve4, "Now that the party is jumping\n"},
}
