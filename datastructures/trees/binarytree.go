package trees

import "errors"

type BinaryTreeNode struct {
	Value int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func NewBinaryTreeNode(value int) *BinaryTreeNode {
	return &BinaryTreeNode{
		Value: value,
		Left:  nil,
		Right: nil,
	}
}

func (t *BinaryTreeNode) AddChild(value int, right bool) error {
	if right && t.Right == nil {
		t.Right = NewBinaryTreeNode(value)
		return nil
	} else if right {
		return errors.New("right child already exists")
	}

	if t.Left == nil {
		t.Left = NewBinaryTreeNode(value)
		return nil
	} else {
		return errors.New("left child already exists")
	}
}

func (n *BinaryTreeNode) CountNodes() int {
	if n.Left != nil && n.Right != nil {
		return 1 + n.Left.CountNodes() + n.Right.CountNodes()
	}

	if n.Left != nil && n.Right == nil {
		return n.Left.CountNodes() + 1
	}

	if n.Left == nil && n.Right != nil {
		return 1 + n.Right.CountNodes()
	}

	return 1
}

func (n *BinaryTreeNode) GetLeaves() []*BinaryTreeNode {
	if n.Left != nil && n.Right != nil {
		return append(n.Left.GetLeaves(), n.Right.GetLeaves()...)
	}

	if n.Left != nil && n.Right == nil {
		return n.Left.GetLeaves()
	}

	if n.Left == nil && n.Right != nil {
		return n.Right.GetLeaves()
	}

	return []*BinaryTreeNode{n}
}
