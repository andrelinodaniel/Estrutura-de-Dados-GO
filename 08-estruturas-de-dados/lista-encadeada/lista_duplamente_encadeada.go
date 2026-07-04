package main

import "fmt"


type Node struct {
	Value int
	Prev  *Node
	Next  *Node
}


type DoublyLinkedList struct {
	Head *Node
	Tail *Node
	Size int
}


func (l *DoublyLinkedList) Append(value int) {
	newNode := &Node{Value: value}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
	
		newNode.Prev = l.Tail
		l.Tail.Next = newNode
		l.Tail = newNode
	}
	l.Size++
}


func (l *DoublyLinkedList) RemoveByValue(value int) bool {
	current := l.Head

	for current != nil {
		if current.Value == value {

			if current.Prev != nil {
				current.Prev.Next = current.Next
			} else {

				l.Head = current.Next
			}

			if current.Next != nil {
				current.Next.Prev = current.Prev
			} else {
		
				l.Tail = current.Prev
			}

			l.Size--
			return true
		}
		current = current.Next
	}
	return false
}

func (l *DoublyLinkedList) Print() {
	current := l.Head
	fmt.Print("[")
	for current != nil {
		fmt.Printf("%d", current.Value)
		if current.Next != nil {
			fmt.Print(" <-> ")
		}
		current = current.Next
	}
	fmt.Println(" -> nil]")
}


func (l *DoublyLinkedList) PrintReverse() {
	current := l.Tail
	fmt.Print("[")
	for current != nil {
		fmt.Printf("%d", current.Value)
		if current.Prev != nil {
			fmt.Print(" <-> ")
		}
		current = current.Prev
	}
	fmt.Println(" -> nil]")
}

func main() {
	list := &DoublyLinkedList{}
	list.Append(101)
	list.Append(202)
	list.Append(303)

	fmt.Println("Lista antes da remoção:")
	list.Print()

	removed := list.RemoveByValue(202)
	fmt.Println("Pedido 202 removido:", removed)

	fmt.Println("Lista depois da remoção:")
	list.Print()

	fmt.Println("Lista em ordem inversa:")
	list.PrintReverse()
}
