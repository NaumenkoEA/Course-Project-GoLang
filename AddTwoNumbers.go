package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	leader := 0
	head := new(ListNode)
	cur := head
	for l1 != nil || l2 != nil || leader != 0 {
		num1 := 0
		num2 := 0
		if l1 != nil {
			num1, l1 = l1.Val, l1.Next
		}
		if l2 != nil {
			num2, l2 = l2.Val, l2.Next
		}
		num := num1 + num2 + leader
		leader = num / 10
		cur.Next = &ListNode{num % 10, nil}
		cur = cur.Next
	}
	return head.Next

}
