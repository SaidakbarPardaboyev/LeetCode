# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0):
        self.val = val
        self.next = None

class linkedList:
    def __init__(self):
        self.head = None

    def Appending(self, val):
        new_value = ListNode(val)

        if self.head is None:
            self.head = new_value
            return

        current = self.head
        while current.next:
            current = current.next
        current.next = new_value

    def printList(self):
        if self.head is None:
            print("None Element")
            return
        
        current = self.head
        while current:
            print(current.val)
            current = current.next

class Solution:
    def swapPairs(self, head):
        prev_element = None
        first_element = None
        second_element = None
        post_element = None
        
        if head is None:
            return head

        if head.next:
            tem = head
            head = head.next
            tem2 = head.next
            head.next = tem
            head.next.next = tem2
            # save position of four element
            prev_element = head.next
        else:
            return head     

        current = head.next.next
        sch = 0
        while current:
            if sch == 0:
                first_element = current
                sch += 1
                current = current.next
            elif sch == 1:
                second_element = current
                post_element = current.next
                sch = 0
                prev_element.next = second_element
                second_element.next = first_element
                first_element.next = post_element

                prev_element = first_element
                current = post_element
            
        return head


ls1 = linkedList()
ls1.Appending(1)
ls1.Appending(2)
ls1.Appending(3)
ls1.Appending(4)
ls1.Appending(5)
# ls1.Appending(6)

# ls1.printList()

# ls = [ls1.head, ls2.head, ls3.head]
ls = Solution().swapPairs(ls1.head)

while ls:
    print(ls.val)
    ls = ls.next

# # print()
# ls2 = linkedList()
# ls2.Appending(1)
# ls2.Appending(3)
# ls2.Appending(4)
# # ls2.printList()

# # print()
# ls3 = linkedList()
# ls3.Appending(2)
# ls3.Appending(6)
# # ls3.printList()