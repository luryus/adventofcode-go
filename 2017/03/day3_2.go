package main

import (
	"fmt"
	"os"
	"strconv"
)

func get_elem(arr []int32, x, y, w, h int) int32 {
	if x < 0 || y < 0 || x >= w || y >= h {
		return 0
	}

	return arr[y*w+x]
}
func set_elem(arr []int32, x, y, w int, val int32) {
	arr[y*w+x] = val
}
func sum_around(arr []int32, x, y, w, h int) int32 {
	return get_elem(arr, x-1, y-1, w, h) +
		get_elem(arr, x-1, y, w, h) +
		get_elem(arr, x-1, y+1, w, h) +
		get_elem(arr, x, y-1, w, h) +
		get_elem(arr, x, y+1, w, h) +
		get_elem(arr, x+1, y-1, w, h) +
		get_elem(arr, x+1, y, w, h) +
		get_elem(arr, x+1, y+1, w, h)
}

const n = 1333

func main() {
	x, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	arr := make([]int32, n*n)
	set_elem(arr, n/2, n/2, n, 1)
	r := 1

	for r < n/2 {
		i, j := n/2+r, n/2+r-1
		ij_min, ij_max := n/2-r, n/2+r

		for {
			set_elem(arr, i, j, n, sum_around(arr, i, j, n, n))

			fmt.Printf("%d\n", get_elem(arr, i, j, n, n))
			if get_elem(arr, i, j, n, n) >= int32(x) {
				return
			}

			if i == ij_max && j == ij_max {
				r++
				break
			} else if i == ij_max && j > ij_min {
				j--
			} else if j == ij_min && i > ij_min {
				i--
			} else if i == ij_min && j < ij_max {
				j++
			} else if j == ij_max && i < ij_max {
				i++
			}
		}
	}
}
