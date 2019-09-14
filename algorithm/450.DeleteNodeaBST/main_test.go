package deleteNode

import (
	"testing"

	"github.com/pengliheng/leetcode/Helper"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  *TreeNode
	N2  int
	ans *TreeNode
}{
	{
		Helper.Ints2TreeNode([]int{5, 3, 6, 2, 4, -1 << 63, 7}),
		3,
		Helper.Ints2TreeNode([]int{5, 4, 6, 2, -1 << 63, -1 << 63, 7}),
	},
	{
		Helper.Ints2TreeNode([]int{0}),
		0,
		Helper.Ints2TreeNode([]int{}),
	},
	{
		Helper.Ints2TreeNode([]int{}),
		0,
		Helper.Ints2TreeNode([]int{}),
	},
	{
		Helper.Ints2TreeNode([]int{1, -1 << 63, 2}),
		2,
		Helper.Ints2TreeNode([]int{1}),
	},
	{
		Helper.Ints2TreeNode([]int{2, 1}),
		2,
		Helper.Ints2TreeNode([]int{1}),
	},
	{
		Helper.Ints2TreeNode([]int{1, -1 << 63, 2}),
		2,
		Helper.Ints2TreeNode([]int{1}),
	},
	{
		Helper.Ints2TreeNode([]int{1, -1 << 63, 2}),
		1,
		Helper.Ints2TreeNode([]int{2}),
	},
	{
		Helper.Ints2TreeNode([]int{3, 1, 4, -1 << 63, 2}),
		3,
		Helper.Ints2TreeNode([]int{4, 1, -1 << 63, -1 << 63, 2}),
	},
	// {
	// 	Helper.Ints2TreeNode([]int{2, 0, 33, -1 << 63, 1, 25, 40, -1 << 63, -1 << 63, 11, 31, 34, 45, 10, 18, 29, 32, -1 << 63, 36, 43, 46, 4, -1 << 63, 12, 24, 26, 30, -1 << 63, -1 << 63, 35, 39, 42, 44, -1 << 63, 48, 3, 9, -1 << 63, 14, 22, -1 << 63, -1 << 63, 27, -1 << 63, -1 << 63, -1 << 63, -1 << 63, 38, -1 << 63, 41, -1 << 63, -1 << 63, -1 << 63, 47, 49, -1 << 63, -1 << 63, 5, -1 << 63, 13, 15, 21, 23, -1 << 63, 28, 37, -1 << 63, -1 << 63, -1 << 63, -1 << 63, -1 << 63, -1 << 63, -1 << 63, -1 << 63, 8, -1 << 63, -1 << 63, -1 << 63, 17, 19, -1 << 63, -1 << 63, -1 << 63, -1 << 63, -1 << 63, -1 << 63, -1 << 63, 7, -1 << 63, 16, -1 << 63, -1 << 63, 20, 6}),
	// 	33,
	// 	Helper.Ints2TreeNode([]int{2, 0, 34, -1 << 63, 1, 25, 40, -1 << 63, -1 << 63, 11, 31, 35, 45, 10, 18, 29, 32, -1 << 63, 36, 43, 46, 4, -1 << 63, 12, 24, 26, 30, -1 << 63, -1 << 63, -1 << 63, 39, 42, 44, -1 << 63, 48, 3, 9, -1 << 63, 14, 22, -1 << 63, -1 << 63, 27, -1 << 63, -1 << 63, 38, -1 << 63, 41, -1 << 63, -1 << 63, -1 << 63, 47, 49, -1 << 63, -1 << 63, 5, -1 << 63, 13, 15, 21, 23, -1 << 63, 28, 37, -1 << 63, -1 << 63, -1 << 63, -1 << 63, -1 << 63, -1 << 63, -1 << 63, -1 << 63, 8, -1 << 63, -1 << 63, -1 << 63, 17, 19, -1 << 63, -1 << 63, -1 << 63, -1 << 63, -1 << 63, -1 << 63, -1 << 63, 7, -1 << 63, 16, -1 << 63, -1 << 63, 20, 6}),
	// },
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, deleteNode(tc.N1, tc.N2), "输入:%v", tc)
	}
}
