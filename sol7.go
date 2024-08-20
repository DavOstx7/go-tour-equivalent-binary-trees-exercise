package main

import (
	"fmt"
	"golang.org/x/tour/tree"
	"reflect"
)

func walkRecursive(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		walkRecursive(t.Left, ch)
	}

	ch <- t.Value

	if t.Right != nil {
		walkRecursive(t.Right, ch)
	}
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walkRecursive(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	var tree1, tree2 []int
	var v1, v2 int
	var ok1, ok2 bool = true, true

	for ok1 || ok2 {
		select {
		case v1, ok1 = <-ch1:
			if ok1 {
				tree1 = append(tree1, v1)
			}
		case v2, ok2 = <-ch2:
			if ok2 {
				tree2 = append(tree2, v2)
			}
		}
	}
	return reflect.DeepEqual(tree1, tree2)
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
