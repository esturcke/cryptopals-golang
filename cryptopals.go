package cryptopals

var vanilla = string(readFile("data/play-that-funky-music.txt"))

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
	{5, solve5, "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"},
	{6, solve6, vanilla},
}
