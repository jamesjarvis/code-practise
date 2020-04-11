/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
			return head
	}
	
//     head of list after two nodes
	remaining := head.Next.Next
	
//     Change head
	newhead := head.Next
	
//     Change next of second node
	head.Next.Next = head
	
//     recursive for remaining list and change next of nodes
	head.Next = swapPairs(remaining)
	
	return newhead
}