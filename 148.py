# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def sortList(self, head: Optional[ListNode]) -> Optional[ListNode]:
        ls_position = list()
        if head is None:
            return head
        
        while head:
            ls_position.append(head)
            head = head.next

        ls_position.sort(key=lambda x: x.val)
        head = ls_position[0]
        cur = head
        for i in range(1, len(ls_position)):
            cur.next = ls_position[i]
            cur = cur.next
        cur.next = None
        return head
