package main

import "fmt"

// func lenAndUpper(name string) (int, string) {
// 	return len(name), strings.ToUpper(name)
// }

//start point
func main() {
	//maps
	han := map[string]string{"name":"han", "age":"33"}
	for key, value := range han{
		fmt.Println(key, value)
		fmt.Println(value)
		
	}
}
