package wbt

import "cmp"

// https://yoichihirai.com/bst.pdf

const (
	delta_p, delta_q = 5, 2
	gamma_p, gamma_q = 3, 2
)

func (tree *Tree[K, V]) left_rebalance() *Tree[K, V] {
	if !is_heavy(tree.Left(), tree.Right()) {
		return tree.fixup()
	}
	frst := *tree.left
	if is_single(frst.Right(), frst.Left()) {
		tree.left = frst.right
		frst.right = tree.fixup()
		return frst.fixup()
	}
	scnd := *frst.right
	tree.left = scnd.right
	scnd.right = tree.fixup()
	frst.right = scnd.left
	scnd.left = frst.fixup()
	return scnd.fixup()
}

func (tree *Tree[K, V]) right_rebalance() *Tree[K, V] {
	if !is_heavy(tree.Right(), tree.Left()) {
		return tree.fixup()
	}
	frst := *tree.right
	if is_single(frst.Left(), frst.Right()) {
		tree.right = frst.left
		frst.left = tree.fixup()
		return frst.fixup()
	}
	scnd := *frst.left
	tree.right = scnd.left
	scnd.left = tree.fixup()
	frst.left = scnd.right
	scnd.right = frst.fixup()
	return scnd.fixup()
}

func is_heavy[K cmp.Ordered, V any](a, b *Tree[K, V]) bool {
	// Nodes are are at least 4 machine words,
	// so we'd run out of memory before this overflows.
	return delta_q*(a.Len()+1) > delta_p*(b.Len()+1)
}

func is_single[K cmp.Ordered, V any](a, b *Tree[K, V]) bool {
	// Nodes are are at least 4 machine words,
	// so we'd run out of memory before this overflows.
	return gamma_q*(a.Len()+1) < gamma_p*(b.Len()+1)
}

func (tree *Tree[K, V]) fixup() *Tree[K, V] {
	tree.childs = tree.Left().Len() + tree.Right().Len()
	return tree
}
