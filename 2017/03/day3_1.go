package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	x, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}
	n := 1
	for x > n*n {
		n += 2
	}
	fmt.Printf("n: %d\n", n)

	center_dist := n / 2

	x_off := x - (n-2)*(n-2)
	for x_off > n-1 {
		fmt.Printf("x_off: %d\n", x_off)
		x_off -= (n - 1)
	}

	fmt.Printf("x_off: %d\n", x_off)

	edge_dist := x_off - ((n - 1) / 2)
	if edge_dist < 0 {
		edge_dist *= -1
	}

	fmt.Printf("Center: %d, Edge: %d => %d\n", center_dist, edge_dist, center_dist+edge_dist)
}
