package avltree

import (
	"github.com/zergon321/mempool"
	"golang.org/x/exp/constraints"
)

type AVLTreeOption[
	TKey constraints.Ordered, TValue any,
] func(tree *AVLTree[TKey, TValue]) error

func AVLTreeOptionWithMemoryPool[
	TKey constraints.Ordered, TValue any,
](
	pool *mempool.Pool[*AVLNode[TKey, TValue]],
) AVLTreeOption[TKey, TValue] {
	return func(tree *AVLTree[TKey, TValue]) error {
		tree.pool = pool
		return nil
	}
}

type UnrestrictedAVLTreeOption[
	TKey Comparable, TValue any,
] func(tree *UnrestrictedAVLTree[TKey, TValue]) error

func UnrestrictedAVLTreeOptionWithMemoryPool[
	TKey Comparable, TValue any,
](
	pool *mempool.Pool[*UnrestrictedAVLNode[TKey, TValue]],
) UnrestrictedAVLTreeOption[TKey, TValue] {
	return func(tree *UnrestrictedAVLTree[TKey, TValue]) error {
		tree.pool = pool
		return nil
	}
}
