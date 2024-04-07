// Given an unsorted integer array nums. 
// Return the smallest positive integer that is not present in nums.
// You must implement an algorithm that runs in O(n) time 
// and uses O(1) auxiliary space.
package main

import "fmt"

func main(){

	nums :=[]int{1,2,6,3,5,4}

	n := firstMissingPositive(nums)

	fmt.Println("The number is :", n)

}

func checkNumberExist(temp map [int]bool, missingNumber int) int{

	if !temp[missingNumber]{

		return missingNumber
	}

	
   return checkNumberExist(temp, missingNumber +1 )



}


func firstMissingPositive(nums []int) int {

	// let's assume missingNumber is 1, 
	// which is First Postive integger
	missingNumber := 1 

	temp :=make(map[int]bool)
	
	for i :=0; i < len(nums); i++{

		temp[nums[i]] = true
		//check if the assumed missingNumber is already exist in nums array
		if(nums[i]== missingNumber  ){
			missingNumber++ 

		}

		
	}

	

	n := checkNumberExist(temp, missingNumber)
	return n
	
}