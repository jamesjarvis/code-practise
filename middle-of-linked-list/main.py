# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None
class Solution:
    def middleNode(self, head: ListNode) -> ListNode:
        midNode = head
        length = 1
        while head.next:
            length += 1
            if length % 2 == 0:
                midNode = midNode.next
            head = head.next
        return midNode
        