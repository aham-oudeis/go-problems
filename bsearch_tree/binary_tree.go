package main

import "fmt"

type TreeNode struct {
	val        int
	leftChild  *TreeNode
	rightChild *TreeNode
}

func newTreeNode(val int) *TreeNode {
	return &TreeNode{val: val}
}

type SearchTree struct {
	root *TreeNode
}

func newSearchTree(rootVal int) *SearchTree {
	root := newTreeNode(rootVal)
	tree := SearchTree{root}

	return &tree
}

func main() {
	node1 := newTreeNode(25)
	node2 := newTreeNode(72)

	tree := newSearchTree(50)
	root := tree.root

	root.leftChild = node1
	root.rightChild = node2

	fmt.Println(root)
}
