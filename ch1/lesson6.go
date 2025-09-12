package ch1
//leaarning maps in go
import "fmt"

func Lesson6(){
	//map declaration
	m := make(map[string]int)

	//map initialization
	m["a"] = 1
	m["b"] = 2
	m["c"] = 3

	//map declaration and initialization in one line
	n := map[string]int{"x": 10, "y": 20, "z": 30}

	//map access
	fmt.Println("m:", m)
	fmt.Println("n:", n)
	fmt.Println("m[a]:", m["a"])
	fmt.Println("n[x]:", n["x"])

	//map length
	fmt.Println("len(m):", len(m))
	fmt.Println("len(n):", len(n))

	//map deletion
	delete(m, "b")
	fmt.Println("m after deletion:", m)

	//map existence check
	v, ok := m["b"]
	fmt.Println("m[b]:", v, "exists:", ok)
	v, ok = m["a"]
	fmt.Println("m[a]:", v, "exists:", ok)

	//iterating over map using for range
	for k, v := range n {
		fmt.Println("n[", k, "] =", v)
	}

	//iterating over map using for loop
	keys := []string{}
	for k := range n {
		keys = append(keys, k)
	}
	for i := 0; i < len(keys); i++ {
		fmt.Println("n[", keys[i], "] =", n[keys[i]])
	}

	//nested maps
	nestedMap := map[string]map[string]int{
		"first":  {"a": 1, "b": 2},
		"second": {"x": 10, "y": 20},
	}
	fmt.Println("nestedMap:", nestedMap)
	fmt.Println("nestedMap[first]:", nestedMap["first"])
	fmt.Println("nestedMap[first][a]:", nestedMap["first"]["a"])
}