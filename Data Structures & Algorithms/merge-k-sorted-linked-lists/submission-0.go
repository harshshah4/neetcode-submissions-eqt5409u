
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	}
	return split(lists, 0, n-1)
}

func split(lists []*ListNode, start, end int) *ListNode{
	if end - start <= 0 {
		return lists[end]
	}
	list1 := split(lists, start, (start+end)/2)
	list2 := split(lists, (start+end)/2 + 1 , end)
	list := merge(list1, list2)
	return list
}

func merge(list1, list2 *ListNode) *ListNode{
	var head, tail *ListNode
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			if head == nil {
				head = list1
				tail = list1
			} else {
				tail.Next = list1
				tail = tail.Next
			}
			list1 = list1.Next
			tail.Next = nil
		} else {
			if head == nil {
				head = list2
				tail = list2
			} else {
				tail.Next = list2
				tail = tail.Next
			}
			list2 = list2.Next
			tail.Next = nil
		}
	}
	if list1 != nil {
		if head == nil {
			head = list1
		} else {
			tail.Next = list1
		}
	}
	if list2 != nil {
		if head == nil {
			head = list2
		} else {
			tail.Next = list2
		}
	}
	return head
}

