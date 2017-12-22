package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	clean    = 0
	weakened = -1
	infected = -2
	flagged  = 1
)
const leftTurn complex128 = complex(0, 1)
const rightTurn complex128 = complex(0, -1)

func part1() {
	virusMap, initSize := readInput()
	pos := complex(float64(initSize/2), float64(-initSize/2))
	dir := complex(0, 1)

	infections := 0

	for i := 0; i < 10000; i++ {
		status := virusMap[pos]
		if status == infected {
			dir *= rightTurn
			virusMap[pos] = clean
		} else {
			infections++
			dir *= leftTurn
			virusMap[pos] = infected
		}
		pos += dir
	}
	fmt.Println(infections)
}

func part2() {
	virusMap, initSize := readInput()
	pos := complex(float64(initSize/2), float64(-initSize/2))
	dir := complex(0, 1)

	infections := 0

	for i := 0; i < 10000000; i++ {
		status := virusMap[pos]
		if status == clean {
			virusMap[pos] = weakened
			dir *= leftTurn
		} else if status == weakened {
			virusMap[pos] = infected
			infections++
		} else if status == infected {
			virusMap[pos] = flagged
			dir *= rightTurn
		} else {
			virusMap[pos] = clean
			dir *= -1
		}
		pos += dir
	}
	fmt.Println(infections)
}

func main() {
	if os.Args[1] == "2" {
		part2()
	} else {
		part1()
	}
}

func readInput() (map[complex128]int, int) {
	sc := bufio.NewScanner(os.Stdin)
	line := 0
	rowLen := -1
	virusMap := make(map[complex128]int)

	for sc.Scan() {
		if rowLen == -1 {
			rowLen = len(sc.Text())
		}
		for i, node := range sc.Text() {
			coord := complex(float64(i), float64(line))
			val := 0
			if node == '#' {
				val = infected
			} else if node == '.' {
				val = clean
			} else {
				panic("Invalid input")
			}
			virusMap[coord] = val
		}
		line--
	}

	return virusMap, rowLen
}
