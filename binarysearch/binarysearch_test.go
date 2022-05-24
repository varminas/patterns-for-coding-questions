package binarysearch

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	bst := &Bst{}

	values := []int{3, 5, 8, 0, 4, 2, 1, 6}

	for _, val := range values {
		bst.Insert(val)
	}

	fmt.Printf("Printing In-Order:\n")
	bst.InOrder()
	fmt.Printf("\nPrinting Pre-Order\n")
	bst.PreOrder()
	fmt.Printf("\nPrinting Post-Order\n")
	bst.PostOrder()
	fmt.Printf("\nPrinting with Breadth-First-Search\n")
	bst.TraverseLevelOrder()
	fmt.Println("\n\nFinding values:")
	bst.Find(3)
	bst.Find(33)
	fmt.Println("\n\nRemoving values")
	bst.Remove(5)
	bst.InOrder()
	//bst.Reset()
}
