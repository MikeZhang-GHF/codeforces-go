// Code generated by copypasta/template/nowcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [a]")
	examples := [][]string{
		{
			`2`,`1`,`1`,
			`1`,
		},
		{
			`5`,`1`,`2`,
			`11`,
		},
		// TODO 测试参数的下界和上界
		
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, nthElement, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://ac.nowcoder.com/acm/contest/6741/a