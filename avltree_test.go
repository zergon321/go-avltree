package avltree_test

import (
	"math/rand"
	"sort"
	"strconv"
	"testing"

	godsavl "github.com/emirpasic/gods/trees/avltree"
	godsrb "github.com/emirpasic/gods/trees/redblacktree"
	"github.com/stretchr/testify/assert"
	"github.com/zergon321/go-avltree"
	"github.com/zergon321/mempool"
	"github.com/zergon321/rb"
)

func TestVisitInOrder(t *testing.T) {
	tree, err := avltree.NewAVLTree[int, int]()
	assert.Nil(t, err)

	tree.Add(1, 1)
	tree.Add(2, 2)
	tree.Add(3, 3)
	tree.Add(4, 4)
	tree.Add(5, 5)
	tree.Add(6, 6)
	tree.Add(7, 7)
	tree.Add(8, 8)

	str := ""

	err = tree.VisitInOrder(func(node *avltree.AVLNode[int, int]) error {
		str += strconv.Itoa(node.Value)
		return nil
	})

	assert.Nil(t, err)
	assert.Equal(t, "12345678", str)
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

func BenchmarkGodsAVLInsert(b *testing.B) {
	tree := godsavl.NewWithIntComparator()

	for i := 0; i < b.N; i++ {
		value := rand.Int()
		tree.Put(value, value)
	}
}

func BenchmarkGodsRBInsert(b *testing.B) {
	tree := godsrb.NewWithIntComparator()

	for i := 0; i < b.N; i++ {
		value := rand.Int()
		tree.Put(value, value)
	}
}

func BenchmarkAVLInsert(b *testing.B) {
	tree := &avltree.AVLTree[int, int]{}

	for i := 0; i < b.N; i++ {
		value := rand.Int()
		tree.Add(value, value)
	}
}

func BenchmarkAVLInsertThenRemove(b *testing.B) {
	tree := &avltree.AVLTree[int, int]{}

	for i := 0; i < b.N; i++ {
		value := rand.Int()
		tree.Add(value, value)
		tree.Remove(value)
	}
}

func BenchmarkAVLInsertThenRemoveMemoryPool(b *testing.B) {
	pool, _ := mempool.NewPool(func() *avltree.AVLNode[int, int] {
		return &avltree.AVLNode[int, int]{}
	}, mempool.PoolOptionInitialLength[*avltree.AVLNode[int, int]](1))
	tree, _ := avltree.NewAVLTree(avltree.AVLTreeOptionWithMemoryPool(pool))

	for i := 0; i < b.N; i++ {
		value := rand.Int()
		tree.Add(value, value)
		tree.Remove(value)
	}
}
