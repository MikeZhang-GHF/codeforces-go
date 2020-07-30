// Code generated by copypasta/template/nowcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/leetcode/testutil"
	"testing"
)

func Test(t *testing.T) {
	t.Log("Current test is [c]")
	examples := [][]string{
		{
			`2`,`2`,`[(1,1),(1,2)]`,
			`2`,
		},
		// TODO 测试参数的下界和上界
		{
			`4`,`100`,`[(1,2),(2,1),(2,2),(2,3),(3,2)]`,
			`5`,
		},
		{
			`10`,`100`,`[(10,10),(1,1),(1,10),(10,1)]`,
			`3`,
		},
	}
	targetCaseNum := 0
	if err := testutil.RunLeetCodeFuncWithExamples(t, BoomKill, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
// https://ac.nowcoder.com/acm/contest/6741/c
