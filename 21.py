class Solution:
    def mergeTwoLists(self, list1: Optional[ListNode], list2: Optional[ListNode]) -> Optional[ListNode]:
        head = None
        tem = []

        while list1 or list2:

            if list1:
                tem.append(list1.val)
                list1 = list1.next

            if list2:
                tem.append(list2.val)
                list2 = list2.next

        tem.sort()

        for i in tem:
            new_item = ListNode(i)
            if head is None:
                head = new_item
            else:
                res_cur_head = head
                while res_cur_head.next:
                    res_cur_head = res_cur_head.next
                res_cur_head.next = new_item

        return head