package lists

type Node struct {
	Value int
	next  *Node
}

type LinkedList struct {
	head   *Node
	tail   *Node
	Length int
}

func (list *LinkedList) Push(value int) {
	newNode := &Node{value, nil}
	if list.Length == 0 {
		list.head = newNode
	} else {
		list.tail.next = newNode
	}
	list.tail = newNode
	list.Length++
}

func (list *LinkedList) Pop() {
	sentinel := list.head
	for i := 0; i < list.Length-2; i++ {
		sentinel = sentinel.next
	}
	sentinel.next = nil
	list.tail = sentinel
	list.Length--
}

func (list *LinkedList) goTo(index int) *Node {
	sentinel := list.head
	for i := 0; (sentinel != nil) && (i < index); i++ {
		sentinel = sentinel.next
	}
	return sentinel
}

func (list *LinkedList) Get(index int) Node {
	return *(list.goTo(index))
}

func (list *LinkedList) Delete(index int) {
	if index == list.Length-1 {
		list.Pop()
		return
	}
	if index == 0 {
		list.head = list.head.next
	} else {
		sentinel := list.goTo(index - 1)
		sentinel.next = sentinel.next.next
	}
	list.Length--

}
