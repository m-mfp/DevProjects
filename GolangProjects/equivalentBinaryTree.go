package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same determines whether the trees t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	seq1 := []int{}
	seq2 := []int{}

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := 0; i < 10; i++ {
		seq1 = append(seq1, <-ch1)
		seq2 = append(seq2, <-ch2)
	}

	if len(seq1) != len(seq2) {
		return false
	}
	for i := range seq1 {
		if seq1[i] != seq2[i] {
			return false
		}
	}
	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	a := Same(tree.New(1), tree.New(1))
	b := Same(tree.New(1), tree.New(2))

	fmt.Println(a)
	fmt.Println(b)
}
