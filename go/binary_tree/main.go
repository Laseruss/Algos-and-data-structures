package main

import "fmt"

func main() {
	root := TreeNode{
		Val:   10,
		Left:  nil,
		Right: nil,
	}

	nodes := []int{1, 4, 8, 9, 76, 54, 0, 12}

	for _, v := range nodes {
		root.insert(v)
	}

	fmt.Println("Preorder")
	fmt.Println("------------")
	fmt.Println(root.preOrderTraversal())
	fmt.Println()
	fmt.Println("Inorder")
	fmt.Println("------------")
	fmt.Println(root.inOrderTraversal())
	fmt.Println()
	fmt.Println("Postorder")
	fmt.Println("------------")
	fmt.Println(root.postOrderTraversal())
	fmt.Println()

	fmt.Printf("Min value in tree: %d\n", root.min())
	fmt.Printf("Max value in tree: %d\n", root.max())
}
