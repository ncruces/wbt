package wbt

import "cmp"

// https://yoichihirai.com/bst.pdf

const (
	delta = 3
	gamma = 2
)

func (tree *Tree[K, V]) left_rebalance() *Tree[K, V] {
	if is_heavy(tree.Left(), tree.Right()) {
		if is_single(tree.left.Right(), tree.left.Left()) {
			frst := *tree.left
			tree.left = frst.right
			frst.right = tree.fixup()
			return frst.fixup()
		}
		frst := *tree.left
		scnd := *frst.right
		tree.left = scnd.right
		scnd.right = tree.fixup()
		frst.right = scnd.left
		scnd.left = frst.fixup()
		return scnd.fixup()
	}
	return tree.fixup()
}

func (tree *Tree[K, V]) right_rebalance() *Tree[K, V] {
	if is_heavy(tree.Right(), tree.Left()) {
		if is_single(tree.right.Left(), tree.right.Right()) {
			frst := *tree.right
			tree.right = frst.left
			frst.left = tree.fixup()
			return frst.fixup()
		}
		frst := *tree.right
		scnd := *frst.left
		tree.right = scnd.left
		scnd.left = tree.fixup()
		frst.left = scnd.right
		scnd.right = frst.fixup()
		return scnd.fixup()
	}
	return tree.fixup()
}

func is_heavy[K cmp.Ordered, V any](a, b *Tree[K, V]) bool {
	// Nodes are are at least 4 machine words,
	// so we'd run out of memory before this overflows.
	return (a.Len() + 1) > delta*(b.Len()+1)
}

func is_single[K cmp.Ordered, V any](a, b *Tree[K, V]) bool {
	// Nodes are are at least 4 machine words,
	// so we'd run out of memory before this overflows.
	return (a.Len() + 1) < gamma*(b.Len()+1)
}

func (tree *Tree[K, V]) fixup() *Tree[K, V] {
	tree.childs = tree.Left().Len() + tree.Right().Len()
	return tree
}
