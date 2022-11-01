package main
import (
    "fmt"
    "strconv"
    )

func isPalindrome(x int) bool {
    str:=strconv.Itoa(x)
       var revStr string
       for i:=len(str)-1;i>=0;i--{
           revStr+=string(str[i])
       }
    if str == revStr{
        return true 
    }
    return false
}


func main(){
  var num int  //creating a integer variable
  fmt.Println("Enter a number greater than 9")
  fmt.Scan(&num)
  fmt.Println(isPalindrome(num))
}
