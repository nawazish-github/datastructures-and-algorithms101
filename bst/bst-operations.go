package main

import (
	"fmt"
	"math"
)

type node struct {
	value       int
	left, right *node
}

type queue struct {
	queueChan chan node
}

func main() {

	//root := initializeBST()

	//fmt.Println("printing in order...")
	//inOrder(root)

	//fmt.Println("printing pre order")
	//preOrder(root)

	//fmt.Println("printing post order...")
	//postOrder(root)

	//x := make(chan node, 20)
	//qc := &queue{queueChan: x}

	//breadthFirstTraversal(root, qc)

	//root1 := initializeSmallBST()
	//root2 := initializeSmallMirrorBST()

	// root1 := initializeBST()
	// root2 := initializeMirrorBST()

	// isMirror := //isBSTsMirrorImage(root1, root2)
	// 	isBSTsMirror(root1, root2)

	// fmt.Println("Is Mirror: ", isMirror)
	fmt.Println("Is BST: ", isBST(initializeMirrorBST(), math.MinInt64, math.MaxInt64))
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

func breadthFirstTraversal(root *node, qc *queue) {
	for len(qc.queueChan) > 0 {
		if root != nil {
			fmt.Println(root.value)
			qc.queueChan <- *(root.left)
			qc.queueChan <- *root.right
		}
	}
}

func initializeSmallBST() *node {
	root := &node{10, nil, nil}
	n5 := &node{5, nil, nil}
	n20 := &node{20, nil, nil}

	root.left = n5
	root.right = n20
	return root
}

func initializeSmallMirrorBST() *node {
	root := &node{10, nil, nil}
	n5 := &node{5, nil, nil}
	n20 := &node{20, nil, nil}

	root.left = n20
	root.right = n5
	return root
}

func initializeMirrorBST() *node {
	root := &node{10, nil, nil}
	n5 := &node{5, nil, nil}
	n20 := &node{20, nil, nil}
	n3 := &node{3, nil, nil}
	n8 := &node{8, nil, nil}
	n15 := &node{15, nil, nil}
	n30 := &node{30, nil, nil}

	n5.left = n8
	n5.right = n3

	n20.left = n30
	n20.right = n15

	root.left = n20
	root.right = n5
	return root
}

func isBSTsMirrorImage(n1, n2 *node) bool {
	if n1 == nil && n2 == nil {
		return true
	}

	if n1 != nil && n2 != nil { //checks to avoid nil pointer exception
		if n1.value != n2.value {
			return false
		}
	} else {
		return false
	}

	return isBSTsMirrorImage(n1.left, n2.right) &&
		isBSTsMirrorImage(n1.right, n2.left)
}

func isBSTsMirror(n1, n2 *node) bool {
	if n1 == nil && n2 == nil {
		return true
	}

	if n1 == nil || n2 == nil {
		return false
	}
	return n1.value == n2.value &&
		isBSTsMirror(n1.left, n2.right) &&
		isBSTsMirror(n1.right, n2.left)
}

func isBST(rt *node, lowerLimit, upperLimit int) bool {
	if rt == nil {
		return true
	}

	if rt.value < lowerLimit || rt.value > upperLimit {
		return false
	}

	return isBST(rt.left, lowerLimit, rt.value) &&
		isBST(rt.right, rt.value, upperLimit)
}
