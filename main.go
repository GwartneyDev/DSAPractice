package main

import (
	"fmt"
)

// Define the Node structure
type Node struct {
	data int
	next *Node
}

// Define the LinkedList structure
type LinkedList struct {
	head *Node
}

// Initialize a new empty linked list
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// Add a new node at the end of the linked list
func (ll *LinkedList) Append(data int) {
	newNode := &Node{data: data, next: nil}

	if ll.head == nil {
		ll.head = newNode
		return
	}

	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

// Display the linked list
func (ll *LinkedList) Display() {
	current := ll.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	// Create a new linked list
	ll := NewLinkedList()

	// Add elements to the linked list
	ll.Append(10)
	ll.Append(20)
	ll.Append(30)

	// Display the linked list
	fmt.Println("Linked List:")
	ll.Display()
}


///UPDATED 

package main

type Node struct {
	value int
	next  *Node
}

type linkedList struct {
	head *Node
}

func (l *linkedList) add(value int) {
	newNode := &Node{value: value} // Create a new node with the given value

	if l.head == nil {
		// If the list is empty, set the new node as the head
		l.head = newNode
	} else {
		// Otherwise, traverse the list to find the last node
		current := l.head
		for current.next != nil {
			current = current.next // Move to the next node
		}
		// After the loop, current points to the last node in the list
		current.next = newNode // Link the new node to the end of the list
	}
}

// adds to the front of a linkd list
func (l *linkedList) push(value int) {

	newNode := &Node{value: value}

	if l.head == nil {
		l.head = newNode

	} else {
		temp := l.head
		l.head = newNode
		l.head.next = temp
	}
}

