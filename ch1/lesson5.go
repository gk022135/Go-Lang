package goplio


func main() {
	// learning about arrays in go
	var a [5]int
	fmt.Println("emp:", a)

	// set value at index 4
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// length of array
	fmt.Println("len:", len(a))

	// declare and initialize an array
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// multidimensional array
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// array of strings
	var s [3]string
	s[0] = "Hello"
	s[1] = "World"
	s[2] = "Go"
	fmt.Println("string array:", s)

	// array of booleans
	var bools [3]bool
	bools[0] = true
	bools[1] = false
	bools[2] = true
	fmt.Println("boolean array:", bools)

	// array of floats
	var floats [3]float64
	floats[0] = 1.1
	floats[1] = 2.2
	floats[2] = 3.3
	fmt.Println("float array:", floats)

	// array of arrays
	var arrayOfArrays [2][2]string
	arrayOfArrays[0][0] = "a"
	arrayOfArrays[0][1] = "b"
	arrayOfArrays[1][0] = "c"
	arrayOfArrays[1][1] = "d"
	fmt.Println("array of arrays:", arrayOfArrays)
	
	// array of structs
	type person struct {
		name string
		age  int
	}
	var people [2]person
	people[0] = person{name: "Alice", age: 30}
	people[1] = person{name: "Bob", age: 25}
	fmt.Println("array of structs:", people)

	// array of pointers
	var pointers [2]*int
	a1, a2 := 10, 20
	pointers[0] = &a1
	pointers[1] = &a2
	fmt.Println("array of pointers:", pointers)
	fmt.Println("values pointed to:", *pointers[0], *pointers[1])

	// array of interfaces
	var interfaces [2]interface{}
	interfaces[0] = "Hello"
	interfaces[1] = 123
	fmt.Println("array of interfaces:", interfaces)

	// array of functions
	var funcs [2]func(int) int
	funcs[0] = func(x int) int { return x + 1 }
	funcs[1] = func(x int) int { return x * 2 }
	fmt.Println("array of functions:", funcs)
	fmt.Println("calling functions:", funcs[0](10), funcs[1](10))

	// array of complex numbers
	var complexes [2]complex128
	complexes[0] = complex(1, 2)
	complexes[1] = complex(3, 4)
	fmt.Println("array of complex numbers:", complexes)

	// array of bytes
	var bytes [3]byte
	bytes[0] = 'a'
	bytes[1] = 'b'
	bytes[2] = 'c'
	fmt.Println("array of bytes:", bytes)

	// array of runes
	var runes [3]rune
	runes[0] = 'あ'
	runes[1] = 'い'
	runes[2] = 'う'
	fmt.Println("array of runes:", runes)

	// array of uints
	var uints [3]uint
	uints[0] = 1
	uints[1] = 2
	uints[2] = 3
	fmt.Println("array of uints:", uints)

	// array of ints with ellipsis
	c := [...]int{1, 2, 3, 4, 5}
	fmt.Println("array with ellipsis:", c)

	// iterating over array using for loop
	for i := 0; i < len(c); i++ {
		fmt.Println("c[", i, "] =", c[i])
	}

	// iterating over array using range
	for i, v := range c {
		fmt.Println("c[", i, "] =", v)
	}

	// modifying array elements
	for i := range c {
		c[i] = c[i] * 2
	}
	fmt.Println("modified array:", c)

	// multi-dimensional array
	var multiD [2][2]int
	multiD[0][0] = 1
	multiD[0][1] = 2
	multiD[1][0] = 3
	multiD[1][1] = 4
	fmt.Println("multi-dimensional array:", multiD)

	// passing array to function
	printArray(c)

	// returning array from function
	d := getArray()
	fmt.Println("array from function:", d)



	// slicing array
	slice := c[1:4]
	fmt.Println("sliced array:", slice)

	// modifying sliced array
	slice[0] = 100
	fmt.Println("modified slice:", slice)
	fmt.Println("original array after modifying slice:", c)	


	//dynamic array using slices
	var s1 []int
	fmt.Println("emp:", s1)

	// append to slice
	s1 = append(s1, 1)
	s1 = append(s1, 2, 3)
	fmt.Println("apd:", s1)

	// length and capacity of slice
	fmt.Println("len:", len(s1))
	fmt.Println("cap:", cap(s1))

	// declare and initialize a slice
	s2 := []int{4, 5, 6}
	fmt.Println("dcl:", s2)

	// make a slice
	s3 := make([]int, 5)
	fmt.Println("make:", s3)

	// make a slice with capacity
	s4 := make([]int, 5, 10)
	fmt.Println("make with cap:", s4)
	fmt.Println("len:", len(s4))
	fmt.Println("cap:", cap(s4))

	// slicing a slice
	s5 := s2[1:3]
	fmt.Println("sliced slice:", s5)

	// iterating over slice using for loop
	for i := 0; i < len(s2); i++ {
		fmt.Println("s2[", i, "] =", s2[i])
	}

	// iterating over slice using range
	for i, v := range s2 {
		fmt.Println("s2[", i, "] =", v)
	}

	// modifying slice elements
	for i := range s2 {
		s2[i] = s2[i] * 2
	}
	fmt.Println("modified slice:", s2)

	// multi-dimensional slice
	var multiDSlice [][]int
	multiDSlice = append(multiDSlice, []int{1, 2})
	multiDSlice = append(multiDSlice, []int{3, 4})
	fmt.Println("multi-dimensional slice:", multiDSlice)

	// passing slice to function
	printSlice(s2)
}