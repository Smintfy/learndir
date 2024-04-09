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

type Node[T comparable] struct {
	data T
	next *Node[T]
}

type LinkedList[T comparable] struct {
	head *Node[T]
}

// add data at the back of the linked list
// O(n) as we need to traverse through the list to find the last node
func (list *LinkedList[T]) append(data T) {
	newNode := &Node[T]{data: data}

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
func (list *LinkedList[T]) push(data T) {
	// new entry
	if list.head == nil {
		newNode := &Node[T]{data: data}
		list.head = newNode
		return 
	}

	// the new node points to the list head
	// and the list head refers to the new node
	newNode := &Node[T]{data: data, next: list.head}
	list.head = newNode
}

// add data at a specified index of a list
// O(n)
func (list *LinkedList[T]) insertAt(index int, data T) (error) {
	if index > list.len() || index < 0 {
		return errors.New("index out of bound")
	}

	newNode := &Node[T]{data: data}
	
	// insert at the beginning
	if index == 0 {
		newNode.next = list.head
		list.head = newNode
		return nil
	}

	current := list.head
	count := 0
	for current != nil {
		// [x_cur] -> [new] -> [x_next]
		// we are inserting the value between current node and the next node of the current node
		if count + 1 == index {
			newNode.next = current.next // [new] -> [x_next]
			current.next = newNode // [x_cur] -> [new]
			return nil
		}
		current = current.next
		count++
	}
	return nil
}

// delete the last node of the linked list
func (list *LinkedList[T]) pop() (T, error) {
	if list.head == nil {
		var t T
		return t, errors.New("linked list is empty")
	}

	// if the list length is 1
	// [x] -> nil
	// then we just make the head nil
	if list.head.next == nil {
		d := list.head.data 
		list.head = nil
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
	return d, nil
}

// returns a new linked list with an element that is common to all linked list
// X = [1] -> [3] -> [5] -> [7] -> [9]
// Y = [2] -> [3] -> [6] -> [9] -> [14]
// X intersect Y = [3] -> [9]
func (list *LinkedList[T]) intersect(otherList *LinkedList[T]) (*LinkedList[T]) {
	// if the other list is bigger then we switch for X and Y
	if list.len() < otherList.len() {
		temp := list
		list = otherList
		otherList = temp
	}

	// this shit is O(n^2)
	// compare each of the bigger list value to the each of the smaller list value
	intersectNode := &LinkedList[T]{}
	X := list.head
	for X != nil {
		Y := otherList.head
		for Y != nil {
			if Y.data == X.data {
				intersectNode.append(Y.data)
			}
			Y = Y.next
		}
		X = X.next
	}
	return intersectNode
}

// returns a new linked list that contains all the items from the original linked list.
// X = [1] -> [3] -> [5] -> [7] -> [9]
// Y = [2] -> [3] -> [6] -> [9] -> [14]
// X union Y = [1] -> [2] -> [3] -> [5] -> [6] -> [7] -> [9] -> [14]
func (list *LinkedList[T]) union(otherList *LinkedList[T]) (*LinkedList[T]) {
	visited := make(map[T]bool)
	unionNode := &LinkedList[T]{}

	// we visit every node in X
	// marking every unique value as visited so there are no duplicates
	X := list.head
	for X != nil {
		if !visited[X.data] {
			unionNode.append(X.data)
			visited[X.data] = true
		}
		X = X.next
	}

	// now we just find Y node value that is not exist on the visited map
	Y := otherList.head
	for Y != nil {
		if !visited[Y.data] {
			unionNode.append(Y.data)
			visited[Y.data] = true
		}
		Y = Y.next
	}
	return unionNode
}

// print out the linked list
// x1 -> ... -> xi -> nil
func (list *LinkedList[T]) print() {
	current := list.head
	for current != nil {
		fmt.Printf("%v -> ", current.data)
		current = current.next
	}
	fmt.Printf("nil")
	fmt.Println()
}

// get the length of the linked list
func (list *LinkedList[T]) len() int {
	count := 0
	current := list.head
	for current != nil {
		current = current.next
		count++
	}
	return count
}

func main() {
	linkedList := &LinkedList[int]{}

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

	fmt.Printf("\n\nInserting: \n")

	linkedList.insertAt(3, 19)
	linkedList.print()
	fmt.Printf("The length of list: %d\n", linkedList.len())

	A := &LinkedList[int]{}
	B := &LinkedList[int]{}

	A_datas := [5]int{1, 8, 3, 4, 5}
	for _, data := range A_datas {
		A.append(data)
	}

	B_datas := [6]int{6, 4, 7, 3, 1, 8}
	for _, data := range B_datas {
		B.append(data)
	}

	fmt.Printf("\n\nA & B:\n")
	A.print()
	B.print()

	C := A.intersect(B)

	fmt.Printf("\nIntersect: \n")
	C.print()

	D := A.union(B)
	fmt.Printf("\nUnion: \n")
	D.print()
}