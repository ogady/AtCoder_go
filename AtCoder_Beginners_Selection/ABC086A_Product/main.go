package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	r := (a * b) % 2
	if r == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
