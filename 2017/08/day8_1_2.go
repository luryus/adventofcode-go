package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var lineRegexp *regexp.Regexp = regexp.MustCompile(`^(\w+) (inc|dec) (-?\d+) if (\w+) ([<>=!]{1,2}) (-?\d+)$`)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	regs := make(map[string]int)
	totalMax := 0

	for scanner.Scan() {
		ins := scanner.Text()
		instr := lineRegexp.FindAllStringSubmatch(ins, -1)[0]
		cmpOp := instr[5]
		cmp1 := regs[instr[4]]
		cmp2, _ := strconv.Atoi(instr[6])
		if evalComparison(cmpOp, cmp1, cmp2) {
			reg := instr[1]
			mul := 1
			if instr[2] == "dec" {
				mul = -1
			}
			val, _ := strconv.Atoi(instr[3])
			regs[reg] += mul * val

			if regs[reg] > totalMax {
				totalMax = regs[reg]
			}
		}
	}

	maxVal := -0x80000000
	fmt.Println(maxVal)
	for _, v := range regs {
		if v > maxVal {
			maxVal = v
		}
	}

	fmt.Println(maxVal)
	fmt.Println(totalMax)
}

func evalComparison(cmpOp string, cmp1, cmp2 int) bool {
	switch cmpOp {
	case "==":
		return cmp1 == cmp2
	case ">":
		return cmp1 > cmp2
	case "<":
		return cmp1 < cmp2
	case "<=":
		return cmp1 <= cmp2
	case ">=":
		return cmp1 >= cmp2
	case "!=":
		return cmp1 != cmp2
	}

	panic(cmpOp)
}
