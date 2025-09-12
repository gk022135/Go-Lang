package ch1
//learning functions in go
import "fmt"

func Lesson8() {
	//function declaration and definition
	add := func(a int, b int) int {
		return a + b
	}

	//function call
	result := add(10, 20)
	fmt.Println("Result of add(10, 20) =", result)

	//function with multiple return values
	swap := func(a int, b int) (int, int) {
		return b, a
	}

	//function call with multiple return values
	x, y := swap(10, 20)
	fmt.Println("Result of swap(10, 20) =", x, y)

	//function with named return values
	multiply := func(a int, b int) (result int) {
		result = a * b
		return
	}

	//function call with named return values
	result = multiply(10, 20)
	fmt.Println("Result of multiply(10, 20) =", result)

	//function as first class citizen
	operate := func(a int, b int, op func(int, int) int) int {
		return op(a, b)
	}

	//function call with function as argument
	result = operate(10, 20, add)
	fmt.Println("Result of operate(10, 20, add) =", result)

	//anonymous function
	func(msg string) {
		fmt.Println(msg)
	}("Hello from anonymous function")

	//immediately invoked function expression (IIFE)
	result = func(a int, b int) int {
		return a - b
	}(20, 10)
	fmt.Println("Result of IIFE (20 - 10) =", result)

	//closure
	counter := func() func() int {
		count := 0
		return func() int {
			count++
			return count
		}
	}()

	//calling closure multiple times
	fmt.Println("Counter:", counter())
	fmt.Println("Counter:", counter())
	fmt.Println("Counter:", counter())
}	