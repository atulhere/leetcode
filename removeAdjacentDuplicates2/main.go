package main

import (
	"fmt"
	"github.com/atulhere/go-utils/stack"
)




func main(){
	str := "yfttttfbbbbnnnnffbgffffgbbbbgssssgthyyyy"
	k :=4
	fmt.Println("Input Sring is ", str)
	str = removeDuplicates(str, k)
	fmt.Println("Output Sring is ", str)
}

func removeDuplicates(str string, k int) string {

	// Create a has map
	arr := make(map[string]int)

	s :=stack.Stack{}

	for _, v := range str {
		arr[string(v)]++
		if arr[string(v)] == k {
			//Remove that char from the map
			delete(arr, string(v))
			step :=k 
			for(step >0){
				s.Pop()
				step--
			}
			
		}
		s.Push(string(v))
	}
	
	// Now, iterate throuh Map and create a string
	// from the remining Character
	fmt.Println(s.Item)
	updatedString := ""
	for _, v :=range str{

		if(arr[string(v)]> 0){
			updatedString =  updatedString + string(v)
			arr[string(v)]--
		}
	}

	return updatedString
}


