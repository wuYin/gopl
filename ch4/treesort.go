package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type tree struct {
	value       int
	left, right *tree
}

func main() {
	nums := make([]int, 20)
	for i := range nums {
		nums[i] = rand.Int() % 20
	}
	sortTree(nums)
	if !sort.IntsAreSorted(nums) {
		fmt.Printf("not sorted: %v\n", nums)
	}
	fmt.Println(nums)
}

func sortTree(vals []int) {
	var root *tree
	for _, v := range vals {
		root = add(root, v)
	}
	appendValues(vals[:0], root)
}

// 将元素按顺序追加到 values 中，返回结果 slice
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

// 递归将元素放到树中合适位置处
func add(t *tree, val int) *tree {
	if t == nil {
		return &tree{value: val}
	}
	if val < t.value {
		t.left = add(t.left, val)
	} else {
		t.right = add(t.right, val)
	}
	return t
}
