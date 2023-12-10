package combinator

type TreeNode struct {
	Left   *TreeNode
	Right  *TreeNode
	Parent *TreeNode
	IsLeaf bool
	IsRoot bool
	Leaf   string
}

func getRoot(tree *TreeNode) *TreeNode {
	if tree.IsRoot {
		return tree
	}
	return getRoot(tree.Parent)
}

func getLeftMostLeaf(tree *TreeNode) *TreeNode {
	if tree.IsLeaf {
		return tree
	}
	return getLeftMostLeaf(tree.Left)
}

func getNthParent(tree *TreeNode, n int) *TreeNode {
	if tree.IsRoot || n == 0 {
		return tree
	}
	return getNthParent(tree.Parent, n-1)
}

func numNodesToRoot(descendent *TreeNode, root *TreeNode) int {
	if descendent == root {
		return 0
	}
	return 1 + numNodesToRoot(descendent.Parent, root)
}

func getNRightSiblings(descendent *TreeNode, n int) []*TreeNode {
	siblings := []*TreeNode{}
	current := descendent
	for i := 0; i < n; i++ {
		current = current.Parent
		siblings = append(siblings, current.Right)
	}
	return siblings
}

func copy(root *TreeNode) *TreeNode {
	if root.IsLeaf {
		return &TreeNode{
			IsRoot: root.IsRoot,
			Parent: root.Parent,
			IsLeaf: root.IsLeaf,
			Leaf:   root.Leaf,
		}
	}

	node := &TreeNode{
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
