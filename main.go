package main

import "fmt"

// func lenAndUpper(name string) (int, string) {
// 	return len(name), strings.ToUpper(name)
// }

//start point
func main() {
	//array = need [number] according to number
	//slice = dont need [number] and no limit
	names := []string{"nico", "han", "bae", "lee"}
	names = append(names, "choi")
	fmt.Println(names)
}
