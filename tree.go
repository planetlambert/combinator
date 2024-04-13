package combinator

// A simple binary tree
type Tree struct {
	Left   *Tree
	Right  *Tree
	Parent *Tree
	IsLeaf bool
	IsRoot bool
	Leaf   string // If `IsLeaf`, the leaf's content
}

// Returns the tree's root given any tree
func getRoot(tree *Tree) *Tree {
	if tree.IsRoot {
		return tree
	}
	return getRoot(tree.Parent)
}

// Returns the left-most leaf from the current node
func getLeftMostLeaf(tree *Tree) *Tree {
	if tree.IsLeaf {
		return tree
	}
	return getLeftMostLeaf(tree.Left)
}

// Return's the n'th parent from the current node
func getNthParent(tree *Tree, n int) *Tree {
	if tree.IsRoot || n == 0 {
		return tree
	}
	return getNthParent(tree.Parent, n-1)
}

// Returns the number of nodes to get to the root from the current node
func numNodesToRoot(descendent *Tree, root *Tree) int {
	if descendent == root {
		return 0
	}
	return 1 + numNodesToRoot(descendent.Parent, root)
}

// Returns `n` "Right Siblings" (this is what I am calling them, but I'm sure
// there must be some correct technical name for them). The first "Right Sibling"
// of a node is it's parent's right child. The second "Right Sibling" is the
// node's grandparent's right child, and so on...
func getNRightSiblings(descendent *Tree, n int) []*Tree {
	siblings := []*Tree{}
	current := descendent
	for i := 0; i < n; i++ {
		current = current.Parent
		siblings = append(siblings, current.Right)
	}
	return siblings
}

// Recursively copies an entire subtree
func copy(root *Tree) *Tree {
	if root.IsLeaf {
		return &Tree{
			IsRoot: root.IsRoot,
			Parent: root.Parent,
			IsLeaf: root.IsLeaf,
			Leaf:   root.Leaf,
		}
	}

	node := &Tree{
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

func join(trees []*Tree) *Tree {
	if len(trees) == 0 {
		return nil
	}

	if len(trees) == 1 {
		return copy(trees[0])
	}

	firstRootCopy := copy(trees[0])
	secondRootCopy := copy(trees[1])
	firstRootCopy.IsRoot = false
	secondRootCopy.IsRoot = false

	newSlice := []*Tree{
		{
			Left:   firstRootCopy,
			Right:  secondRootCopy,
			IsRoot: true,
		},
	}
	newSlice = append(newSlice, trees[2:]...)
	return join(newSlice)
}
