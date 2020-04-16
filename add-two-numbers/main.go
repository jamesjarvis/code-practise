/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sumlist := &ListNode{}
	templist := sumlist
	carry := 0
	for l1 != nil || l2 != nil {
			l, r := 0, 0
			if l1 != nil {
					l = l1.Val
					l1 = l1.Next
			}
			if l2 != nil {
					r = l2.Val
					l2 = l2.Next
			}
			temp := l + r + carry
			carry = temp / 10
			templist.Next = &ListNode{
					Val: temp % 10,
					Next: nil,
			}

			templist = templist.Next
	}
	if carry > 0 {
			templist.Next = &ListNode{
					Val: carry,
					Next: nil,
			}
	}
	return sumlist.Next
}