package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var ops []int
	for scanner.Scan() {
		line := scanner.Text()
		val, err := strconv.Atoi(line)
		if err != nil {
			fmt.Printf("Int parsing failed for %s", line)
			panic(err)
		}
		ops = append(ops, val)
	}

	pc, iterations := 0, 0
	for pc < len(ops) {
		iterations++
		op := ops[pc]
		ops[pc]++
		pc += op
	}
	fmt.Printf("Iterations: %d\n", iterations)
}
