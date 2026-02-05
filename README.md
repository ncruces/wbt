# Immutable Weight-Balanced Trees

[![Go Reference](https://pkg.go.dev/badge/image)](https://pkg.go.dev/github.com/ncruces/wbt)
[![Go Report](https://goreportcard.com/badge/github.com/ncruces/wbt)](https://goreportcard.com/report/github.com/ncruces/wbt)
[![Go Coverage](https://github.com/ncruces/wbt/wiki/coverage.svg)](https://raw.githack.com/wiki/ncruces/wbt/coverage.html)

A [persistent](https://en.wikipedia.org/wiki/Persistent_data_structure) implementation of
[weight-balanced trees](https://en.wikipedia.org/wiki/Weight-balanced_tree),
including the [join-based tree algorithms](https://en.wikipedia.org/wiki/Join-based_tree_algorithms)
and [order statistics](https://en.wikipedia.org/wiki/Order_statistic_tree).

This package will perform better than a [B-tree](https://en.wikipedia.org/wiki/B-tree) [^1]
(in terms of CPU _and_ memory) **if** you make use of persistence.

Otherwise, i.e. if copying/cloning your B-tree is infrequent
(because either the tree is immutable after construction,
 or hundreds of modifications can be batched in a single ‘transaction’)
B-trees will perform better.

[^1]: compared with both
 [`github.com/tidwall/btree`](https://github.com/tidwall/btree) and
 [`github.com/google/btree`](https://github.com/google/btree).
