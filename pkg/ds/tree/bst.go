package tree

import "fmt"

type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

type BinarySearchTree struct {
	Root *TreeNode
}

func (bst *BinarySearchTree) Insert(data int) {
	bst.Root = insertNode(bst.Root, data)
}

func (bst *BinarySearchTree) Search(data int) *TreeNode {
	return searchNode(bst.Root, data)
}

func (bst *BinarySearchTree) Delete(data int) {
	bst.Root = delete(bst.Root, data)
}

func (bst *BinarySearchTree) InorderTraversal() {
	inorder(bst.Root)
}

func (bst *BinarySearchTree) PreOrderTraversal() {
	preorder(bst.Root)
}

func (bst *BinarySearchTree) PostOrderTraversal() {
	postorder(bst.Root)
}

func insertNode(root *TreeNode, data int) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}

	if root.Data < data {
		root.Right = insertNode(root.Right, data)
	} else {
		root.Left = insertNode(root.Left, data)
	}

	return root
}

func searchNode(root *TreeNode, data int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Data == data {
		return root
	}

	if data < root.Data {
		return searchNode(root.Left, data)
	} else {
		return searchNode(root.Right, data)
	}
}

func delete(root *TreeNode, data int) *TreeNode {
	if root == nil {
		return nil
	}

	if data < root.Data {
		root.Left = delete(root.Left, data)
	} else if data > root.Data {
		root.Right = delete(root.Right, data)
	} else {

		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}

		succesor := findInorderSuccesor(root.Right)
		root.Data = succesor.Data
		root.Right = delete(root.Right, succesor.Data)

	}

	return root
}

func findInorderSuccesor(root *TreeNode) *TreeNode {
	current := root

	for current.Left != nil {
		current = current.Left
	}
	return current
}

func inorder(root *TreeNode) {
	if root == nil {
		return
	}

	inorder(root.Left)
	fmt.Print(root.Data, " ")
	inorder(root.Right)
}

func preorder(root *TreeNode) {
	if root == nil {
		return
	}

	fmt.Print(root.Data, " ")
	preorder(root.Left)
	preorder(root.Right)
}

func postorder(root *TreeNode) {
	if root == nil {
		return
	}

	postorder(root.Left)
	postorder(root.Right)
	fmt.Print(root.Data, " ")
}
