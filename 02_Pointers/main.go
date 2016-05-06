package main

import "fmt"

func zero(z *int) {
	*z = 0
}
func main() {
	x := 5
	fmt.Printf("Before zero function %d \n", x)
	zero(&x)
	fmt.Printf("After zero it becomes %d \n", x) //X is still 5
}
