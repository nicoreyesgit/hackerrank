package datastructures

type Tree struct {
	root *NodeTree
}

type NodeTree struct {
	value int
	left  *NodeTree
	right *NodeTree
}

func CompareBinaryTrees(a *NodeTree, b *NodeTree) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}

	if a.value != b.value {
		return false
	}

	return CompareBinaryTrees(a.left, b.left) && CompareBinaryTrees(a.right, b.right)
}
