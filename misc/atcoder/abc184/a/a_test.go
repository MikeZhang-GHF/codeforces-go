// Code generated by copypasta/template/atcoder/generator_test.go
package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

// 提交地址：https://atcoder.jp/contests/abc184/submit?taskScreenName=abc184_a
func Test_run(t *testing.T) {
	t.Log("Current test is [a]")
	testCases := [][2]string{
		{
			`1 2
3 4`,
			`-2`,
		},
		{
			`0 -1
1 0`,
			`1`,
		},
		{
			`100 100
100 100`,
			`0`,
		},
		// TODO 测试参数的下界和上界
		
	}
	testutil.AssertEqualStringCase(t, testCases, 0, run)
}
// https://atcoder.jp/contests/abc184/tasks/abc184_a