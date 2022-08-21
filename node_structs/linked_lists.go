package main

import "fmt"

type Node struct {
	val  int
	next *Node
	prev *Node
}

func newNode(val int) *Node {
	return &Node{val: val}
}

type List struct {
	head   *Node
	tail   *Node
	length int
}

func newList() *List {
	return &List{}
}

func (lst *List) enque(num int) {
	curTail := lst.tail

	//if no node exists in the list
	if curTail == nil {
		lst.tail = newNode(num)
		lst.head = lst.tail
		lst.length += 1
		return
	}

	//if there is already an existing node
	newTail := newNode(num)
	curTail.next = newTail
	lst.tail = newTail
	lst.length += 1
}

func (lst *List) deque() (int, bool) {
	if lst.head == nil {
		return 0, false
	}

	curHead := lst.head
	newHead := curHead.next

	lst.head = newHead
	lst.length -= 1

	return curHead.val, true
}

func (lst *List) toSlice() []int {
	col := make([]int, 0, lst.length)

	curNode := lst.head

	for i := 0; i < lst.length; i++ {
		col = append(col, curNode.val)
		curNode = curNode.next
	}

	return col
}

func main() {
	fmt.Println("vim-go")
	q := newList()

	q.enque(4)
	q.enque(8)
	q.enque(9)
	q.enque(6)
	fmt.Println(q.head)
	fmt.Println(q.toSlice())
	q.deque()
	fmt.Println(q.head)
	q.deque()
	fmt.Println(q.head)
	fmt.Println(q.toSlice())
	fmt.Println(q.head)
}
