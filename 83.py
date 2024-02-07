# Definition for singly-linked list.
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
    def deleteDuplicates(self, head):
        if head is None:
            return head
        if head.next is None:
            return head
        
        p1 = head
        p2 = head.next
        while p1:
            if p2 is None or p1.next is None:
                break
            while p1.val == p2.val and p2.next:
                p2 = p2.next
            if p2.val != p1.val:
                p1.next = p2
                p1 = p1.next
                p2 = p1.next
            else:
                p1.next = None
                break

        return head
 

hd = LinkedList()
hd.Appending(1)
hd.Appending(2)
hd.Appending(3)
hd.Appending(3)
hd.Appending(3)
hd.Appending(4)
hd.Appending(4)
hd.Appending(4)

cur = Solution().deleteDuplicates(hd.head)

# cur = hd.head
while cur:
    print(cur.val)
    cur = cur.next
