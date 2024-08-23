package main

import (
	"fmt"
	"reflect"
)

// TreeNode 定义二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// preorderTraversal 前序遍历二叉树的迭代实现
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var result []int
	stack := []*TreeNode{root} // 使用栈来模拟递归过程

	for len(stack) > 0 {
		node := stack[len(stack)-1]       // 取出栈顶节点
		stack = stack[:len(stack)-1]      // 弹出栈顶节点
		result = append(result, node.Val) // 访问当前节点

		// 先将右子树压栈
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		// 再将左子树压栈
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return result
}

// Test函数用于测试前序遍历
func Test() {
	// 构建二叉树
	//     1
	//      \
	//       2
	//      /
	//     3
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}

	// 期望的前序遍历结果
	expected := []int{1, 2, 3}

	// 执行前序遍历
	result := preorderTraversal(root)

	// 输出测试结果
	if reflect.DeepEqual(result, expected) {
		fmt.Println("Test passed!")
	} else {
		fmt.Printf("Test failed! Expected %v, got %v\n", expected, result)
	}
}

func main() {
	Test()
}
