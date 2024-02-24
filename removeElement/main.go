package main

import "fmt"

func removeElement(nums []int, n int) int {

	i, j := 0, 0

	for i < len(nums) {

		if nums[i] != n {
			nums[j] = nums[i]
			j++
		}
		i++
	}

	fmt.Println("The array is ", nums[:j])
	return j

}

func main() {

	nums := []int{1, 1, 2, 3, 2, 1, 2, 5}
	n := 2
	fmt.Println("The given array is ", nums)
	fmt.Println("We need to remove ", n)

	noOfElements := removeElement(nums, n)
	fmt.Printf("\nNumber of elements after removing %d from the array are %d ", n, noOfElements)

}
