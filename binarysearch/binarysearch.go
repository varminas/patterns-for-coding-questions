package binarysearch

import (
	"container/list"
	"fmt"
	"github.com/varminas/patterns-for-coding-questions/models"
)

type Bst struct {
	Root *models.BstNode
}

func (bst *Bst) Reset() {
	bst.Root = nil
}

func (bst *Bst) Insert(value int) {
	bst.insertRec(bst.Root, value)
}

func (bst *Bst) insertRec(node *models.BstNode, value int) *models.BstNode {
	if bst.Root == nil {
		bst.Root = &models.BstNode{
			Value: value,
		}
		return bst.Root
	}

	if node == nil {
		return &models.BstNode{Value: value}
	}
	if value <= node.Value {
		node.Left = bst.insertRec(node.Left, value)
	}
	if value > node.Value {
		node.Right = bst.insertRec(node.Right, value)
	}

	return node
}

func (bst *Bst) InOrder() {
	bst.inOrderRec(bst.Root)
}

func (bst *Bst) inOrderRec(node *models.BstNode) {
	if node != nil {
		bst.inOrderRec(node.Left)
		fmt.Printf("%d ", node.Value)
		bst.inOrderRec(node.Right)
	}
}

func (bst *Bst) PreOrder() {
	bst.preOrderRec(bst.Root)
}

func (bst *Bst) preOrderRec(node *models.BstNode) {
	if node != nil {
		fmt.Printf("%d ", node.Value)
		bst.preOrderRec(node.Left)
		bst.preOrderRec(node.Right)
	}
}

func (bst *Bst) PostOrder() {
	bst.postOrderRec(bst.Root)
}

func (bst *Bst) postOrderRec(node *models.BstNode) {
	if node != nil {
		bst.postOrderRec(node.Left)
		bst.postOrderRec(node.Right)
		fmt.Printf("%d ", node.Value)
	}
}

func (bst *Bst) Find(value int) {
	node := bst.findRec(bst.Root, value)
	if node == nil {
		fmt.Printf("Value %d not found.\n", value)
	} else {
		fmt.Printf("Value %d found.\n", value)
	}
}

func (bst *Bst) findRec(node *models.BstNode, value int) *models.BstNode {
	if node == nil {
		return nil
	}
	if value == node.Value {
		return bst.Root
	}
	if value < node.Value {
		return bst.findRec(node.Left, value)
	}
	return bst.findRec(node.Right, value)
}

func (bst *Bst) Remove(value int) {
	bst.removeRec(bst.Root, value)
	fmt.Println("")
}

func (bst *Bst) removeRec(node *models.BstNode, value int) *models.BstNode {
	if node == nil {
		return nil
	} else if value < node.Value {
		node.Left = bst.removeRec(node.Left, value)
	} else if value > node.Value {
		node.Right = bst.removeRec(node.Right, value)
	} else {
		fmt.Printf("Found value %d for removing.\n", value)
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		}

		// If the node has two children
		// Place the inorder successor in position of the node to be deleted
		node.Value = bst.minValue(node.Right)

		// Delete the inorder successor
		node.Right = bst.removeRec(node.Right, node.Value)
	}
	return node
}

func (bst *Bst) minValue(node *models.BstNode) int {
	if node.Left == nil {
		return node.Value
	} else {
		return bst.minValue(node.Left)
	}
}

func (bst *Bst) TraverseLevelOrder() {
	bst.traverseLevelOrderRec(bst.Root)
}

func (bst *Bst) traverseLevelOrderRec(root *models.BstNode) {
	if bst.Root == nil {
		return
	}

	nodes := list.New()
	nodes.PushBack(root)

	for nodes.Len() > 0 {
		node := nodes.Front()
		nodeVal := node.Value.(*models.BstNode)
		nodes.Remove(node)
		fmt.Printf("%d ", nodeVal.Value)
		if nodeVal.Left != nil {
			nodes.PushBack(nodeVal.Left)
		}
		if nodeVal.Right != nil {
			nodes.PushBack(nodeVal.Right)
		}
	}
}
