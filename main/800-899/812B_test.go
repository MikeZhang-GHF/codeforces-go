package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestCF812B(t *testing.T) {
	// just copy from website
	rawText := `
2 2
0010
0100
outputCopy
5
inputCopy
3 4
001000
000010
000010
outputCopy
12
inputCopy
4 3
01110
01110
01110
01110
outputCopy
18
inputCopy
5 93
00000000000000000000000000000000000000000000000000000000100000000000000000000000000000000001010
00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000
00000010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000
00000000000000000000000000000010000000000000000000100000000000000000000000000000000000000000000
00000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000000
outputCopy
265
inputCopy
8 8
0011101110
0110010100
0100111110
0111111100
0011010100
0001101110
0111100000
0110111000
outputCopy
77`
	testutil.AssertEqualCase(t, rawText, 0, CF812B)
}