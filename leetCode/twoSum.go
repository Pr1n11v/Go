package main
import "fmt"

func twoSum(nums []int, target int) []int {
    for i,data:=range nums{
      for j,data2:=range nums{
        if i==j{
            continue
        }else{
        if (data+data2)==target{
            return []int{i,j} 
        }
        }
    }  
    }
    return []int{0}
}

func main(){
  var size,target int 
  fmt.Println("Enter the array size and the target sum")
  fmt.Scan(&size,&target)
  var arr = make([]int,size)
  fmt.Println("Enter the array elements")
  for i:=0;i<size;i++{
       fmt.Scanf("%d",&arr[i])
   
  } 
  
   arr2:=twoSum(arr,target)
   fmt.println(arr2)
 
}

