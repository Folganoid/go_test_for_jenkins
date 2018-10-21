package main

import "fmt"

func sum(a, b int64) int64 {
	return a + b
}

func mult(a, b int64) int64 {
	return a * b
}


func main() {

	fmt.Println(sum(10, 10))
	fmt.Println(mult(10, 10))

}