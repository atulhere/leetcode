// Reverse an array in place
package main

import "fmt"

func main() {

	array := []int{1, 2, 3, 4, 5, 6, 7}
	reverse(array)

}

func reverse(array []int) []int {

	middleNumber := len(array) / 2
	i := 0
	j := len(array) - 1
	for i < middleNumber && j >= middleNumber {

		array[i], array[j] = array[j], array[i]

		i++
		j--
	}

	fmt.Println("Reverse of array is ", array)
	return array

}
