package ch1

import "fmt"

//learning pointer
/*
---> poit the memory loacation of variables
*/

//by values pass hots, copy of variales goes
func changeNUm(num int){
	num = 5;
	fmt.Println("int changeNUm", num)
}


//by address passing the values -->> sharing refrence of variables

func changeNUmRef(num *int){
	*num = 5
	fmt.Println("in change NUm Ref", &num)
}


func main(){
	num := 1

	changeNUm((num))

	fmt.Println("Memory address", &num)

	fmt.Println("after changesNUm in Main", num)

	fmt.Print("after passing the values via refere")

	changeNUmRef(&num);

	fmt.Println("value of num after ref changes is ", num)
}