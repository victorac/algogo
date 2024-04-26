package trees

type Data struct {
	Datetime    int    `json:"datetime"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Hierarchy struct {
	Data     Data        `json:"data"`
	Children []Hierarchy `json:"children"`
}

var nilData = &Data{-1, "", ""}
var nilHierarchy = &Hierarchy{*nilData, nil}

type AVL struct {
	root *AvlNode
}

type AvlNode struct {
	value      Data
	left       *AvlNode
	right      *AvlNode
	height     int
	difference int
}

func (tree *AVL) Print() *Hierarchy {
	var hierarchy = &Hierarchy{*nilData, nil}
	if tree.root != nil {
		tree.root.print(hierarchy)
	}
	return hierarchy
}

func (node *AvlNode) print(hierarchy *Hierarchy) {
	hierarchy.Data = node.value
	if node.left == nil && node.right == nil {
		return
	}
	hierarchy.Children = make([]Hierarchy, 2)
	if node.left != nil {
		childHierarchy := &Hierarchy{}
		node.left.print(childHierarchy)
		hierarchy.Children[0] = *childHierarchy
	} else {
		hierarchy.Children[0] = *nilHierarchy
	}
	if node.right != nil {
		childHierarchy := &Hierarchy{}
		node.right.print(childHierarchy)
		hierarchy.Children[1] = *childHierarchy
	} else {
		hierarchy.Children[1] = *nilHierarchy
	}
}

func (tree *AVL) Push(data Data) {
	if tree.root == nil {
		tree.root = &AvlNode{data, nil, nil, 0, 0}
	} else {
		tree.root.push(data)
	}
}

func (node *AvlNode) push(value Data) {
	var childHeight int
	if value.Datetime < node.value.Datetime {
		if node.left == nil {
			node.left = &AvlNode{value, nil, nil, 0, 0}
		} else {
			node.left.push((value))
		}
		childHeight = node.left.height
		if node.right == nil || (node.right.height < childHeight) {
			node.height = childHeight + 1
		} // else this node already has a greater height on the right
	} else {
		if node.right == nil {
			node.right = &AvlNode{value, nil, nil, 0, 0}
		} else {
			node.right.push((value))
		}
		childHeight = node.right.height
		if node.left == nil || (node.left.height < childHeight) {
			node.height = childHeight + 1
		} // else this node already has a greater height on the left
	}

	node.balance()
}

func (node *AvlNode) balance() {
	var leftHeight, rightHeight int
	if node.left != nil {
		leftHeight = node.left.height
	}
	if node.right != nil {
		rightHeight = node.right.height
	}
	node.difference = leftHeight - rightHeight
	if rightHeight > leftHeight+1 {
		// right heavy
		child := node.right
		var childLeftHeight, childRightHeight = -1, -1
		if child.left != nil {
			childLeftHeight = child.left.height
		}
		if child.right != nil {
			childRightHeight = child.right.height
		}
		if childLeftHeight > childRightHeight {
			// double rotation
			node.right.rightRotation2()
		}
		node.leftRotation2()
	} else if leftHeight > rightHeight+1 {
		// left heavy
		child := node.left
		var childLeftHeight, childRightHeight = -1, -1
		if child.left != nil {
			childLeftHeight = child.left.height
		}
		if child.right != nil {
			childRightHeight = child.right.height
		}
		if childRightHeight > childLeftHeight {
			// double rotation
			node.left.leftRotation2()
		}
		node.rightRotation2()
	}
}

func (node *AvlNode) rightRotation() {
	previousParentValue := node.value
	node.value = node.left.value
	node.left.value = previousParentValue

	node.left.left = node.left.right
	node.left.right = node.right
}

func (node *AvlNode) leftRotation() {
	previousParentValue := node.value
	node.value = node.right.value
	node.right.value = previousParentValue

	node.right.left = node.left
	node.left = node.right
	node.right = node.left.right
}

func (node *AvlNode) rightRotation2() {
	previousParentValue := node.value
	node.value = node.left.value
	node.left.value = previousParentValue

	previousLeftLeft := node.left.left
	node.left.left = node.left.right
	node.left.right = node.right
	node.right = node.left
	node.left = previousLeftLeft
}

func (node *AvlNode) leftRotation2() {
	previousParentValue := node.value
	node.value = node.right.value
	node.right.value = previousParentValue

	previousRightRight := node.right.right
	node.right.right = node.right.left
	node.right.left = node.left
	node.left = node.right
	node.right = previousRightRight
}
