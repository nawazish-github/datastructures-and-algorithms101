package tree

//Bst is the Binary Search Tree representation.
type Bst struct {
	Size int
	Root *Node
}

//Node represents each Node in the Binary Search Tree
type Node struct {
	value      int
	leftChild  *Node
	rightChild *Node
}

//Insert inserts elements into the BST.
func (b *Bst) Insert(num int, root *Node) {
	b.Root = b.insertIntoBST(num, b.Root)
}

func (b *Bst) insertIntoBST(num int, root *Node) *Node {
	node := Node{value: num}
	if root == nil {
		root = &node
		b.Size++
		return root
	}
	if num < root.value {
		root.leftChild = b.insertIntoBST(num, root.leftChild)
	} else {
		root.rightChild = b.insertIntoBST(num, root.rightChild)
	}
	return root
}
