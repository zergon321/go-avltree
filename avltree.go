package avltree

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// AVLTree[TKey constraints.Ordered, TValue any] structure. Public methods are Add, Remove, Update, Search, DisplayTreeInOrder.
type AVLTree[TKey constraints.Ordered, TValue any] struct {
	root *AVLNode[TKey, TValue]
}

func (t *AVLTree[TKey, TValue]) Add(key TKey, value TValue) {
	t.root = t.root.add(key, value)
}

func (t *AVLTree[TKey, TValue]) Remove(key TKey) {
	t.root = t.root.remove(key)
}

func (t *AVLTree[TKey, TValue]) Update(oldKey TKey, newKey TKey, newValue TValue) {
	t.root = t.root.remove(oldKey)
	t.root = t.root.add(newKey, newValue)
}

func (t *AVLTree[TKey, TValue]) Search(key TKey) (node *AVLNode[TKey, TValue]) {
	return t.root.search(key)
}

func (t *AVLTree[TKey, TValue]) DisplayInOrder() {
	t.root.displayNodesInOrder()
}

// AVLNode structure
type AVLNode[TKey constraints.Ordered, TValue any] struct {
	key   TKey
	Value TValue

	// height counts nodes (not edges)
	height int
	left   *AVLNode[TKey, TValue]
	right  *AVLNode[TKey, TValue]
}

// Adds a new node
func (n *AVLNode[TKey, TValue]) add(key TKey, value TValue) *AVLNode[TKey, TValue] {
	if n == nil {
		return &AVLNode[TKey, TValue]{key, value, 1, nil, nil}
	}

	if key < n.key {
		n.left = n.left.add(key, value)
	} else if key > n.key {
		n.right = n.right.add(key, value)
	} else {
		// if same key exists update value
		n.Value = value
	}
	return n.rebalanceTree()
}

// Removes a node
func (n *AVLNode[TKey, TValue]) remove(key TKey) *AVLNode[TKey, TValue] {
	if n == nil {
		return nil
	}
	if key < n.key {
		n.left = n.left.remove(key)
	} else if key > n.key {
		n.right = n.right.remove(key)
	} else {
		if n.left != nil && n.right != nil {
			// node to delete found with both children;
			// replace values with smallest node of the right sub-tree
			rightMinNode := n.right.findSmallest()
			n.key = rightMinNode.key
			n.Value = rightMinNode.Value
			// delete smallest node that we replaced
			n.right = n.right.remove(rightMinNode.key)
		} else if n.left != nil {
			// node only has left child
			n = n.left
		} else if n.right != nil {
			// node only has right child
			n = n.right
		} else {
			// node has no children
			n = nil
			return n
		}

	}
	return n.rebalanceTree()
}

// Searches for a node
func (n *AVLNode[TKey, TValue]) search(key TKey) *AVLNode[TKey, TValue] {
	if n == nil {
		return nil
	}
	if key < n.key {
		return n.left.search(key)
	} else if key > n.key {
		return n.right.search(key)
	} else {
		return n
	}
}

// Displays nodes left-depth first (used for debugging)
func (n *AVLNode[TKey, TValue]) displayNodesInOrder() {
	if n.left != nil {
		n.left.displayNodesInOrder()
	}
	fmt.Print(n.key, " ")
	if n.right != nil {
		n.right.displayNodesInOrder()
	}
}

func (n *AVLNode[TKey, TValue]) getHeight() int {
	if n == nil {
		return 0
	}
	return n.height
}

func (n *AVLNode[TKey, TValue]) recalculateHeight() {
	n.height = 1 + max(n.left.getHeight(), n.right.getHeight())
}

// Checks if node is balanced and rebalance
func (n *AVLNode[TKey, TValue]) rebalanceTree() *AVLNode[TKey, TValue] {
	if n == nil {
		return n
	}
	n.recalculateHeight()

	// check balance factor and rotateLeft if right-heavy and rotateRight if left-heavy
	balanceFactor := n.left.getHeight() - n.right.getHeight()
	if balanceFactor == -2 {
		// check if child is left-heavy and rotateRight first
		if n.right.left.getHeight() > n.right.right.getHeight() {
			n.right = n.right.rotateRight()
		}
		return n.rotateLeft()
	} else if balanceFactor == 2 {
		// check if child is right-heavy and rotateLeft first
		if n.left.right.getHeight() > n.left.left.getHeight() {
			n.left = n.left.rotateLeft()
		}
		return n.rotateRight()
	}
	return n
}

// Rotate nodes left to balance node
func (n *AVLNode[TKey, TValue]) rotateLeft() *AVLNode[TKey, TValue] {
	newRoot := n.right
	n.right = newRoot.left
	newRoot.left = n

	n.recalculateHeight()
	newRoot.recalculateHeight()
	return newRoot
}

// Rotate nodes right to balance node
func (n *AVLNode[TKey, TValue]) rotateRight() *AVLNode[TKey, TValue] {
	newRoot := n.left
	n.left = newRoot.right
	newRoot.right = n

	n.recalculateHeight()
	newRoot.recalculateHeight()
	return newRoot
}

// Finds the smallest child (based on the key) for the current node
func (n *AVLNode[TKey, TValue]) findSmallest() *AVLNode[TKey, TValue] {
	if n.left != nil {
		return n.left.findSmallest()
	} else {
		return n
	}
}

// Returns max number - TODO: std lib seemed to only have a method for floats!
func max[TKey constraints.Ordered](a TKey, b TKey) TKey {
	if a > b {
		return a
	}
	return b
}
