package challenge16

import (
	"fmt"
	"strings"
)

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func recursiveToTree(arr []int, i int) *TreeNode {
	if i >= len(arr) || arr[i] == -1 {
		return nil
	}
	return &TreeNode{
		Value: arr[i],
		Left:  recursiveToTree(arr, 2*i+1),
		Right: recursiveToTree(arr, 2*i+2),
	}
}

func transformTree(tree []int) *TreeNode {
	root := recursiveToTree(tree, 0)
	return root
}
func printTree(root *TreeNode, level int) {
	if root == nil {
		return
	}

	indentation := strings.Repeat("  ", level)

	fmt.Printf("{\n%svalue: %d\n", indentation, root.Value)
	fmt.Printf("%sleft: {\n", indentation)
	printTree(root.Left, level+1)
	fmt.Printf("%s}\n", indentation)
	fmt.Printf("%sright: {\n", indentation)
	printTree(root.Right, level+1)
	fmt.Printf("%s}\n", indentation)
}
func main() {
	// Ejemplo de uso
	treeArray := []int{3, 1, 0, 8, 12, -1, 1}
	treeObject := transformTree(treeArray)
	// Imprimir el árbol
	printTree(treeObject, 0)
}
