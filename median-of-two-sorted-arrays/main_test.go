package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMediam2Arrays(t *testing.T) {
	assert.Equal(t, float64(0), findMedianSortedArrays([]int{0, 0}, []int{0, 0}))
	assert.Equal(t, float64(5), findMedianSortedArrays([]int{1, 2, 3, 4}, []int{5, 6, 7, 8, 9}))
	assert.Equal(t, 4.5, findMedianSortedArrays([]int{1, 2, 4, 6}, []int{3, 5, 7, 8}))
	assert.Equal(t, float64(2), findMedianSortedArrays([]int{1, 3}, []int{2}))
	assert.Equal(t, 2.5, findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
	assert.Equal(t, float64(1), findMedianSortedArrays([]int{}, []int{1}))
	assert.Equal(t, float64(1), findMedianSortedArrays([]int{1}, []int{}))
	assert.Equal(t, float64(1), findMedianSortedArrays([]int{1}, []int{1}))
	assert.Equal(t, float64(2.5), findMedianSortedArrays([]int{1, 2, 3, 4}, []int{}))
	assert.Equal(t, float64(-1), findMedianSortedArrays([]int{3}, []int{-2, -1}))
	assert.Equal(t, float64(1.5), findMedianSortedArrays([]int{1, 2}, []int{1, 2}))
	assert.Equal(t, float64(2), findMedianSortedArrays([]int{1, 2}, []int{1, 2, 3}))
	assert.Equal(t, float64(2), findMedianSortedArrays([]int{1, 1, 3, 3}, []int{1, 1, 3, 3}))
	assert.Equal(t, float64(1.5), findMedianSortedArrays([]int{1, 2}, []int{-1, 3}))
	assert.Equal(t, float64(2.5), findMedianSortedArrays([]int{1}, []int{2, 3, 4}))
	assert.Equal(t, float64(2.5), findMedianSortedArrays([]int{4}, []int{1, 2, 3}))
	assert.Equal(t, float64(2.5), findMedianSortedArrays([]int{1, 2, 3}, []int{4}))
	assert.Equal(t, float64(2.5), findMedianSortedArrays([]int{1, 3}, []int{2, 4}))
}
