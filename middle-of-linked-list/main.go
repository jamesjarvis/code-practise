/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func middleNode(head *ListNode) *ListNode {
	return helper(head, head, 1)
//     So you want to start at 0 with both the same, then increase both by one
//     Then increase the current by one, leaving the middle
//     Then repeat
}

func helper(midNode *ListNode, curNode *ListNode, length int) *ListNode {
	curNode = curNode.Next
	if(curNode == nil) {
			return midNode
	}
	length++
	if(length % 2 == 0) {
			midNode = midNode.Next
	}
	return helper(midNode, curNode, length)
}