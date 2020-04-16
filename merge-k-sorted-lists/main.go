/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func mergeKLists(lists []*ListNode) *ListNode {
	newlist := &ListNode{}
	dummy := newlist
	for !allempty(lists) {
			i := smallest(lists)
			dummy.Next = lists[i]
			lists[i] = lists[i].Next
			dummy = dummy.Next
	}
	return newlist.Next
}

func smallest(lists []*ListNode) int {
	smallest := 1<<31 - 1
	smallestindex := 0
	for i, v := range lists {
			if v != nil && v.Val < smallest {
					smallest = v.Val
					smallestindex = i
			}
	}
	return smallestindex
}

func allempty(lists []*ListNode) bool {
	for _, v := range lists {
			if v != nil {
					return false
			}
	}
	return true
}