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
