// Given a string path, which is an absolute path (starting with a slash '/') 
// to a file or directory in a Unix-style file system, 
//  convert it to the simplified canonical path.
// In a Unix-style file system, a period '.' refers to the current directory,
// a double period '..' refers to the directory up a level, 
// and any multiple consecutive slashes (i.e. '//') are treated as a single slash '/'. 
// For this problem, any other format of periods such as '...' are treated as file/directory names.

package main

import (
	"fmt"
	"errors"
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

	str := "/a//b////c/d//././/.." 

	str = simplifyPath(str)

	fmt.Println(str)

}

func simplifyPath(str string) string{

	s :=Stack{}

	word :="" 
	output := ""

	for _, v := range str{

		currentChar := string(v)

		if (currentChar == "/" && word == ""){
			continue // move to the next iteration
		}

		if(currentChar !="/"){
			word = word + currentChar
		}

		if(currentChar == "/" && word != ""){

			if (word == "."){
				word =""

				continue // move to the next iteration because "."" represnt current directory
			}else if(word ==".."){

				// POP the last element if Stack is not empty
				if(len(s.item)> 0){
					s.pop()
					word =""
				}
				word =""

			}else{
				// push the element into Stack
				s.push(word)
				word = ""
			}

		}
	}

	//if last char is not "/", handle this case
	if(word == ".."){

		if(len(s.item)> 0){
			s.pop()
		}
	}
	if(word !=".." && word !="." && word !=""){
		s.push(word)
	}
	for _,v := range s.item{
		output = output + "/"+ v.(string) 
	}
	
	if(len(s.item) == 0){
		return "/"
	}
	return output
	
}