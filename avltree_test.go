package avltree_test

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/zergon321/go-avltree"
	"github.com/zergon321/rb"
)

func BenchmarkInsert(b *testing.B) {
	tree := &avltree.AVLTree[int, int]{}

	for i := 0; i < b.N; i++ {
		value := rand.Int()
		tree.Add(value, value)
	}
}

func BenchmarkRBInsert(b *testing.B) {
	tree := rb.NewTree[int, int]()

	for i := 0; i < b.N; i++ {
		value := rand.Int()
		tree.Insert(value, value)
	}
}

func BenchmarkSliceInsert(b *testing.B) {
	slice := []int{}

	for i := 0; i < b.N; i++ {
		value := rand.Int()
		length := len(slice)

		ind := sort.Search(length, func(i int) bool {
			return slice[i] >= value
		})

		if ind == 0 {
			slice = append(slice, 0)
			copy(slice[1:], slice)
			slice[0] = value
		} else if ind < length {
			slice = append(slice[:ind+1], slice[ind:]...)
			slice[ind] = value
		} else {
			slice = append(slice, value)
		}
	}
}
