package tree

import "testing"

func TestTree(t *testing.T) {
	bst := NewTree()
	bst.Root = NewTreeNode(10)
	bst.Root.Left = NewTreeNode(5)
	bst.Root.Left.Left = NewTreeNode(3)
	bst.Root.Left.Right = NewTreeNode(7)
	bst.Root.Right = NewTreeNode(15)
	bst.Root.Right.Right = NewTreeNode(18)
}

/*
      10
    /    \
   5      15
  / \       \
 3   7      18
*/

func TestDFS(t *testing.T) {
	bst := NewTree()
	bst.Root = NewTreeNode(10)
	bst.Root.Left = NewTreeNode(5)
	bst.Root.Left.Left = NewTreeNode(3)
	bst.Root.Left.Right = NewTreeNode(7)
	bst.Root.Right = NewTreeNode(15)
	bst.Root.Right.Right = NewTreeNode(18)
}
