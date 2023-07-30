package avltree

type Comparable interface {
	Less(Comparable) bool
	Greater(Comparable) bool
	Equal(Comparable) bool
}
