package main


func main(){
	// implemetation of while loop using for loop
	i := 1
	for i <= 10 {
		println(i)
		i = i + 1
	}
	println("Loop Ended")

	// infinite loop
	// for {
	// 	println("infinite loop")
	// }

	// for loop with init and post statement
	for j := 1; j <= 10; j++ {
		println(j)
	}

	// for loop with range
	nums := []int{2, 3, 4, 5}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	println("sum:", sum)

	// for loop with range to get index and value
	for i, num := range nums {
		if num == 3 {
			println("index:", i)
		}
	}
}