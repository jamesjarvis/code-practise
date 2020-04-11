/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func preorderTraversal(root *TreeNode) []int {
	var things []int
	if root == nil {
			return things
	}
	things = []int{root.Val}
	return append(append(things, preorderRecur(root.Left)...),preorderRecur(root.Right)...)
}
