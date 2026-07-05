package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	res := head
	check := head
	count := 0

	for check != nil {
		count++
		check = check.Next
	}
	if count < 2 {
		return nil
	}

	del := head.Next

	if count-n == 0 {
		return del
	}

	for i := 0; i < count-n; i++ {
		if i == count-n-1 {
			head.Next = del.Next
			break
		}
		head = head.Next
		del = del.Next
	}

	return res
}
