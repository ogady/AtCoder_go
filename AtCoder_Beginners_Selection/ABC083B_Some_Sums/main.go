package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n, a, b, sum int
	fmt.Scanf("%d %d %d", &n, &a, &b)
	for i := 0; i <= n; i++ {
		s := strconv.Itoa(i)
		var x int
		for _, v := range s {
			j, _ := strconv.Atoi(string(v))
			x += j
		}
		if x >= a && x <= b {
			sum += i
		}
	}

	fmt.Println(sum)
}
