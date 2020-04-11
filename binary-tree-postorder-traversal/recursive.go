/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func postorderTraversal(root *TreeNode) []int {
	var things []int
	if root == nil {
			return things
	}
	things = []int{root.Val}
	return append(postorderTraversal(root.Left),append(postorderTraversal(root.Right), things...)...)
}