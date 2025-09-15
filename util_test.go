package wbt

import "cmp"

func (tree *Tree[K, V]) check() int {
	if tree == nil {
		return 0
	}

	// BST invariants.
	if tree.Left() != nil && !cmp.Less(tree.Left().Key(), tree.Key()) {
		panic("the left child's key must be less than this key")
	}
	if tree.Right() != nil && !cmp.Less(tree.Key(), tree.Right().Key()) {
		panic("this key must be less than the right child's key")
	}

	// WBT tree invariants.
	if is_heavy(tree.Left(), tree.Right()) || is_heavy(tree.Right(), tree.Left()) {
		panic("the tree is unbalanced")
	}

	// OST invariant.
	len := 1 + tree.Left().check() + tree.Right().check()
	if len != tree.Len() {
		panic("the length of this tree is one plus the length of both children")
	}
	return len
}
