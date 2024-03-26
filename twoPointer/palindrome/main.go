package main

import(

	"fmt"
	"regexp"
	"strings"
)

func main(){

	s := "ab_a"

	result := isPalindrome(s)

    if (result){
		fmt.Println("String is palindrome!")

	}else{
		fmt.Println("String is not palindrome!")
	}

}

func isPalindrome(s string) bool {
    
	// Remove Special char from the input string if any
	reg :=regexp.MustCompile(`[\W_]+`)
    cs := reg.ReplaceAllString(s,"")
	fmt.Println(cs)
    left :=0
    right :=len(cs) -1

    for left < right {
        ll := strings.ToLower(string(cs[left])) 
        rr := strings.ToLower(string(cs[right])) 

        if(ll==rr){

            left++
            right--
        }else{
            return false
        }

      
    }

      return true


}