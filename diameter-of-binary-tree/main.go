import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    if root == nil {
        return 0
    }
    lheight := height(root.Left)
    rheight := height(root.Right)
    
    ldiameter := diameterOfBinaryTree(root.Left)
    rdiameter := diameterOfBinaryTree(root.Right)
    
    // Return max of following three 
    // 1) Diameter of left subtree 
    // 2) Diameter of right subtree 
    // 3) Height of left subtree + height of right subtree
    return int(math.Max(float64(lheight+rheight),math.Max(float64(ldiameter),float64(rdiameter))))
}

func height(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    return 1 + int(math.Max(float64(height(root.Left)),float64(height(root.Right))))
}