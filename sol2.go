package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	ch <- t.Value
	Walk(t.Left, ch)
	Walk(t.Right, ch)
}

func walkTree(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	close(ch)
}

func countValues(m map[int]int, ch chan int) {
	for value := range ch {
		if _, ok := m[value]; !ok {
			m[value] = 0
		}

		m[value] += 1
	}
}

func isSameValues(m map[int]int, ch chan int) bool {
	for value := range ch {
		if _, ok := m[value]; !ok {
			return false
		}
		m[value] -= 1
	}

	for _, value := range m {
		if value != 0 {
			return false
		}
	}

	return true
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	go walkTree(t1, ch1)

	ch2 := make(chan int)
	go walkTree(t2, ch2)

	m := make(map[int]int)
	countValues(m, ch1)

	return isSameValues(m, ch2)
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
