package combinator

// A simple binary tree
type treeNode struct {
	Left   *treeNode
	Right  *treeNode
	Parent *treeNode
	IsLeaf bool
	IsRoot bool
	Leaf   string // If `IsLeaf`, the leaf's content
}

// Returns the tree's root given any treeNode
func getRoot(tree *treeNode) *treeNode {
	if tree.IsRoot {
		return tree
	}
	return getRoot(tree.Parent)
}

// Returns the left-most leaf from the current node
func getLeftMostLeaf(tree *treeNode) *treeNode {
	if tree.IsLeaf {
		return tree
	}
	return getLeftMostLeaf(tree.Left)
}

// Return's the n'th parent from the current node
func getNthParent(tree *treeNode, n int) *treeNode {
	if tree.IsRoot || n == 0 {
		return tree
	}
	return getNthParent(tree.Parent, n-1)
}

// Returns the number of nodes to get to the root from the current node
func numNodesToRoot(descendent *treeNode, root *treeNode) int {
	if descendent == root {
		return 0
	}
	return 1 + numNodesToRoot(descendent.Parent, root)
}

// Returns `n` "Right Siblings" (this is what I am calling them, but I'm sure
// there must be some correct technical name for them). The first "Right Sibling"
// of a node is it's parent's right child. The second "Right Sibling" is the
// node's grandparent's right child, and so on...
func getNRightSiblings(descendent *treeNode, n int) []*treeNode {
	siblings := []*treeNode{}
	current := descendent
	for i := 0; i < n; i++ {
		current = current.Parent
		siblings = append(siblings, current.Right)
	}
	return siblings
}

// Recursively copies an entire subtree
func copy(root *treeNode) *treeNode {
	if root.IsLeaf {
		return &treeNode{
			IsRoot: root.IsRoot,
			Parent: root.Parent,
			IsLeaf: root.IsLeaf,
			Leaf:   root.Leaf,
		}
	}

	node := &treeNode{
		IsRoot: root.IsRoot,
		Parent: root.Parent,
		Left:   copy(root.Left),
		Right:  copy(root.Right),
		IsLeaf: root.IsLeaf,
		Leaf:   root.Leaf,
	}

	node.Left.Parent = node
	node.Right.Parent = node
	return node
}
