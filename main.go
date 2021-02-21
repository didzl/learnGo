package main

import "fmt"

// func lenAndUpper(name string) (int, string) {
// 	return len(name), strings.ToUpper(name)
// }

//start point
func main() {
	//structs
	favFood := []string{"bulgogi", "galbi"}
	han:= person{name : "han", age: 33, favFood: favFood}
	bae:= person{"bae", 33, favFood}
	fmt.Println(han)
	fmt.Println(han.favFood)
	fmt.Println(bae)
}

// first of all, you define structs
type person struct {
	name string
	age int
	favFood []string

}
