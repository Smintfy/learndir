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
func (list *LinkedList[T]) Append(data T) {
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
func (list *LinkedList[T]) Push(data T) {
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
func (list *LinkedList[T]) InsertAt(index int, data T) (error) {
	if index > list.Length() || index < 0 {
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
func (list *LinkedList[T]) Pop() (T, error) {
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
func (list *LinkedList[T]) Intersect(otherList *LinkedList[T]) (*LinkedList[T]) {
	// count the occurrences of the node value in list
	visited := make(map[T]int)
	intersectNode := &LinkedList[T]{}

	// we visit every node in X
	// increment the count value on each visit
	X := list.head
	for X != nil {
		visited[X.data]++
		X = X.next
	}

	// we compare every visited node to the Y node
	Y := otherList.head
	for Y != nil {
		if k, v := visited[Y.data]; v && k > 0 {
			intersectNode.Push(Y.data)
			// decrement after pushing the data so no duplicate occurs
			visited[Y.data]--
		}
		Y = Y.next
	}
	return intersectNode
}

// returns a new linked list that contains all the items from the original linked list.
// X = [1] -> [3] -> [5] -> [7] -> [9]
// Y = [2] -> [3] -> [6] -> [9] -> [14]
// X union Y = [1] -> [2] -> [3] -> [5] -> [6] -> [7] -> [9] -> [14]
func (list *LinkedList[T]) Union(otherList *LinkedList[T]) (*LinkedList[T]) {
	visited := make(map[T]bool)
	unionNode := &LinkedList[T]{}

	// we visit every node in X
	// marking every unique value as visited so there are no duplicates
	X := list.head
	for X != nil {
		if !visited[X.data] {
			unionNode.Append(X.data)
			visited[X.data] = true
		}
		X = X.next
	}

	// now we just find Y node value that is not exist on the visited map
	Y := otherList.head
	for Y != nil {
		if !visited[Y.data] {
			unionNode.Append(Y.data)
			visited[Y.data] = true
		}
		Y = Y.next
	}
	return unionNode
}

// print out the linked list
// x1 -> ... -> xi -> nil
func (list *LinkedList[T]) Print() {
	current := list.head
	for current != nil {
		fmt.Printf("%v->", current.data)
		current = current.next
	}
	fmt.Printf("nil")
	fmt.Println()
}

// get the length of the linked list
func (list *LinkedList[T]) Length() int {
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
		linkedList.Append(data)
	}
	linkedList.Print()
	fmt.Printf("The length of list: %d\n", linkedList.Length())

	// data reserved for testing the append
	d_append := 13
	fmt.Printf("\n\nAppending %d: \n", d_append)

	linkedList.Append(d_append)
	linkedList.Print()
	fmt.Printf("The length of list: %d\n", linkedList.Length())

	// data reserved for testing the push
	d_push := 17
	fmt.Printf("\n\nPushing %d: \n", d_push)

	linkedList.Push(d_push)
	linkedList.Print()
	fmt.Printf("The length of list: %d\n", linkedList.Length())

	fmt.Printf("\n\nDeleting last node: \n")

	linkedList.Pop()
	linkedList.Print()
	fmt.Printf("The length of list: %d\n", linkedList.Length())

	fmt.Printf("\n\nInserting: \n")

	linkedList.InsertAt(3, 19)
	linkedList.Print()
	fmt.Printf("The length of list: %d\n", linkedList.Length())

	A := &LinkedList[int]{}
	B := &LinkedList[int]{}

	A_datas := [6]int{1, 8, 3, 4, 5, 4}
	for _, data := range A_datas {
		A.Append(data)
	}

	B_datas := [4]int{6, 4, 8, 8}
	for _, data := range B_datas {
		B.Append(data)
	}

	fmt.Printf("\n\nA & B:\n")
	A.Print()
	B.Print()

	C := A.Intersect(B)

	fmt.Printf("\nIntersect: \n")
	C.Print()

	D := A.Union(B)
	fmt.Printf("\nUnion: \n")
	D.Print()
}