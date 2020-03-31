package main

import (
	"fmt"
	"strings"
)

func main() {

	keyWords := []string{"dream", "dreamer", "erase", "eraser"}

	var s string

	fmt.Scanf("%s", &s)

	for {
		var status bool
		if s != "" {
			for _, v := range keyWords {
				if strings.HasSuffix(s, v) {
					s = strings.TrimSuffix(s, v)
					status = true
				}
			}
			if !status {
				break
			}
		} else {
			break
		}
	}
	if s != "" {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}

}
