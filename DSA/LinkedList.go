package main

import (
	"errors"
	"fmt"
)

/*
[x1] [p] -> [x2] [p] -> nil

[p] holds the address or points to the next node.
[x1] refers as a head of the linked list and serves as an entry point.

nil != null
*/


type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

// add data at the back of the linked list
// O(n) as we need to traverse through the list to find the last node
func (list *LinkedList) append(data int) {
	newNode := &Node{data: data, next: nil}

	// new entry
	if list.head == nil {
		list.head = newNode
		return
	}

	// iterate through the node until the next pointer is nil
	current := list.head
	for current.next != nil {
		current = current.next
	}

	current.next = newNode
}

// add data at the front of the linked list
// O(1) as we just need to perform constant operations of adding the new Node
func (list *LinkedList) push(data int) {
	// new entry
	if list.head == nil {
		newNode := &Node{data: data, next: nil}
		list.head = newNode
		return 
	}

	// the new node points to the list head
	// and the list head refers to the new node
	newNode := &Node{data: data, next: list.head}
	list.head = newNode
}

// delete the last node of the linked list
func (list *LinkedList) pop() (int, error) {
	if list.head == nil {
		return 0, errors.New("linked list is empty")
	}

	// if the list length is 1
	// [x] -> nil
	// then we just make the head nil
	if list.head.next == nil {
		d := list.head.data 
		list.head = nil
		fmt.Println("successfully deleted last node")
		return d, nil
	}

	// iterate through the list
	// when the next of the next node is nil
	// it turns to ... -> [xi] -> nil
	current := list.head
	for current.next.next != nil {
		current = current.next
	}
	d := current.next.data
	current.next = nil
	fmt.Println("successfully deleted last node")
	return d, nil
}

// print out the linked list
// x1 -> ... -> xi -> nil
func (list *LinkedList) print() {
	current := list.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Printf("nil")
	fmt.Println()
}

// get the length of the linked list
func (list *LinkedList) len() int {
	count := 0
	current := list.head
	for current != nil {
		count++
		current = current.next
	}
	return count
}

func main() {
	linkedList := &LinkedList{}

	datas := [5]int{2, 3, 5, 7, 11}

	// populate
	for _, data := range datas {
		linkedList.append(data)
	}
	linkedList.print()
	fmt.Printf("The length of list: %d\n", linkedList.len())

	// data reserved for testing the append
	d_append := 13
	fmt.Printf("\n\nAppending %d: \n", d_append)

	linkedList.append(d_append)
	linkedList.print()
	fmt.Printf("The length of list: %d\n", linkedList.len())

	// data reserved for testing the push
	d_push := 17
	fmt.Printf("\n\nPushing %d: \n", d_push)

	linkedList.push(d_push)
	linkedList.print()
	fmt.Printf("The length of list: %d\n", linkedList.len())

	fmt.Printf("\n\nDeleting last node: \n")

	linkedList.pop()
	linkedList.print()
	fmt.Printf("The length of list: %d\n", linkedList.len())
}