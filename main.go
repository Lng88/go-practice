package main

import "fmt"

type LinkedList struct {
	head *Node
	tail *Node
}

type Node struct {
	next *Node
	val  string
}

func (ll *LinkedList) Insert(val string) {
	newNode := Node{val: val}
	if ll.head == nil {
		ll.head = &newNode
		ll.tail = &newNode
		return
	}

	curr := ll.head
	for {
		if curr.next == nil {
			break
		}
		curr = curr.next
	}
	curr.next = &newNode
	ll.tail = &newNode
}

// Update to remove node. Will have to keep track of ll index
func (ll *LinkedList) Remove(val string) {

}

func (ll *LinkedList) Scan() {
	curr := ll.head
	if ll.head == nil {
		return
	}
	for {
		fmt.Println("current value", curr.val)
		if curr.next == nil {
			break
		}

		curr = curr.next
	}
	fmt.Println("End of scan")
}

func main() {
	fmt.Println("Hello")
	list := LinkedList{}
	list.Insert("1")
	list.Insert("2")
	list.Insert("3")
	list.Insert("4")
	list.Insert("5")
	list.Scan()
	fmt.Printf("%#v", list)
}
