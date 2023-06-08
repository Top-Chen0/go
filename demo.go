package main

import "fmt"

//func main() {
//	fmt.Print("test")
//}

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Print(add(42, 13))
}
