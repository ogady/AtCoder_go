package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var n int
	var aa []string
	var ia []int
	var alice int
	var bob int

	fmt.Scanf("%d", &n)
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	aa = strings.Fields(s.Text())
	for _, v := range aa {
		i, _ := strconv.Atoi(v)
		ia = append(ia, i)
	}
	sort.Slice(ia, func(i, j int) bool { return ia[i] > ia[j] })

	for i, v := range ia {
		if i == 0 || i%2 == 0 {
			alice += v
		} else {
			bob += v
		}
	}

	fmt.Println(alice - bob)
}
