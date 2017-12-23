package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

type regOrImm interface {
	value(regs map[byte]int) int
}
type immVal int
type regVal byte

func (imm immVal) value(regs map[byte]int) int {
	return int(imm)
}
func (reg regVal) value(regs map[byte]int) int {
	val, ok := regs[byte(reg)]
	if !ok {
		return 0
	}
	return val
}

const (
	opSet = iota
	opSub = iota
	opMul = iota
	opJnz = iota
)

type operation struct {
	reg regOrImm
	val regOrImm
	typ int
}

func readInput() []*operation {
	var ops []*operation
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		parts := strings.Split(sc.Text(), " ")
		opcode := 0
		switch parts[0] {
		case "sub":
			opcode = opSub
			break
		case "mul":
			opcode = opMul
			break
		case "jnz":
			opcode = opJnz
			break
		case "set":
			opcode = opSet
			break
		}
		op := operation{typ: opcode}

		imm, err := strconv.Atoi(parts[1])
		if err != nil {
			op.reg = regVal(parts[1][0])
		} else {
			op.reg = immVal(imm)
		}
		imm, err = strconv.Atoi(parts[2])
		if err != nil {
			op.val = regVal(parts[2][0])
		} else {
			op.val = immVal(imm)
		}

		ops = append(ops, &op)
	}
	return ops
}

func main() {

	ops := readInput()
	regs := make(map[byte]int)
	pc, cycles := 0, 0
	mulCount := 0

	if len(os.Args) > 1 {
		regs['a'] = 1
	}

	for pc >= 0 && pc < len(ops) {
		if pc == 8 {
			b := regs['b']
			c := regs['c']
			for {
				d := 2
				mulCount += (b - d) * (b - 2)
				if !big.NewInt(int64(b)).ProbablyPrime(0) {
					regs['h'] += 1
				}
				if c == b {
					break
				} else {
					b += 17
				}
			}
			pc = 32
			continue
		}
		cycles++
		op := ops[pc]
		switch op.typ {
		case opSet:
			regs[byte(op.reg.(regVal))] = op.val.value(regs)
			pc++
			break
		case opSub:
			regs[byte(op.reg.(regVal))] -= op.val.value(regs)
			pc++
			break
		case opMul:
			mulCount++
			regs[byte(op.reg.(regVal))] *= op.val.value(regs)
			pc++
			break
		case opJnz:
			offset := op.val.value(regs)
			if op.reg.value(regs) != 0 {
				pc += offset
			} else {
				pc++
			}
			break
		}
	}
	fmt.Println("h", regs['h'])
	fmt.Println("mul", mulCount)
}
