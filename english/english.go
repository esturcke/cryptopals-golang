package english

import (
	"math"
)

var englishFreqencyVector = [256]float64{
	byte('a'): 0.0651738,
	byte('b'): 0.0124248,
	byte('c'): 0.0217339,
	byte('d'): 0.0349835,
	byte('e'): 0.1041442,
	byte('f'): 0.0197881,
	byte('g'): 0.0158610,
	byte('h'): 0.0492888,
	byte('i'): 0.0558094,
	byte('j'): 0.0009033,
	byte('k'): 0.0050529,
	byte('l'): 0.0331490,
	byte('m'): 0.0202124,
	byte('n'): 0.0564513,
	byte('o'): 0.0596302,
	byte('p'): 0.0137645,
	byte('q'): 0.0008606,
	byte('r'): 0.0497563,
	byte('s'): 0.0515760,
	byte('t'): 0.0729357,
	byte('u'): 0.0225134,
	byte('v'): 0.0082903,
	byte('w'): 0.0171272,
	byte('x'): 0.0013692,
	byte('y'): 0.0145984,
	byte('z'): 0.0007836,
	byte(' '): 0.1918182,
}

func getFrequencyVector(data []byte) [256]float64 {
	var v [256]float64
	d := float64(1) / float64(len(data))
	for _, b := range data {
		v[b] += d
	}
	return v
}

func distance(v1, v2 [256]float64) float64 {
	var sumSquares float64
	for i := 0; i < 256; i++ {
		diff := v2[i] - v1[i]
		sumSquares += diff * diff
	}
	return math.Sqrt(sumSquares)
}

// LikeEnglish returns a score for how close the byte string is to
// an English string
func LikeEnglish(data []byte) float64 {
	return 1 - distance(englishFreqencyVector, getFrequencyVector(data))
}
