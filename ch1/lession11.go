package goplio

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


//by address passing the values
func main(){
	num := 1

	changeNUm((num))

	fmt.Println("after changesNUm in Main", num)
}