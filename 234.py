class Solution:
    def isPalindrome(self, head) -> bool:
        ls = list()
        ls_rev = list()

        while head:
            ls.append(head.val)
            ls_rev.insert(0, head.val)
            head = head.next
        
        if ls == ls_rev:
            return True
        else:
            return False

class ListNode:
    def __init__(self, val):
        self.val = val
        self.next = None

class LinkedList:
    def __init__(self) -> None:
        self.head = None

    def Appending(self, val):
        new_item = ListNode(val)

        if self.head is None:
            self.head = new_item
            return
    
        cur_item = self.head
        while cur_item.next:
            cur_item = cur_item.next
    
        cur_item.next = new_item

head = LinkedList()
head.Appending(1)
head.Appending(2)
head.Appending(2)
head.Appending(1)

print(Solution().isPalindrome(head.head))