package main
import "fmt"


//go does not have ternary operator

func main() {
	//learning conditional statements
	x := 10
	if x%2 == 0 {
		fmt.Println("x is even")
	} else {
		fmt.Println("x is odd")
	}


	//direct variable declaration in if statement
	if x := 20; x%2 == 0 {
		fmt.Println("x is even")
	} else {
		fmt.Println("x is odd")
	}

	//else if  ladder
	if x < 0 {
		fmt.Println("x is negative")
	} else if x == 0 {
		fmt.Println("x is zero")
	} else {
		fmt.Println("x is positive")
	}

	if x < 0 {
		fmt.Println("x is negative")
	} else if x == 0 {
		fmt.Println("x is zero")
	} else if x > 0 && x <= 10 {
		fmt.Println("x is positive and less than or equal to 10")
	} else {
		fmt.Println("x is positive and greater than 10")
	}	




	//switch case
	switch x := 10; {
	case x < 0:
		fmt.Println("x is negative")
	case x == 0:
		fmt.Println("x is zero")
	case x > 0:
		fmt.Println("x is positive")
	}

	switch x := 10; x {
	case 1, 3, 5, 7, 9:
		fmt.Println("x is odd")
	case 2, 4, 6, 8, 10:
		fmt.Println("x is even")
	default:
		fmt.Println("x is out of range")
	}

	switch {
	case x < 0:
		fmt.Println("x is negative")
	case x == 0:
		fmt.Println("x is zero")
	case x > 0:
		fmt.Println("x is positive")
	}

	switch x := 10; {
	case x < 0:
		fmt.Println("x is negative")
		fallthrough
	case x == 0:
		fmt.Println("x is zero")
		fallthrough
	case x > 0:
		fmt.Println("x is positive")
	}

	switch x := 10; {
	case x < 0:
		fmt.Println("x is negative")
	case x == 0:
		fmt.Println("x is zero")
	case x > 0:
		fmt.Println("x is positive")
	}

	switch x := 10; x {
	case 1, 3, 5, 7, 9:
		fmt.Println("x is odd")
	case 2, 4, 6, 8, 10:
		fmt.Println("x is even")
	default:
		fmt.Println("x is out of range")
	}

	switch {
	case x < 0:
		fmt.Println("x is negative")
	case x == 0:
		fmt.Println("x is zero")
	case x > 0:
		fmt.Println("x is positive")
	}

	switch x := 10; {
	case x < 0:
		fmt.Println("x is negative")
		fallthrough
	case x == 0:
		fmt.Println("x is zero")
		fallthrough
	case x > 0:
		fmt.Println("x is positive")
	}

}