package avltree_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zergon321/go-avltree"
)

type Range struct {
	A float32
	B float32
}

func (ls Range) Less(other avltree.Comparable) bool {
	if point, ok := other.(Point); ok {
		return point.Num < ls.A
	}

	otherLS := other.(Range)
	return ls.B < otherLS.A
}

func (ls Range) Greater(other avltree.Comparable) bool {
	if point, ok := other.(Point); ok {
		return point.Num > ls.B
	}

	otherLS := other.(Range)
	return ls.A > otherLS.B
}

func (ls Range) Equal(other avltree.Comparable) bool {
	if point, ok := other.(Point); ok {
		return point.Num >= ls.A && point.Num <= ls.B
	}

	otherLS := other.(Range)
	return ls.A == otherLS.A && ls.B == otherLS.B
}

type Point struct {
	Num float32
}

func (p Point) Less(other avltree.Comparable) bool {
	if r, ok := other.(Range); ok {
		return p.Num < r.A
	}

	otherLS := other.(Point)
	return p.Num < otherLS.Num
}

func (p Point) Greater(other avltree.Comparable) bool {
	if r, ok := other.(Range); ok {
		return p.Num > r.B
	}

	otherLS := other.(Point)
	return p.Num > otherLS.Num
}

func (p Point) Equal(other avltree.Comparable) bool {
	if r, ok := other.(Range); ok {
		return p.Num >= r.A && p.Num <= r.B
	}

	otherLS := other.(Point)
	return p.Num == otherLS.Num
}

type Geometric interface {
	avltree.Comparable
}

func TestFill(t *testing.T) {
	tree, err := avltree.NewUnrestrictedAVLTree[Range, struct{}]()
	assert.Nil(t, err)

	for i := 0; i < 10; i += 2 {
		r := Range{A: float32(i), B: float32(i + 1)}
		tree.Add(r, struct{}{})
	}

	str := ""

	err = tree.VisitInOrder(func(node *avltree.UnrestrictedAVLNode[Range, struct{}]) error {
		str += strconv.Itoa(int(node.Key().A))
		str += strconv.Itoa(int(node.Key().B))

		return nil
	})

	assert.Nil(t, err)
	assert.Equal(t, "0123456789", str)
}

func TestRangeTreeSearch(t *testing.T) {
	tree, err := avltree.NewUnrestrictedAVLTree[Geometric, struct{}]()
	assert.Nil(t, err)

	for i := 0; i < 10; i += 2 {
		r := Range{A: float32(i), B: float32(i + 1)}
		tree.Add(r, struct{}{})
	}

	node := tree.Search(Point{Num: 0.5})
	assert.NotNil(t, node)
	assert.Equal(t, Range{A: 0, B: 1}, node.Key())

	node = tree.Search(Point{Num: -1})
	assert.Nil(t, node)

	node = tree.Search(Point{Num: 10})
	assert.Nil(t, node)

	node = tree.Search(Point{Num: 1.5})
	assert.Nil(t, node)
}
