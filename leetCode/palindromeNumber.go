package main
import (
    "fmt"  
    "strconv" //strconv for to and from string conversions
    )

func isPalindrome(x int) bool {
    str:=strconv.Itoa(x) //converts integer x to string
       var revStr string  
       for i:=len(str)-1;i>=0;i--{  // for loop for iterating through the string in reverse 
           revStr+=string(str[i])
       }
    if str == revStr{ //Checks if the string  and reverse of the string is equal
        return true 
    }
    return false
}


func main(){
  var num int  //creating an integer variable
  fmt.Println("Enter a number greater than 9")
  fmt.Scan(&num) //accepting the number from user
  fmt.Println(isPalindrome(num)) //calling the ispalindrome function
}
