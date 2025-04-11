package btree

import "fmt";

type BTree[T any] struct {
    Left *BTree[T]
    Right *BTree[T]
    Value T
}

func PrettyPrint[T any](btree *BTree[T], indent, prefix string) {
    if btree == nil {
        fmt.Printf("%s%s(nil)\n", indent, prefix)
        return
    }

    fmt.Printf("%s%sNode: %v\n", indent, prefix, btree.Value)
    PrettyPrint(btree.Left, indent+"  |", "L -> ")
    PrettyPrint(btree.Right, indent+"  |", "R -> ")


}

func (b BTree[T]) String() string {
    return fmt.Sprintf("\nL: {%v}\nR: {%v}\nV: {%v}\n", b.Left, b.Right, b.Value)
}


