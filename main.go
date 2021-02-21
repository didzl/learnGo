package main

import "fmt"

// func lenAndUpper(name string) (int, string) {
// 	return len(name), strings.ToUpper(name)
// }

//start point
func main() {
	a:= 2
	b:= &a // point a's ram
	a = 5
	*b = 20 // a = 20(pointer) and see through wright a
	fmt.Println(a, b, *b, a)
}
