package main

import "fmt"

type node struct {
	data int
	next *node
}

func AddNode(head *node, value int) *node {
	newNode := &node{data: value}
	//in case head have nothing or head is empty
	if head == nil {
		head = newNode
	} else {
		head1 := head
		for head1.next != nil {
			head1 = head1.next
		}
		head1.next = newNode
	}
	return head
}
func main() {
	var head *node
	head = AddNode(head, 5)
	head = AddNode(head, 15)
	head = AddNode(head, 115)
	head = AddNode(head, 125)
	for head != nil {
		fmt.Print(head.data, "-->")
		head = head.next
	}

}
