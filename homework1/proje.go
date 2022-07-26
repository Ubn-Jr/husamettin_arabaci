package main

import "fmt"

func main() {

	var x int
	var y int
	var z int
	x = 10
	y = 20
	z = add(x, y)
	fmt.Println(z)
}

func add(x int, y int) int {
	var z int
	z = x + y
	return z
}
