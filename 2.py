# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def addTwoNumbers(self, l1, l2):
        butun = 0
        ls1 = list()
        l11 = l1
        while l11:
            ls1.append(l11.val)
            l11 = l11.next

        ls2 = list()
        l22 = l2
        while l22:
            ls2.append(l22.val)
            l22 = l22.next

        size_min = None
        kichik = None
        katta = None

        if len(ls1) > len(ls2):
            size_min = len(ls2)
            size_max = len(ls1)
            kichik = ls2
            katta = ls1
        else:
            size_min = len(ls1)
            size_max = len(ls2)
            kichik = ls1
            katta = ls2
        
        res = list()
        butun = 0
        for i in range(size_min):
            tem = katta[i] + kichik[i] + butun
            res.append(tem % 10)
            butun = tem // 10

        if size_min != size_max:
            for i in range(size_min, size_max):
                tem = katta[i] + butun
                res.append(tem%10)
                butun = tem // 10
        
        if butun > 0:
            res.append(butun)

        head = ListNode(res[0])

        run = head
        for i in range(1, len(res)):
            run.next = ListNode(res[i])
            run = run.next

        return head






# dfghyjukjhgfds