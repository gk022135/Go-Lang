package ch1

import "fmt"

//vardic function in go
/*
use for:- receive N numbers of parameters
like printLn(1,1,1,12.23,45,2)
*/

//declaration
func sum(nums ...int) int{
	total := 0

	for _, num:= 0 range nums{
		total = total + num;
	}
	return total
}

func main(){
	result := sum(2,3,4,5);
	fmt.Println(sum)


	nums := []int(3,4,5,6) //slice

	res :=sum(nums...)
	fmt.Println((res))
}
