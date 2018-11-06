package main

import "fmt"

type node struct {
	value       int
	left, right *node
}

func main() {

	root := initializeBST()

	//fmt.Println("printing in order...")
	//inOrder(root)

	//fmt.Println("printing pre order")
	//preOrder(root)

	fmt.Println("printing post order...")
	postOrder(root)
}

func initializeBST() *node {
	root := &node{10, nil, nil}
	n5 := &node{5, nil, nil}
	n20 := &node{20, nil, nil}
	n3 := &node{3, nil, nil}
	n8 := &node{8, nil, nil}
	n15 := &node{15, nil, nil}
	n30 := &node{30, nil, nil}

	n5.left = n3
	n5.right = n8

	n20.left = n15
	n20.right = n30

	root.left = n5
	root.right = n20
	return root
}

func inOrder(root *node) {

	if root == nil {
		return
	}

	inOrder(root.left)
	fmt.Println(root.value)
	inOrder(root.right)
}

func preOrder(root *node) {
	if root == nil {
		return
	}

	fmt.Println(root.value)
	preOrder(root.left)
	preOrder(root.right)

}

func postOrder(root *node) {
	if root == nil {
		return
	}

	postOrder(root.left)
	postOrder(root.right)
	fmt.Println(root.value)

}
