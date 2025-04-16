package hnode

import "fmt"

func PrettyPrint(btree *HNode, indent, prefix string) {
	if btree == nil {
		fmt.Printf("%s%s(nil)\n", indent, prefix)
		return
	}

	fmt.Printf("%s%sNode: %v\n", indent, prefix, btree.Value)
	PrettyPrint(btree.Left, indent+"    ", "L -> ")
	PrettyPrint(btree.Right, indent+"    ", "R -> ")
}

func (b HNode) String() string {
	return fmt.Sprintf("\nL: {%v}\nR: {%v}\nV: {%v}\n", b.Left, b.Right, b.Value)
}
