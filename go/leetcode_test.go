package main_test

import (
	"fmt"
	"testing"
)

func Test00042(t *testing.T) {
	trap := func(height []int) int {
		length := len(height)
		leftMax := make(map[int]int, len(height))
		rightMax := make(map[int]int, len(height))
		leftM := 0
		rightM := 0
		for i := range height {
			if leftM > height[i] {
			} else {
				leftM = height[i]
			}
			leftMax[i] = leftM

			if rightM > height[length-1-i] {
			} else {
				rightM = height[length-1-i]
			}
			rightMax[length-1-i] = rightM
		}

		s := 0
		for i := range height {
			if leftMax[i] > rightMax[i] {
				s += rightMax[i] - height[i]
			} else {
				s += leftMax[i] - height[i]
			}
		}
		return s
	}

	input := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}

	fmt.Println(trap(input))
}

func Test00146(t *testing.T) {
	// type LRUCache struct {

	// }

	// Constructor :=func (capacity int) LRUCache {

	// }

	// Get := func (this *LRUCache) (key int) int {

	// }

	// func (this *LRUCache) Put(key int, value int)  {

	// }
}
