package main

import (
	"fmt"
	"strconv"

	"github.com/zergon321/go-avltree"
)

func main() {
	tree := new(avltree.AVLTree[int, string])
	keys := []int{3, 2, 4, 1, 5}

	for _, key := range keys {
		tree.Add(key, strconv.Itoa(key))
	}

	tree.Remove(2)
	tree.Update(5, 6, strconv.Itoa(6))
	tree.DisplayInOrder()

	val := tree.Search(3).Value
	fmt.Println(val)
}
