package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBinaryTreeNode(t *testing.T) {
	tree := NewBinaryTreeNode(1)

	assert.Equal(t, 1, tree.Value)
}

func TestBinaryTreeAddChild(t *testing.T) {
	tree := NewBinaryTreeNode(1)
	tree.AddChild(2, false)
	tree.AddChild(3, true)
	tree.Right.AddChild(4, false)

	assert.Equal(t, 1, tree.Value)
	assert.Equal(t, 2, tree.Left.Value)
	assert.Equal(t, 3, tree.Right.Value)
	assert.Equal(t, 4, tree.Right.Left.Value)
	assert.Nil(t, tree.Right.Right)
}

func TestBinaryTreeCountNodes(t *testing.T) {
	tree := NewBinaryTreeNode(1)
	tree.AddChild(2, false)
	tree.AddChild(3, true)
	tree.Right.AddChild(4, false)

	assert.Equal(t, 4, tree.CountNodes())
}

func TestBinaryTreeGetLeaves(t *testing.T) {
	tree := NewBinaryTreeNode(1)
	tree.AddChild(2, false)
	tree.AddChild(3, true)
	tree.Right.AddChild(4, false)

	assert.Equal(t, 2, len(tree.GetLeaves()))

	tree.Right.AddChild(5, true)

	assert.Equal(t, 3, len(tree.GetLeaves()))
}
