package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{}
	n := 2

	dummy := &ListNode{Val: -1}
	dummy.Next = head
	cur := head
	latter := dummy
	former := head

	for i := 0; i < n; i++ {
		former = former.Next
	}

	for former != nil {
		former = former.Next
		latter = cur
		cur = cur.Next
	}
	latter.Next = cur.Next

}
