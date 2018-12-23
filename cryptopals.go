package cryptopals

import (
	"github.com/esturcke/cryptopals-golang/challenge10"
	"github.com/esturcke/cryptopals-golang/challenge11"
	"github.com/esturcke/cryptopals-golang/challenge12"
	"github.com/esturcke/cryptopals-golang/challenge13"
	"github.com/esturcke/cryptopals-golang/challenge8"
	"github.com/esturcke/cryptopals-golang/challenge9"
	"github.com/esturcke/cryptopals-golang/io"
)

var vanilla = string(io.Read(("data/play-that-funky-music.txt")))
var moreVanilla = string(io.Read("data/ice-ice-baby.txt"))

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
	{7, solve7, vanilla},
	{8, challenge8.Solve, "d880619740a8a19b7840a8a31c810a3d08649af70dc06f4fd5d2d69c744cd283e2dd052f6b641dbf9d11b0348542bb5708649af70dc06f4fd5d2d69c744cd2839475c9dfdbc1d46597949d9c7e82bf5a08649af70dc06f4fd5d2d69c744cd28397a93eab8d6aecd566489154789a6b0308649af70dc06f4fd5d2d69c744cd283d403180c98c8f6db1f2a3f9c4040deb0ab51b29933f2c123c58386b06fba186a"},
	{9, challenge9.Solve, "YELLOW SUBMARINE" + string([]byte{4, 4, 4, 4})},
	{10, challenge10.Solve, vanilla},
	{11, challenge11.Solve, "guessed it"},
	{12, challenge12.Solve, moreVanilla},
	{13, challenge13.Solve, ""},
}
