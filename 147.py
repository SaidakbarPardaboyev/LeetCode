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
    def insertionSortList(self, head):
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

hd = LinkedList()
hd.Appending(4)
hd.Appending(2)
hd.Appending(1)
hd.Appending(3)
# hd.Appending(5)
# hd.Appending(6)
# hd.Appending(7)


cur = Solution().insertionSortList(hd.head)

# cur = hd.head
while cur:
    print(cur.val)
    cur = cur.next
 