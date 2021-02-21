package main

import "fmt"

// func lenAndUpper(name string) (int, string) {
// 	return len(name), strings.ToUpper(name)
// }

//start point
func main() {
	//array = need [number]
	//slice = dont need [number]
	names := []string{"nico", "han", "bae", "lee"}
	names = append(names, "choi")
	fmt.Println(names)
}
