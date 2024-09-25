package algo_test

import (
	"algo"
	"testing"
)

func TestBTreeInsert(t *testing.T) {
	testNode := algo.Node{Val: 10, Left: nil, Right: nil}
	testNode.InsertAll([]int{5, 15, 2, 7, 12, 18})
	if testNode.Val != 10 {
		t.Errorf("Error BTree insert test: Expected Val = 10 for Root node, got %d", testNode.Val)
	}
	if testNode.Left.Val != 5 {
		t.Errorf("Error BTree insert test: Expected Val = 5 for Root.Left node, got %d", testNode.Left.Val)
	}
	if testNode.Right.Val != 15 {
		t.Errorf("Error BTree insert test: Expected Val = 15 for Root.Right node, got %d", testNode.Right.Val)
	}
	if testNode.Left.Left.Val != 2 {
		t.Errorf("Error BTree insert test: Expected Val = 2 for Root.Left.Left node, got %d", testNode.Left.Left.Val)
	}
	if testNode.Left.Right.Val != 7 {
		t.Errorf("Error BTree insert test: Expected Val = 7 for Root.Left.Right node, got %d", testNode.Left.Right.Val)
	}
	if testNode.Right.Left.Val != 12 {
		t.Errorf("Error BTree insert test: Expected Val = 12 for Root.Right.Left node, got %d", testNode.Right.Left.Val)
	}
	if testNode.Right.Right.Val != 18 {
		t.Errorf("Error BTree insert test: Expected Val = 18 for Root.Right.Right node, got %d", testNode.Right.Right.Val)
	}
}

func TestBTreeWalk(t *testing.T) {
	testNode := algo.Node{Val: 10, Left: nil, Right: nil}
	testNode.InsertAll([]int{5, 15, 2, 7, 12, 18})
	walkString := testNode.Walk()
	if walkString != "[ 2 5 7 10 12 15 18 ]" {
		t.Errorf("BTree walk error: expected [ 2 5 7 10 12 15 18 ], got %s", walkString)
	}

	testNode = algo.Node{Val: 20, Left: nil, Right: nil}
	testNode.InsertAll([]int{48, 17, 33, 59, 26, 11, 61, 39, 35, 52, 8, 46, 37, 63, 28, 30, 24, 13, 55, 50, 19, 41, 7, 15, 1, 3, 22, 5, 44, 57})
	walkString = testNode.Walk()
	if walkString != "[ 1 3 5 7 8 11 13 15 17 19 20 22 24 26 28 30 33 35 37 39 41 44 46 48 50 52 55 57 59 61 63 ]" {
		t.Errorf("BTree walk error: expected [ 1 3 5 7 8 11 13 15 17 19 20 22 24 26 28 30 33 35 37 39 41 44 46 48 50 52 55 57 59 61 63 ], got %s", walkString)
	}
}

func TestBTreeDeleteNoChildren(t *testing.T) {
	testNode := algo.Node{Val: 10, Left: nil, Right: nil}
	testNode.InsertAll([]int{5, 15, 2, 7, 12, 18})

	testNode.Delete(2)
	if testNode.Left.Left != nil {
		t.Errorf("BTree delete error: deleted 2 from tree, root.Left.Left should be nil, was %d instead", testNode.Left.Left.Val)
	}

	testNode.Delete(18)
	if testNode.Right.Right != nil {
		t.Errorf("BTree delete error: deleted 18 from tree, root.Right.Right should be nil, was %d instead", testNode.Right.Right.Val)
	}

	testNode.Delete(12)
	if testNode.Right.Left != nil {
		t.Errorf("BTree delete error: deleted 12 from tree, root.Right.Left should be nil, was %d instead", testNode.Right.Left.Val)
	}
}

func TestBTreeDeleteOneChild(t *testing.T) {
	testNode := algo.Node{Val: 10, Left: nil, Right: nil}
	testNode.InsertAll([]int{5, 15, 7, 12})

	testNode.Delete(15)
	if testNode.Right.Left != nil {
		t.Errorf("BTree delete error: deleted 15 from tree, root.Right.Left should be nil, was %d instead", testNode.Right.Left.Val)
	}
	if testNode.Right.Val != 12 {
		t.Errorf("BTree delete error: deleted 15 from tree, root.Right should be 12, was %d instead", testNode.Right.Val)
	}

	testNode.Delete(5)
	if testNode.Left.Right != nil {
		t.Errorf("BTree delete error: deleted 5 from tree, root.Left.Right should be nil, was %d instead", testNode.Left.Right.Val)
	}
	if testNode.Left.Val != 7 {
		t.Errorf("BTree delete error: deleted 5 from tree, root.Left should be 7, was %d instead", testNode.Left.Val)
	}
}

func TestBTreeDeleteRoot(t *testing.T) {
	testNode := algo.Node{Val: 10, Left: nil, Right: nil}
	testNode.InsertAll([]int{5, 15, 2, 7, 12, 18})

	testNode.Delete(10)
	if testNode.Right.Left != nil {
		t.Errorf("BTree delete error: deleted 10 from tree, root.Right.Left should be nil, was %d instead", testNode.Right.Left.Val)
	}
	if testNode.Val != 12 {
		t.Errorf("BTree delete error: deleted 10 from tree, root should be 12, was %d instead", testNode.Val)
	}

	testNode.Delete(12)
	if testNode.Right.Right != nil {
		t.Errorf("BTree delete error: deleted 12 from tree, root.Right.Right should be nil, was %d instead", testNode.Right.Left.Val)
	}
	if testNode.Val != 15 {
		t.Errorf("BTree delete error: deleted 12 from tree, root should be 15, was %d instead", testNode.Val)
	}
}
