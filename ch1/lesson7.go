package ch1

//learning ranges in go
import "fmt"

func Lesson7() {
	//array declaration and initialization
	arr := [5]int{1, 2, 3, 4, 5}

	//iterating over array using for range
	for i, v := range arr {
		fmt.Println("arr[", i, "] =", v)
	}

	//iterating over array using for loop
	for i := 0; i < len(arr); i++ {
		fmt.Println("arr[", i, "] =", arr[i])
	}

	//slice declaration and initialization
	slice := []int{10, 20, 30, 40, 50}

	//iterating over slice using for range
	for i, v := range slice {
		fmt.Println("slice[", i, "] =", v)
	}

	//iterating over slice using for loop
	for i := 0; i < len(slice); i++ {
		fmt.Println("slice[", i, "] =", slice[i])
	}

	//string declaration and initialization
	str := "hello"

	//iterating over string using for range
	for i, c := range str {
		fmt.Println("str[", i, "] =", string(c))
	}

	//iterating over string using for loop
	for i := 0; i < len(str); i++ {
		fmt.Println("str[", i, "] =", string(str[i]))
	}

	//map declaration and initialization
	m := map[string]int{"a": 1, "b": 2, "c": 3}

	//iterating over map using for range
	for k, v := range m {
		fmt.Println("m[", k, "] =", v)
	}

	//iterating over map using for loop
	keys := []string{}
	for k := range m {
		keys = append(keys, k)
	}
	for i := 0; i < len(keys); i++ {
		fmt.Println("m[", keys[i], "] =", m[keys[i]])
	}
}