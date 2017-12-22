package main

import (
	"fmt"
	"os"
	"strconv"
)

const a_mul uint64 = 16807
const b_mul uint64 = 48271
const rounds1 int = 40000000
const rounds2 int = 5000000

const mod uint64 = 0x7fffffff

func main() {
	a, _ := strconv.ParseUint(os.Args[1], 10, 64)
	b, _ := strconv.ParseUint(os.Args[2], 10, 64)

	if len(os.Args) > 3 {
		fmt.Println(part2(a, b))
	} else {
		fmt.Println(part1(a, b))
	}

}

func part1(a, b uint64) int {
	matches := 0
	for i := 0; i < rounds1; i++ {
		a = (a_mul * a) % mod
		b = (b_mul * b) % mod

		if (a & 0xffff) == (b & 0xffff) {
			matches++
		}
	}
	return matches
}

func part2(a, b uint64) int {
	matches := 0
	for i := 0; i < rounds2; i++ {
		for {
			a = (a_mul * a) % mod
			if (a & 0x3) == 0 {
				break
			}
		}
		for {
			b = (b_mul * b) % mod
			if (b & 0x7) == 0 {
				break
			}
		}

		if (a & 0xffff) == (b & 0xffff) {
			matches++
		}
	}

	return matches
}
