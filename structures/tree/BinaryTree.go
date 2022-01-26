package tree

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

type BinaryTree struct {
	Root *TreeNode
}

func NewEmptyBST() *BinaryTree {
	return &BinaryTree{nil}
}

func NewBST(root *TreeNode) *BinaryTree {
	return &BinaryTree{root}
}

func (t *BinaryTree) Insert(data int) {
	if t.Root == nil {
		t.Root = &TreeNode{data, nil, nil}
	} else {
		t.Root.Insert(data)
	}
}

func (node *TreeNode) Insert(data int) {
	if data > node.Val {
		if node.Right == nil {
			node.Right = &TreeNode{data, nil, nil}
		} else {
			node.Right.Insert(data)
		}
	} else {
		if node.Left == nil {
			node.Left = &TreeNode{data, nil, nil}
		} else {
			node.Left.Insert(data)
		}
	}
}

func Ints2Tree(ints []int) *BinaryTree {
	tree := NewEmptyBST()
	for _, val := range ints {
		tree.Insert(val)
	}
	return tree
}

func (t *BinaryTree) GetRoot() *TreeNode {
	return t.Root
}

func Tree2Ints(tree *BinaryTree) []int {
	posMap := map[int]int{}
	maxIndex := 0
	TraversalAsHeap(tree.Root, 1, posMap, &maxIndex)
	output := make([]int, maxIndex+1)
	for i, val := range posMap {
		output[i] = val
	}
	return output[1:] // nil Node will be '0'
}

func TraversalAsHeap(node *TreeNode, i int, posMap map[int]int, maxIndex *int) {
	if node == nil {
		return
	}
	TraversalAsHeap(node.Left, 2*i, posMap, maxIndex)
	TraversalAsHeap(node.Right, 2*i+1, posMap, maxIndex)
	posMap[i] = node.Val
	if i > *maxIndex {
		*maxIndex = i
	}
}

func Tree2InOrder(tree *BinaryTree) []int {
	return TraversalInOrder(tree.Root)
}

func TraversalInOrder(node *TreeNode) []int {
	output := []int{}
	if node == nil {
		return output
	}
	output = append(output, TraversalInOrder(node.Left)...)
	output = append(output, node.Val)
	output = append(output, TraversalInOrder(node.Right)...)
	return output
}

func Tree2PreOrder(tree *BinaryTree) []int {
	return TraversalPreOrder(tree.Root)
}

func TraversalPreOrder(node *TreeNode) []int {
	output := []int{}
	if node == nil {
		return output
	}
	output = append(output, node.Val)
	output = append(output, TraversalPreOrder(node.Left)...)
	output = append(output, TraversalPreOrder(node.Right)...)
	return output
}

// func GetTargetNode(root *TreeNode, target int) *TreeNode {

// }
