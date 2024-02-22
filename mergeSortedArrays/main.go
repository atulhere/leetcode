/*
Merge Sorted Array
You are given two integer arrays nums1 and nums2,
sorted in non-decreasing order, and two integers m and n,
representing the number of elements in nums1 and nums2
respectively.

Merge nums1 and nums2 into a single array sorted in
non-decreasing order.

The final sorted array should not be returned by
the function, but instead be stored inside the array
nums1. To accommodate this, nums1 has a length of m + n,
where the first m elements denote the elements
that should be merged, and the last n elements are
set to 0 and should be ignored. nums2 has a length of n.
*/

package main

import (
	"fmt"
)

func main() {

	num1 := []int{1, 2, 9, 0, 0, 0}
	num2 := []int{2, 5, 6}
	merge(num1, 3, num2, 3)
}

func merge(nums1 []int, m int, nums2 []int, n int) {

	fmt.Println("The given array are: ", nums1, nums2)

	i, j, k := m-1, n-1, (m+n)-1

	for i >= 0 && j >= 0 {

		if nums1[i] >= nums2[j] {
			nums1[k] = nums1[i]
			i--
			k--
		} else {
			nums1[k] = nums2[j]
			j--
			k--
		}
	}

	//if num2 have still some elements, push them into num1
	for j >= 0 {

		nums1[k] = nums2[j]
		j--
		k--

	}

	fmt.Println("The merged sorted array is", nums1)

}
