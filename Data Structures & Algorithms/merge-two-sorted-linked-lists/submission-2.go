/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil  {
		return list2
	} else if list2 == nil {
		return list1
	}
	curr := &ListNode{}
	head := curr
    for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			curr.Next = list1
			curr = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			curr = list2
			list2 = list2.Next
		}
	}
	if list1 != nil {
		curr.Next = list1
	}
	if list2 != nil {
		curr.Next = list2
	}
	return head.Next
}
