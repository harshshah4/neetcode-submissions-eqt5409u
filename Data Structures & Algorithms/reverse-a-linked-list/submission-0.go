/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	temp := head
	head = head.Next
	temp.Next = nil
    for head != nil  {
		temp1 := head.Next
		head.Next = temp
		temp = head
		head = temp1
	}
	return temp
}
