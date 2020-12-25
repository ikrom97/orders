package main

import (
	"fmt"
)
type Node struct {
	name string
	priority int64
	previous *Node
	next *Node
}
type ListNode struct {
	size int64
	firstNode *Node
	lastNode *Node
}
func (receiver *ListNode) PushBack(name string, priority int64) {
	if receiver.size == 0 {
		current := Node {
			name: name,
			priority: priority,
			previous: nil,
			next: nil,
		}
		receiver.size++
		receiver.firstNode = &current
		receiver.lastNode = &current
		return
	}	
	current := Node {
		name: name,
		priority: priority,
		previous: receiver.lastNode,
		next: nil,
	}
	receiver.lastNode.next = &current
	receiver.lastNode = &current
	receiver.size++
}
func (receiver *ListNode) PushFront(name string, priority int64) {
	if receiver.size == 0 {
		current := Node {
			name: name,
			priority: priority,
			previous: nil,
			next: nil,
		}
		receiver.size++
		receiver.firstNode = &current
		receiver.lastNode = &current
		return
	}
	current := Node {
		name: name,
		priority: priority,
		previous: nil,
		next: receiver.firstNode,
	}
	receiver.firstNode.previous = &current
	receiver.firstNode = &current
	receiver.size++
}
func (receiver *ListNode) RemoveBack() {
	if receiver.size == 0 {
		return
	} else if receiver.size == 1 {
		receiver.firstNode = nil
		receiver.lastNode = nil
		receiver.size--
		return
	}
	receiver.lastNode.previous.next = nil
	current := receiver.lastNode.previous
	receiver.lastNode.previous = nil
	receiver.lastNode = current
	receiver.size--
}
func (receiver *ListNode) RemoveFront() {
	if receiver.size == 0 {
		return
	} else if receiver.size == 1 {
		receiver.firstNode = nil
		receiver.lastNode = nil
		receiver.size--
		return
	}
	receiver.firstNode.next.previous = nil
	current := receiver.firstNode.next
	receiver.firstNode.next = nil
	receiver.firstNode = current
	receiver.size--
}
func (receiver *ListNode) PRINT(comment string) {
	fmt.Println(comment)
	current := *receiver.firstNode
	fmt.Println(current)
	for current.next != nil {
		current = *current.next
		fmt.Println(current)
	}
	fmt.Println("Лист:",*receiver)
}
func main() {
	// У нас есть очередь из 2 человек:
	order := []Node{
		{
			name: "Ikrom",
			priority: 9,
			previous: nil,
			next: nil,
		},
		{
			name: "Suhrob",
			priority: 4,
			previous: nil,
			next: nil,
		},
	}
	orderList := ListNode {
		size: int64(len(order)),
		firstNode: &order[0],
		lastNode: &order[len(order)-1],
	}
	// Сортируем очередь:
	for i := 0; i < len(order); i++ {
		for g := 1; g < len(order); g++ {
			if order[g-1].priority < order[g].priority {
				help := order[g-1]
				order[g-1] = order[g]
				order[g] = help
			}
		}
	}
	for i := 1; i < len(order); i++ {
		order[i-1].next = &order[i]
	}
	for i := 1; i < len(order); i++ {
		order[i].previous = &order[i-1]
	}
	// Выводим в консоль:
	orderList.PRINT("\n Сортированный очередь без модификаций:")
	// Добавляем в конец очереди:
	orderList.PushBack("Nekruz", 10)
	orderList.PRINT("\n Очередь после метода PushBack:")
	// Удаляем последний элемент:
	orderList.RemoveBack()
	orderList.PRINT("\n Очередь после метода RemoveBack")
	// Добавляем в начало очереди:
	orderList.PushFront("Surush", 1)
	orderList.PRINT("\n Очередь после метода PushFront:")
	// Удаляем первый элемент:
	orderList.RemoveFront()
	orderList.PRINT("\n Очередь после метода RemoveFront:")
}

