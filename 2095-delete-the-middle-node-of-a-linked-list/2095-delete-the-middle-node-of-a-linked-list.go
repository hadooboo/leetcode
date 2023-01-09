/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
	if head.Next == nil {
		return nil
	}

	n := head
	for p, c := head, 0; p != nil; p, c = p.Next, c+1 {
		if c >= 3 && c%2 == 1 {
			n = n.Next
		}
	}
	n.Next = n.Next.Next

	return head
}