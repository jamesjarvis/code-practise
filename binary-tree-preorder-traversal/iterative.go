type stack []*TreeNode

func (s stack) Empty() bool { return len(s) == 0 }
func (s *stack) Put(i *TreeNode)  { (*s) = append((*s), i) }
func (s *stack) Pop() *TreeNode {
d := (*s)[len(*s)-1]
(*s) = (*s)[:len(*s)-1]
return d
}


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
	
	var s stack
	s.Put(root)
	for !s.Empty() {
			temp := s.Pop()
			things = append(things, temp.Val)
			if temp.Right != nil {
					s.Put(temp.Right)
			}
			if temp.Left != nil {
					s.Put(temp.Left)
			}
	}
	return things
}
