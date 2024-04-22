package main

import (
	"fmt"
	"github.com/atulhere/go-utils/stack"
)




func main(){
	str := "abbaca"
	fmt.Println("Input Sring is ", str)
	str = removeAdjacentDuplicates(str)
	fmt.Println("Output Sring is ", str)
}

func removeAdjacentDuplicates(str string) string{

		//Create an instance of Stack
		st :=stack.Stack{}

		for _, v :=range str{
			
			// if Stack is empty push the element
			if(st.IsEmpty()){
				st.Push(string(v))
				continue
			}

			// Check if the current element is already 
			// at the Top of the Stack, Pop the element at the Top 
			// and do not push new Element

			item,err := st.Top()
			if(err == nil){
				if(string(v) == item){
					st.Pop()
					continue
				}	
				st.Push(string(v))
			}
			
		}

	updatedStr := ""
	for _, v := range st.Item{
		updatedStr = updatedStr + v.(string)

	}

	return updatedStr
}


