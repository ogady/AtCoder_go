package main

import "fmt"

func main() {
	var cnt, tar int
	fmt.Scanf("%d %d", &cnt, &tar)

	for i := 0; i <= cnt; i++ {
		for j := 0; j <= cnt; j++ {
			sum := 10000*i + 5000*j + 1000*(cnt-i-j)
			if (cnt - i - j) >= 0 {
				if sum == tar {
					fmt.Println(i, j, (cnt - i - j))
				}
			}
		}
	}
	fmt.Println(-1, -1, -1)
}
