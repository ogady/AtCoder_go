package main

import (
	"fmt"
	"math"
)

type plan struct {
	t float64
	x float64
	y float64
}

func main() {
	var n int
	fmt.Scanf("%d", &n)

	plans := make([]plan, n+1)
	plans[0] = plan{
		t: 0,
		x: 0,
		y: 0,
	}

	var t, x, y float64
	for i := 1; i <= n; i++ {
		fmt.Scanf("%f %f %f", &t, &x, &y)
		plans[i] = plan{
			t: t,
			x: x,
			y: y,
		}
	}

	var canMove bool

	for i := 1; i <= n; i++ {

		dist := math.Abs(plans[i].x-plans[i-1].x) + math.Abs(plans[i].y-plans[i-1].y)
		dt := math.Abs(plans[i].t - plans[i-1].t)

		if dist <= dt && judgeEven(dist) == judgeEven(dt) {
			canMove = true
		} else {
			canMove = false
			break
		}
	}

	if canMove {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func judgeEven(tar float64) bool {
	if int(tar)%2 == 0 {
		return true
	}
	return false
}
