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

func walkTree(t *tree.Tree) chan int {
	ch := make(chan int)
	go func() {
		Walk(t, ch)
		close(ch)
	}()
	return ch
}

func countValues(ch chan int) map[int]int {
	m := make(map[int]int)
	for value := range ch {
		if _, ok := m[value]; !ok {
			m[value] = 0
		}

		m[value] += 1
	}

	return m
}

func IsEqualMaps(m1 map[int]int, m2 map[int]int) bool {
	if len(m1) != len(m2) {
		return false
	}

	for key, value1 := range m1 {
		value2, ok := m2[key]

		if !ok {
			return false
		}
		if value1 != value2 {
			return false
		}
	}

	return true
}

func Same(t1, t2 *tree.Tree) bool {
	ch1 := walkTree(t1)
	ch2 := walkTree(t2)

	m1 := countValues(ch1)
	m2 := countValues(ch2)

	return IsEqualMaps(m1, m2)
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
