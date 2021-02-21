package main

import "fmt"

// func lenAndUpper(name string) (int, string) {
// 	return len(name), strings.ToUpper(name)
// }

func canIDrink(age int) bool {
	switch koreanAge :=age+ 2; koreanAge{
	case 18:
		return false
	case 8:
		return true
	}
	return false
}


//start point
func main() {
	fmt.Println(canIDrink(16))
}
