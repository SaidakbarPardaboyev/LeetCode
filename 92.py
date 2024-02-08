# Definition for singly-linked list.
import os
os.system("cls")
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class LinkedList:
    def __init__(self):
        self.head = None

    def Appending(self, val):
        new_val = ListNode(val)

        if self.head is None:
            self.head = new_val
            return
        cur = self.head
        while cur.next:
            cur = cur.next
        
        cur.next = new_val

class Solution:
    def reverseBetween(self, head, left: int, right: int):
        if head.next is None:
            return head
        reversed_part = None
        cur = head
        ls_pos = list()
        for i in range(left-1):
            cur = cur.next
        
        for i in range(right-left+1):
            ls_pos.append(cur)
            cur = cur.next
        post = cur

        ls_pos.reverse()
        reversed_part = ls_pos[0]
        i = 1
        cur = reversed_part
        while i < len(ls_pos):
            cur.next = ls_pos[i]
            cur = cur.next
            i += 1
        cur.next = post

        # cur = head
        # for i in range(left-2):
        #     cur = cur.next
        # cur.next = reversed_part

        return head

hd = LinkedList()
hd.Appending(3)
hd.Appending(5)
# hd.Appending(3)
# hd.Appending(4)
# hd.Appending(5)
# hd.Appending(6)
# hd.Appending(7)


cur = Solution().reverseBetween(hd.head, 1, 2)

cur = hd.head
while cur:
    print(cur.val)
    cur = cur.next
 