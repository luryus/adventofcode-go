package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewReader(os.Stdin)
	data, _ := sc.ReadString(0)

	dirs := strings.Split(strings.TrimSpace(data), ",")

	x, y, z := 0, 0, 0
	max := float64(0)

	for _, d := range dirs {
		if d == "n" {
			x++
			y--
		} else if d == "s" {
			y++
			x--
		} else if d == "nw" {
			x++
			z--
		} else if d == "se" {
			x--
			z++
		} else if d == "sw" {
			y++
			z--
		} else if d == "ne" {
			z++
			y--
		}

		m := dist(x, y, z)
		if m > max {
			max = m
		}
	}

	fmt.Println(dist(x, y, z))
	fmt.Println(max)
}

func dist(x, y, z int) float64 {
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	if z < 0 {
		z = -z
	}
	return float64(x+y+z) / 2
}
