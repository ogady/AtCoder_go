package main

import "fmt"

func main() {
	var a, b, c, x, cnt int

	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	fmt.Scanf("%d", &c)
	fmt.Scanf("%d", &x)
	for A := 0; A <= a; A++ {
		for B := 0; B <= b; B++ {
			for C := 0; C <= c; C++ {
				i := 500*A + 100*B + 50*C
				if x == i {
					cnt++
				}
			}
		}
	}
	fmt.Println(cnt)
}
