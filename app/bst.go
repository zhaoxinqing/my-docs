package app

import "math"

// tree 结构
type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

// 二叉搜索树有效性验证
func IsValidBST(root *TreeNode) bool {
	return isValid(root, math.MinInt64, math.MaxInt64)
}

//
func isValid(root *TreeNode, min int, max int) bool {
	if root == nil {
		return true
	}
	if root.Value <= min {
		return false
	}
	if root.Value >= max {
		return false
	}
	return isValid(root.Left, min, root.Value) && isValid(root.Right, root.Value, max)
}

// 满二叉树验证
// 完全二叉树验证
