package main

import (
	"fmt"
)

type Tree struct {
	value    int
	leftNode *Tree
	rigtNode *Tree
}

func main() {
	root := InitializeNodes()
	PrintTree(root)
}

func InitializeNodes() *Tree {
	root := &Tree{value: 5}
	root = AddNode(root, &Tree{value: 3})
	root = AddNode(root, &Tree{value: 8})
	root = AddNode(root, &Tree{value: 2})
	root = AddNode(root, &Tree{value: 3})
	root = AddNode(root, &Tree{value: 7})
	root = AddNode(root, &Tree{value: 9})
	return root
}

func PrintTree(root *Tree) {
	if root == nil {
		return
	}

	fmt.Println("root", root.value)
	if root.leftNode != nil {
		fmt.Println("routed left -> ", root.value)
		PrintTree(root.leftNode)
	}
	if root.rigtNode != nil {
		fmt.Println("routed right -> ", root.value)
		PrintTree(root.rigtNode)
	}
}

func AddNode(root *Tree, node *Tree) *Tree {
	if root == nil {
		return &Tree{}
	}

	if node.value < root.value && root.leftNode == nil {
		root.leftNode = node
		return root
	} else if node.value < root.value && root.leftNode != nil {
		AddNode(root.leftNode, node)
	} else if node.value >= root.value && root.rigtNode == nil {
		root.rigtNode = node
		return root
	} else if node.value >= root.value && root.rigtNode != nil {
		AddNode(root.rigtNode, node)
	}

	return root
}
