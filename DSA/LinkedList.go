package main

import (
	"errors"
	"fmt"
)

type Node[T comparable] struct {
	data T
	prev *Node[T]
	next *Node[T]
}

type LinkedList[T comparable] struct {
	head *Node[T]
	tail *Node[T]
}

// add data at the back of the linked list O(n)
func (list *LinkedList[T]) Append(data T) {
	newNode := &Node[T]{data: data}

	if list.head == nil {
		list.head = newNode
		list.tail = newNode
		return
	}

	newNode.prev = list.tail
	list.tail.next = newNode
	list.tail = newNode
}

// add data at the front of the linked list O(1)
func (list *LinkedList[T]) Push(data T) {
	if list.head == nil {
		newNode := &Node[T]{data: data}
		list.head = newNode
		list.tail = newNode
		return
	}
	newNode := &Node[T]{data: data, next: list.head}
	list.head = newNode
}

// add data at a specified index of a list O(n)
func (list *LinkedList[T]) InsertAt(index int, data T) (error) {
	if index >= list.Length() || index < 0 {
		return errors.New("index out of bound")
	}

	newNode := &Node[T]{data: data}

	if index == 0 {
		newNode.next = list.head
		list.head = newNode
		if list.tail == nil {
			list.tail = newNode
		}
		return nil
	}

	current := list.head
	for i := 0; i < index - 1; i++ {
		current = current.next
	}
	newNode.next = current.next
	newNode.prev = current
	current.next.prev = newNode
	current.next = newNode
	return nil
}

// delete the last node of the linked list
func (list *LinkedList[T]) Pop() (T, error) {
	if list.head == nil {
		var t T
		return t, errors.New("linked list is empty")
	}

	if list.head.next == nil {
        d := list.head.data
        list.head = nil
        list.tail = nil
        return d, nil
    }

    d := list.tail.data
	if list.tail.prev != nil {
        list.tail = list.tail.prev
        list.tail.next = nil
		return d, nil
    }
	// if the head becomes the tail and no longer has head
    list.head = nil
    list.tail = nil
    return d, nil
}

// returns a new linked list with an element that is common to all linked list
func (list *LinkedList[T]) Intersect(otherList *LinkedList[T]) (*LinkedList[T]) {
	visited := make(map[T]int)
	intersectNode := &LinkedList[T]{}

	X := list.head
	for X != nil {
		// count how often a data is visited
		visited[X.data]++
		X = X.next
	}

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

func (list *LinkedList[T]) Union(otherList *LinkedList[T]) (*LinkedList[T]) {
	visited := make(map[T]bool)
	unionNode := &LinkedList[T]{}

	X := list.head
	for X != nil {
		// marking every unique value as visited so there are no duplicates
		if !visited[X.data] {
			unionNode.Append(X.data)
			visited[X.data] = true
		}
		X = X.next
	}

	Y := otherList.head
	for Y != nil {
		// find Y node value that is not exist on the visited map
		if !visited[Y.data] {
			unionNode.Append(Y.data)
			visited[Y.data] = true
		}
		Y = Y.next
	}
	return unionNode
}

// print out the linked list
func (list *LinkedList[T]) Print() {
	current := list.head
	for current != nil {
		fmt.Printf("%v->", current.data)
		current = current.next
	}
	fmt.Println("nil")
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

	// append
	linkedList.Append(13)
	linkedList.Print()

	// push
	linkedList.Push(17)
	linkedList.Print()

	err := linkedList.InsertAt(6, 19)
	if err != nil {
		fmt.Println(err.Error())
	}
	linkedList.Print()

	for linkedList.head != nil {
		linkedList.Pop()
		linkedList.Print()
	}
}