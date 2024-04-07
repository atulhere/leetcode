// Given an array of integers nums and an integer k, 
// return the number of contiguous subarrays where the product 
// of all the elements in the subarray is strictly less than k.

package main

import "fmt"
func main(){
	

	nums :=[]int{10,5,2,6}

	fmt.Println(&nums[0], &nums[1],&nums[2], &nums[3])

	arr := numSubarrayProductLessThanK(nums,100)

	// Bruteforce approach is to ues two for loops 
	// but that and find out all possible subarrays 

	fmt.Println("The number of sub arrays are : ", arr)

}

func numSubarrayProductLessThanK(nums []int, k int) int{

	count := 0
	
	for i :=0; i < len(nums); i++{
		product := 1
		fmt.Println(product)
		for j :=i; j< len(nums); j++{

			product = product * nums[j]
			if product < k{
				count++
			}

		}
	
	}
	return count
}




[1,2,3,3,3]

1- [1,2,3,3]
2- [1,2,3,3,3]
3- [2,3,3]
4- [2,3,3,3]
5-[3,3]
6-[3,3,3]