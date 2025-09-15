package wbt

import (
	"maps"
	"testing"
)

func TestMakeSet_inc(t *testing.T) {
	tt := MakeSet(0, 1, 2, 3, 4, 5, 6)
	tt.check()

	if n := tt; n.key != 3 {
		t.Fatal(n.key)
	}
	if n := tt.left; n.key != 1 {
		t.Fatal(n.key)
	}
	if n := tt.right; n.key != 5 {
		t.Fatal(n.key)
	}
	if n := tt.left.left; n.key != 0 {
		t.Fatal(n.key)
	}
	if n := tt.left.right; n.key != 2 {
		t.Fatal(n.key)
	}
	if n := tt.right.left; n.key != 4 {
		t.Fatal(n.key)
	}
	if n := tt.right.right; n.key != 6 {
		t.Fatal(n.key)
	}
}

func TestMakeSet_dec(t *testing.T) {
	tt := MakeSet(6, 5, 5, 4, 3, 3, 2, 1, 1, 0)
	tt.check()

	if n := tt; n.key != 3 {
		t.Fatal(n.key)
	}
	if n := tt.left; n.key != 1 {
		t.Fatal(n.key)
	}
	if n := tt.right; n.key != 5 {
		t.Fatal(n.key)
	}
	if n := tt.left.left; n.key != 0 {
		t.Fatal(n.key)
	}
	if n := tt.left.right; n.key != 2 {
		t.Fatal(n.key)
	}
	if n := tt.right.left; n.key != 4 {
		t.Fatal(n.key)
	}
	if n := tt.right.right; n.key != 6 {
		t.Fatal(n.key)
	}
}

func TestMakeMap(t *testing.T) {
	tt := MakeMap(map[int]string{
		0: "zero",
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
		6: "six",
	})
	tt.check()

	if n := tt; n.key != 3 {
		t.Fatal(n.key)
	}
	if n := tt.left; n.key != 1 {
		t.Fatal(n.key)
	}
	if n := tt.right; n.key != 5 {
		t.Fatal(n.key)
	}
	if n := tt.left.left; n.key != 0 {
		t.Fatal(n.key)
	}
	if n := tt.left.right; n.key != 2 {
		t.Fatal(n.key)
	}
	if n := tt.right.left; n.key != 4 {
		t.Fatal(n.key)
	}
	if n := tt.right.right; n.key != 6 {
		t.Fatal(n.key)
	}
}

func TestTree_Collect(t *testing.T) {
	m1 := map[int]string{
		0: "zero",
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
		6: "six",
	}
	m2 := MakeMap(m1).Collect()

	if !maps.Equal(m1, m2) {
		t.Error(m1, m2)
	}
}
