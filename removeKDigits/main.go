package main

import (
	"fmt"
	"errors"
	"strings"
)

type Stack struct{
	item []interface{}
}

func (s *Stack) push(str interface{}){

	s.item = append(s.item,str)

}


func (s *Stack) pop(){
	l :=len(s.item)
	if l > 0{
		s.item = s.item[:l-1]
	}
}

func (s *Stack) top() (interface{}, error) {

	l := len(s.item)
	if l ==0 {

		return nil, errors.New("Stack is empty")
	}
	return s.item[l-1], nil
}

func (s *Stack) isEmpty()bool{

	return len(s.item)==0

}

func main(){

	num := "100"
	k :=1
	s := removeKdigits(num, k)
	fmt.Println(s)



}
func removeKdigits(num string, k int) string {

    l := len(num)

	var fs string 

	if l == k{
 		fs = "0"
		 return fs
	}
	
    
    s :=Stack{} 
    var numberOfPop int

	


    for i :=0; i < l; i++{

		// Check if Stack is empty, add the first element
		// continue with next iteration in the loop

		topElement, err := s.top()
		if(err !=nil) {
			s.push(string(num[i]))
			
			continue
		}
	
		 currentElement := string(num[i])	 

		// Pop element from the Stack if the current element is less 
		// than the number at the top

		if (currentElement < topElement.(string)){
			
			if(numberOfPop < k){
			s.pop() //Pop the Top element
			s.push(currentElement)
			numberOfPop++
			continue
			}

		}
		if (len(s.item) < (l - k)){
			s.push(currentElement)

		}

    }

	



  	for _, v := range s.item{
		fs = fs + v.(string)
  	}
	  // Remove leading zeros from the result.
	fs = strings.TrimLeft(fs, "0")
	
	if fs !=""{
		return fs
	}else{
		return "0"
	}
    
}

