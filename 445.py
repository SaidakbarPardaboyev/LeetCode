# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def addTwoNumbers(self, l1: Optional[ListNode], l2: Optional[ListNode]) -> Optional[ListNode]:
        num1 = 0
        num2 = 0

        cur = l1
        while cur:
            num1 = num1 * 10 + cur.val
            cur = cur.next

        cur = l2
        while cur:
            num2 = num2 * 10 + cur.val
            cur = cur.next

        head = None

        res = str(num1 + num2)
        head = ListNode(int(res[0]))
        cur = head
        for i in range(1, len(res)):
            cur.next = ListNode(int(res[i]))
            cur = cur.next

        return head
    