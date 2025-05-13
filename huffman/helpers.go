package huffman

import "fmt"

func PrettyPrint(btree *Node, indent, prefix string) {
	if btree == nil {
		fmt.Printf("%s%s(nil)\n", indent, prefix)
		return
	}

	fmt.Printf("%s%sNode: %v\n", indent, prefix, btree.Value)
	PrettyPrint(btree.Left, indent+"    ", "L -> ")
	PrettyPrint(btree.Right, indent+"    ", "R -> ")
}

func (b Node) String() string {
	return fmt.Sprintf("\nL: {%v}\nR: {%v}\nV: {%v}\n", b.Left, b.Right, b.Value)
}

func isLeaf(node *Node) bool {
	return node.Left == nil && node.Right == nil && len(node.Value.Value) == 1
}
