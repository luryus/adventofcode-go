package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	sum := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		min := math.MaxInt32
		max := math.MinInt32

		line_sc := bufio.NewScanner(strings.NewReader(line))
		line_sc.Split(bufio.ScanWords)
		for line_sc.Scan() {
			i, err := strconv.Atoi(line_sc.Text())
			if err != nil {
				panic(err)
			}
			if i < min {
				min = i
			}
			if i > max {
				max = i
			}
		}

		sum += max - min
	}

	fmt.Println(sum)
}
