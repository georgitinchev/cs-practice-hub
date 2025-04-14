# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution(object):
    def addTwoNumbers(self, l1, l2):
        dummy = ListNode(0)  # eh, just a dummy node to get us started
        current = dummy  # will use this to build the result
        carry = 0  # gotta keep track of the carry, like in 2nd grade math

        while l1 is not None or l2 is not None:  # keep going while there's stuff in either list
            x = l1.val if l1 is not None else 0  # if l1 ran out, just pretend it's 0
            y = l2.val if l2 is not None else 0  # same for l2
            sum = carry + x + y  # add the digits + whatever's being carried over
            carry = sum // 10  # carry the 1... or whatever's needed
            current.next = ListNode(sum % 10)  # only keep the ones place here
            current = current.next  # move to the next spot in result list

            if l1 is not None: l1 = l1.next  # move along the first list
            if l2 is not None: l2 = l2.next  # and the second too

        if carry > 0:  # oh right, might still have a carry left at the end
            current.next = ListNode(carry)

        return dummy.next  # skip the dummy and return the real head
