package trees

import "fmt"

type Node struct {
	value      int
	leftChild  *Node
	rightChild *Node
}

type BST struct {
	root *Node
}

func (tree *BST) Push(value int) {
	newNode := &Node{value, nil, nil}
	if tree.root == nil {
		tree.root = newNode
	} else {
		sentinel := tree.root
		for sentinel != nil {
			if value < sentinel.value {
				if sentinel.leftChild == nil {
					sentinel.leftChild = newNode
					sentinel = nil
				} else {
					sentinel = sentinel.leftChild
				}
			} else {
				if sentinel.rightChild == nil {
					sentinel.rightChild = newNode
					sentinel = nil
				} else {
					sentinel = sentinel.rightChild
				}
			}
		}
	}
}

func (tree *BST) search(value int) *Node {
	sentinel := tree.root
	for sentinel != nil && sentinel.value != value {
		if value < sentinel.value {
			sentinel = sentinel.leftChild
		} else {
			sentinel = sentinel.rightChild
		}
	}
	return sentinel
}

func (tree *BST) Get(value int) int {
	node := tree.search(value)
	if node == nil {
		return -1
	}
	return node.value
}

func (node *Node) leftmostParent() *Node {
	parent := node
	sentinel := node.leftChild
	for sentinel != nil {
		if sentinel.leftChild != nil {
			parent = sentinel
		}
		sentinel = sentinel.leftChild
	}
	return parent
}

func (node *Node) rightmostParent() *Node {
	parent := node
	sentinel := node.rightChild
	for sentinel != nil {
		if sentinel.rightChild != nil {
			parent = sentinel
		}
		sentinel = sentinel.rightChild
	}
	return parent
}

func (tree *BST) Delete(value int) {
	sentinel := tree.root
	parent := tree.root
	child := ""
	for sentinel != nil && sentinel.value != value {
		if value < sentinel.value {
			parent = sentinel
			sentinel = sentinel.leftChild
			child = "left"
		} else {
			parent = sentinel
			sentinel = sentinel.rightChild
			child = "right"
		}
	}
	if sentinel == nil {
		// value not found
		return
	}

	if sentinel.leftChild != nil {
		greaterLeftChildParent := sentinel.leftChild.rightmostParent()
		greaterLeftChild := greaterLeftChildParent.rightChild
		if greaterLeftChild != nil {
			sentinel.value = greaterLeftChild.value
			greaterLeftChildParent.rightChild = greaterLeftChild.leftChild
		} else {
			sentinel.value = sentinel.leftChild.value
			sentinel.leftChild = sentinel.leftChild.leftChild
		}
	} else if sentinel.rightChild != nil {
		leastRightChildParent := sentinel.rightChild.leftmostParent()
		leastRightChild := leastRightChildParent.leftChild
		if leastRightChild != nil {
			sentinel.value = leastRightChild.value
			leastRightChildParent.leftChild = leastRightChild.rightChild
		} else {
			sentinel.value = sentinel.rightChild.value
			sentinel.rightChild = sentinel.rightChild.rightChild
		}
	} else if child == "left" {
		parent.leftChild = nil
	} else if child == "right" {
		parent.rightChild = nil
	} else if parent == tree.root {
		// no child nodes left
		tree.root = nil
	}

}

func (tree *BST) Print() {
	if tree.root == nil {
		fmt.Println("Empty tree")
	} else {
		fmt.Printf("Root: %v\n", tree.root.value)
		tree.root.print(true)
		fmt.Println("]")
	}
}

func (node *Node) print(leftmost bool) int {
	if node == nil {
		return -1
	}
	res := node.leftChild.print(leftmost)
	if leftmost && res == -1 {
		fmt.Printf("[%d", node.value)
	} else {
		fmt.Printf(", %d", node.value)
	}
	node.rightChild.print(false)
	return 0
}
