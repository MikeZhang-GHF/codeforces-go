// Code generated by generator_test.
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [d]")
	exampleIns := [][]string{{`[6,5,4,3,2,1]`, `2`}, {`[9,9,9]`, `4`}, {`[1,1,1]`, `3`}, {`[7,1,7,1,7,1]`, `3`}, {`[11,111,22,222,33,333,44,444]`, `6`}}
	exampleOuts := [][]string{{`7`}, {`-1`}, {`3`}, {`15`}, {`843`}}
	// custom test cases or WA cases.
	//exampleIns = append(exampleIns, []string{``})
	//exampleOuts = append(exampleOuts, []string{``})
	if err := testutil.RunLeetCodeFuncWithCase(t, minDifficulty, exampleIns, exampleOuts, 0); err != nil {
		t.Fatal(err)
	}
}