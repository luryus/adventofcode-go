package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sndFreq int

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

type op interface {
	exec(regs map[byte]int) int
}

type sndOp struct {
	freq byte
}

func (op *sndOp) exec(regs map[byte]int) int {
	sndFreq = regs[op.freq]
	return 1
}

type setOp struct {
	dst byte
	val regOrImm
}

func (op *setOp) exec(regs map[byte]int) int {
	regs[op.dst] = op.val.value(regs)
	return 1
}

type addOp struct {
	dst byte
	val regOrImm
}

func (op *addOp) exec(regs map[byte]int) int {
	regs[op.dst] += op.val.value(regs)
	return 1
}

type mulOp struct {
	dst byte
	val regOrImm
}

func (op *mulOp) exec(regs map[byte]int) int {
	regs[op.dst] *= op.val.value(regs)
	return 1
}

type modOp struct {
	dst byte
	val regOrImm
}

func (op *modOp) exec(regs map[byte]int) int {
	regs[op.dst] %= op.val.value(regs)
	return 1
}

type rcvOp struct {
	cond byte
}

func (op *rcvOp) exec(regs map[byte]int) int {
	if regs[op.cond] != 0 {
		fmt.Println("rcv", sndFreq)
		os.Exit(0)
	}
	return 1
}

type jgzOp struct {
	cond   byte
	offset regOrImm
}

func (op *jgzOp) exec(regs map[byte]int) int {
	if regs[op.cond] > 0 {
		return op.offset.value(regs)
	}
	return 1
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	var ops []op

	for sc.Scan() {
		opStr := strings.Split(sc.Text(), " ")
		dstReg := opStr[1][0]
		opCode := opStr[0]

		var operand regOrImm
		if len(opStr) > 2 {
			imm, err := strconv.Atoi(opStr[2])
			if err != nil {
				operand = regVal(opStr[2][0])
			} else {
				operand = immVal(imm)
			}
		}

		switch opCode {
		case "set":
			ops = append(ops, &setOp{dst: dstReg, val: operand})
			break
		case "mul":
			ops = append(ops, &mulOp{dst: dstReg, val: operand})
			break
		case "add":
			ops = append(ops, &addOp{dst: dstReg, val: operand})
			break
		case "mod":
			ops = append(ops, &modOp{dst: dstReg, val: operand})
			break
		case "jgz":
			ops = append(ops, &jgzOp{cond: dstReg, offset: operand})
			break
		case "snd":
			ops = append(ops, &sndOp{freq: dstReg})
			break
		case "rcv":
			ops = append(ops, &rcvOp{cond: dstReg})
			break
		}
	}

	regs := make(map[byte]int)
	pc := 0
	maxPc := len(ops) - 1
	for {
		op := ops[pc]
		pc += op.exec(regs)

		if pc < 0 || pc > maxPc {
			break
		}
	}
}
