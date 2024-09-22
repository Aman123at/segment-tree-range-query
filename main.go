package main

import (
	"log"
	"math"
)

type SegmentTree struct {
	tree   []int
	n      int
	height int
}

// creates new segment tree
func NewSegmentTree(arr []int) *SegmentTree {
	n := len(arr)                                       // length of array
	height := int(math.Ceil(math.Log2(float64(n))))     // calculating height of tree
	treeSize := 2*int(math.Pow(2, float64(height))) - 1 // calculating overall size of tree => 2*2^h - 1

	// segment tree instance
	st := &SegmentTree{
		tree:   make([]int, treeSize), // total number of nodes in tree
		n:      n,
		height: height,
	}

	// build segment tree
	st.buildTree(arr, 0, n-1, 0)

	return st
}

// building segment tree , following bottom-up approach
func (st *SegmentTree) buildTree(arr []int, start, end, node int) int {
	// base condition
	if start == end {
		st.tree[node] = arr[start]
		return st.tree[node]
	}

	mid := (start + end) / 2 // finding mid value

	// calculating sum of left segment and right segment and then assigning sum to current node
	st.tree[node] = st.buildTree(arr, start, mid, 2*node+1) + st.buildTree(arr, mid+1, end, 2*node+2)

	// return the current node value
	return st.tree[node]
}

// find sum by range
func (st *SegmentTree) QueryRange(left, right int) int {
	return st.queryRangeHelper(0, st.n-1, left, right, 0)
}

func (st *SegmentTree) queryRangeHelper(start, end, left, right, node int) int {
	// egde case : when query is out of segment
	if left > end || right < start {
		return 0
	}

	// base condition : segment must be present inside range
	if left <= start && right >= end {
		return st.tree[node]
	}

	mid := (start + end) / 2

	// recursion to find sum of segment nodes and returning the value
	return st.queryRangeHelper(start, mid, left, right, 2*node+1) + st.queryRangeHelper(mid+1, end, left, right, 2*node+2)
}

// update in segment tree
func (st *SegmentTree) UpdatePoint(index, value int) {
	st.updatePointHelper(0, st.n-1, index, value, 0)
}

func (st *SegmentTree) updatePointHelper(start, end, index, value, node int) {
	// edge case : when index is out of range
	if index < start || index > end {
		return
	}

	// base case for update
	if start == end {
		st.tree[node] = value
		return
	}

	mid := (start + end) / 2

	// first update left part of segment tree then right
	st.updatePointHelper(start, mid, index, value, 2*node+1)
	st.updatePointHelper(mid+1, end, index, value, 2*node+2)

	// update the sum value for respective parent nodes in tree
	st.tree[node] = st.tree[2*node+1] + st.tree[2*node+2]
}

func main() {
	log.Println("Segment Tree Implementation for range based queries")

	arr := []int{1, 3, 5, 7, 9, 11}
	st := NewSegmentTree(arr)
	log.Println(st.tree)
	log.Println("Sum of values in range [1, 3]:", st.QueryRange(1, 3))
	st.UpdatePoint(2, 10)
	log.Println("Tree after update : ", st.tree)
	log.Println("After update => Sum of values in range [1, 3]:", st.QueryRange(1, 3))
}
