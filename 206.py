# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def reverseList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        ls = list()

        if head is None:
            return head

        cur = head
        while cur:
            ls.append(cur)
            cur = cur.next

        head = ls[-1]

        cur = head
        for i in range(len(ls)-2, -1, -1):
            cur.next = ls[i]
            cur = cur.next
        
        cur.next = None

        return head