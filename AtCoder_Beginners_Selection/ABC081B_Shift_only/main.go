package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var cnt int
	var n int
	var ns []int
	var existOdd bool

	fmt.Scanf("%d", &n)
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	ss := strings.Fields(s.Text())

	for _, v := range ss {
		j, _ := strconv.Atoi(v)
		ns = append(ns, j)
	}

	for {
		for _, v := range ns {

			if v%2 == 1 {
				existOdd = true
			}
		}

		if existOdd {
			break
		}

		for i, v := range ns {
			ns[i] = v / 2
		}
		cnt++
	}

	fmt.Println(cnt)
}
