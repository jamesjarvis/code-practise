/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func inorderTraversal(root *TreeNode) []int {
	var things []int
	if root == nil {
			return things
	}
	things = []int{root.Val}
	return append(append(inorderTraversal(root.Left), things...),inorderTraversal(root.Right)...)
}
