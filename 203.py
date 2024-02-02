class Solution:
    def removeElements(self, head, val: int):
        while head and head.val == val:
            head = head.next

        cur_item = head
        while cur_item and cur_item.next:
            if cur_item.next.val == val:
                cur_item.next = cur_item.next.next
            else:
                cur_item = cur_item.next

        return head

class ListNode:
    def __init__(self, val) -> None:
        self.val = val
        self.next = None

class LinkedList:
    def __init__(self):
        self.head = None

    def Appending(self, val):
        new_val = ListNode(val)

        if self.head is None:
            self.head = new_val
            return
    
        cur_item = self.head
        while cur_item.next is not None:
            cur_item = cur_item.next
        
        cur_item.next = new_val

    def Inserting(self, prev, val):
        new_val = ListNode(val)

        if self.head.val == prev:
            new_val.next = self.head
            self.head = new_val

        cur_val = self.head
        while cur_val.next:
            post = cur_val.next
            if post.val == prev:
                cur_val.next = new_val
                new_val.next = post
                return
            cur_val = post    
        print("Incorrect prev")


head = LinkedList()
head.Appending(7)
head.Appending(7)
head.Appending(7)
head.Appending(7)
head.Appending(7)
head.Appending(7)
head.Appending(7)

print(Solution().removeElements(head.head, 7))

# cur = head.head
# while cur:
#     print(cur.val)
#     cur = cur.next


    