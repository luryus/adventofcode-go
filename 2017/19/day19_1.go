package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	var routes [][]byte
	get := func(x, y int) byte { return routes[y][x] }

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		routes = append(routes, []byte(sc.Text()))
	}

	rows, cols := len(routes), len(routes[0])

	// find start coord
	x, y := 0, 0
	for i := 0; i < cols; i++ {
		if get(i, 0) != ' ' {
			x = i
			break
		}
	}

	dirX, dirY := 0, 1

	x += dirX
	y += dirY

	steps := 1
	var letters []byte

	for x >= 0 && y >= 0 && x < cols && y < rows {
		c := get(x, y)
		if bytes.IndexByte([]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ"), c) != -1 {
			letters = append(letters, c)
		} else if c == '+' {
			// change dir
			if dirX == 0 {
				dirY = 0
				if get(x-1, y) != ' ' {
					dirX = -1
				} else if get(x+1, y) != ' ' {
					dirX = 1
				} else {
					panic("Illegal state in +")
				}
			} else {
				dirX = 0
				if get(x, y-1) != ' ' {
					dirY = -1
				} else if get(x, y+1) != ' ' {
					dirY = 1
				} else {
					panic("Illegal state in +")
				}
			}
		} else if c == ' ' {
			break
		}

		x += dirX
		y += dirY
		steps++
	}
	fmt.Println("Letters", string(letters))
	fmt.Println("steps", steps)
}
