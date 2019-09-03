package main

import (
	"fmt"
)

type node struct {
	value int
	next *node
}

func (node *node) printLink() {
	if node == nil {
		return
	}
	for node != nil {
		fmt.Print(node.value, "  ")
		node = node.next
	}
	fmt.Println()
}

func (head *node) reversalLink() *node {
	cur := head
	var prev *node
	for cur != nil {
		cur.next, prev, cur = prev, cur, cur.next
	}
	prev.printLink()
	return prev
}

func (head *node) swapNodeInPair() {
	if head == nil || head.next == nil {
		return
	}
	pre := &node{}
	pre.next = head

	p := head.next

	for pre.next != nil && pre.next.next != nil {
		a := pre.next
		b := a.next
		pre.next, b.next, a.next = b, a, b.next
		pre = a
	}
	p.printLink()
}

func reverseKGroup(head *node, k int) *node {
	var pre *node
	cur, p := head, head

	groupSize, index := 0, 0

	// 将链表前k个节点分成一组
	for p != nil && groupSize < k {
		p = p.next
		groupSize++
	}

	if groupSize == k {
		// 反转前k个节点
		for cur != nil && index < k {
			cur.next, pre, cur = pre, cur, cur.next
			index++
		}

		if cur != nil {
			head.next = reverseKGroup(cur, k)
		}
		return pre
	} else {
		return head
	}
}

func main() {
	head := &node{
		value: 1,
	}
	node2 := &node{
		value: 2,
	}
	node3 := &node{
		value: 3,
	}
	node4 := &node{
		value: 4,
	}
	node5 := &node{
		value: 5,
	}
	node6 := &node{
		value: 6,
	}
	node7 := &node{
		value: 7,
	}
	head.next = node2
	node2.next = node3
	node3.next = node4
	node4.next = node5
	node5.next = node6
	node6.next = node7

	head.printLink()

	head = head.reversalLink()

	//head.swapNodeInPair()

	head = reverseKGroup(head, 4)
	head.printLink()
}
