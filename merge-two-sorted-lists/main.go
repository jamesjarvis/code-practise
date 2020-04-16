/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
			return l2
	} else if l2 == nil {
			return l1
	}
	returning := &ListNode{}
	dummy := returning
	for l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
					dummy.Next = &ListNode{
							Val: l1.Val,
							Next: nil,
					}
					l1 = l1.Next
			} else {
					dummy.Next = &ListNode{
							Val: l2.Val,
							Next: nil,
					}
					l2 = l2.Next
			}
			dummy = dummy.Next
	}
//     remaining possible lists
	if l1 == nil {
			dummy.Next = l2
	} else if l2 == nil {
			dummy.Next = l1
	}
	return returning.Next
}