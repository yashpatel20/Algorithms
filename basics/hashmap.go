package basics

import (
	"fmt"
)

//Hashmap basic hashmap in go
func Hashmap() {
	m := make(map[string]int)
	fmt.Println("create", m)

	//set
	m["hello"] = 1
	m["namaste"] = 2
	m["bonjour"] = 3
	fmt.Println("set", m)

	//delete
	delete(m, "hello")
	fmt.Println("delete", m)

	//traverse
	fmt.Println("traverse")
	for k, v := range m {
		fmt.Println(k, v)
	}
}
