package main

type Node struct {
	val  int
	next *Node
}

type LinkList struct {
	head *Node
	size int
}

func (l *LinkList) eeque(val int) {

	newNode := &Node{val: val, next: nil}
	if l.head == nil {
		l.head = newNode
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode

		l.size++

	}
}

func (l *LinkList) dnque() int {
	if l.head == nil {
		return 0
	}
	temp := l.head.val
	l.head = l.head.next
	l.size--
	return temp
}

func (l *LinkList) push(val int) int {
	newNode := &Node{val: val, next: nil}
	if l.head == nil {
		l.head = newNode
	} else {
		temp := l.head
		l.head = newNode
		l.head.next = temp
	}
	l.size++

	return l.size
}

func (l *LinkList) pop() int {
	if l.head == nil {
		return 0 // or another sentinel value to indicate the stack is empty
	}

	temp := l.head.val
	l.head = l.head.next
	l.size--
	return temp
}

func (l *LinkList) revdrse(val int) {
	var prev *Node
	var next *Node
	current := l.head //2
	for current.next != nil {
		next = current.next //3
		current.next = prev //nil
		prev = current      // prev = 2
		current = next      // current becomes
	}
	l.head = prev // set the node to the first node
}
