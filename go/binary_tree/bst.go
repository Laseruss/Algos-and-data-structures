package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) search(x int) *TreeNode {
	if t == nil || t.Val == x {
		return t
	}

	if t.Val > x {
		return t.Left.search(x)
	} else {
		return t.Right.search(x)
	}
}

func (t *TreeNode) insert(x int) {
	if t.Val < x {
		if t.Right == nil {
			t.Right = &TreeNode{x, nil, nil}
			return
		} else {
			t.Right.insert(x)
		}
	}

	if t.Val > x {
		if t.Left == nil {
			t.Left = &TreeNode{x, nil, nil}
			return
		} else {
			t.Left.insert(x)
		}
	}
}

func delete(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return node
	}

	if node.Val > val {
		node.Left = delete(node.Left, val)
		return node
	} else if node.Val < val {
		node.Right = delete(node.Right, val)
		return node
	} else if node.Val == val {
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		} else {
			node.Right = lift(node.Right, node)
			return node
		}
	}

	return node
}

func lift(node, nodeToDelete *TreeNode) *TreeNode {
	if node.Left != nil {
		node.Left = lift(node.Left, nodeToDelete)
		return node
	} else {
		nodeToDelete.Val = node.Val
		return node.Right
	}

}

func (t *TreeNode) preOrderTraversal() []int {
	l := []int{}
	preWalk(t, &l)

	return l
}

func preWalk(node *TreeNode, nums *[]int) {
	if node == nil {
		return
	}

	*nums = append(*nums, node.Val)
	preWalk(node.Left, nums)
	preWalk(node.Right, nums)

	return
}

func (t *TreeNode) inOrderTraversal() []int {
	return inOrderWalk(t, &[]int{})
}

func inOrderWalk(node *TreeNode, nums *[]int) []int {
	if node == nil {
		return *nums
	}

	inOrderWalk(node.Left, nums)
	*nums = append(*nums, node.Val)
	inOrderWalk(node.Right, nums)

	return *nums
}

func (t *TreeNode) postOrderTraversal() []int {
	return postOrderWalk(t, &[]int{})
}

func postOrderWalk(node *TreeNode, nums *[]int) []int {
	if node == nil {
		return *nums
	}

	postOrderWalk(node.Left, nums)
	postOrderWalk(node.Right, nums)
	*nums = append(*nums, node.Val)

	return *nums
}

func (t *TreeNode) min() int {
	if t.Left == nil {
		return t.Val
	}

	return t.Left.min()
}

func (t *TreeNode) max() int {
	if t.Right == nil {
		return t.Val
	}

	return t.Right.max()
}
