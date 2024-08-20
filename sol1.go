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

func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	go func() {
		Walk(t1, ch1)
		close(ch1)
	}()

	ch2 := make(chan int)
	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()

	m := make(map[int]int)
	for value := range ch1 {
		if _, ok := m[value]; !ok {
			m[value] = 0
		}

		m[value] += 1
	}

	for value := range ch2 {
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

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
