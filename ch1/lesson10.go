package ch1

import "fmt"

//learning closers
/*
what is closure
--->kisi function ke andar koi outer scope ka variable use ho rha hai, toh use function ke andar hames available rhege, woh function execute hone ke baad bhi

matlab innere funtion mai hames out function wala variable available rahega
*/

func counter() func() int {
	var count int = 0

	return func() int{
		count+=1
	}
}


func main(){

	increment := counter()

	fmt.Println(increment())
}