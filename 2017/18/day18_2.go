package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

type regOrImm interface {
	value(regs map[byte]int) int
}

const (
	opSnd = iota
	opRcv = iota
	opAdd = iota
	opSet = iota
	opMul = iota
	opMod = iota
	opJgz = iota
)

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

type op struct {
	reg regOrImm
	val regOrImm
	typ int
}

func run(id int, ops []op, chin <-chan int, chout chan<- int, count *int, finMutex *sync.RWMutex, fin *bool) {
	regs := make(map[byte]int)
	regs['p'] = id
	pc := 0
	maxPc := len(ops) - 1
	*count = 0
	for {
		op := ops[pc]
		typ := op.typ

		switch typ {
		case opSnd:
			chout <- op.reg.value(regs)
			(*count)++
			if id == 1 {
				fmt.Println(*count)
			}
			pc++
			break
		case opRcv:
		loop:
			for {
				select {
				case regs[byte(op.reg.(regVal))] = <-chin:
					finMutex.Lock()
					*fin = false
					finMutex.Unlock()
					break loop
				default:
					finMutex.Lock()
					*fin = true
					finMutex.Unlock()
				}
			}
			//regs[byte(op.reg.(regVal))] = <-chin
			pc++
			break
		case opAdd:
			regs[byte(op.reg.(regVal))] += op.val.value(regs)
			pc++
			break
		case opSet:
			regs[byte(op.reg.(regVal))] = op.val.value(regs)
			pc++
			break
		case opMul:
			regs[byte(op.reg.(regVal))] *= op.val.value(regs)
			pc++
			break
		case opMod:
			regs[byte(op.reg.(regVal))] %= op.val.value(regs)
			pc++
			break
		case opJgz:
			if op.reg.value(regs) > 0 {
				pc += op.val.value(regs)
			} else {
				pc++
			}
			break
		}

		if pc < 0 || pc > maxPc {
			break
		}
	}
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	var ops []op

	for sc.Scan() {
		opStr := strings.Split(sc.Text(), " ")
		opCode := opStr[0]

		var op1, op2 regOrImm
		imm, err := strconv.Atoi(opStr[1])
		if err != nil {
			op1 = regVal(opStr[1][0])
		} else {
			op1 = immVal(imm)
		}

		if len(opStr) > 2 {
			imm, err := strconv.Atoi(opStr[2])
			if err != nil {
				op2 = regVal(opStr[2][0])
			} else {
				op2 = immVal(imm)
			}
		}

		switch opCode {
		case "set":
			ops = append(ops, op{typ: opSet, reg: op1, val: op2})
			break
		case "mul":
			ops = append(ops, op{typ: opMul, reg: op1, val: op2})
			break
		case "add":
			ops = append(ops, op{typ: opAdd, reg: op1, val: op2})
			break
		case "mod":
			ops = append(ops, op{typ: opMod, reg: op1, val: op2})
			break
		case "jgz":
			ops = append(ops, op{typ: opJgz, reg: op1, val: op2})
			break
		case "snd":
			ops = append(ops, op{typ: opSnd, reg: op1})
			break
		case "rcv":
			ops = append(ops, op{typ: opRcv, reg: op1})
			break
		}
	}

	ch0to1, ch1to0 := make(chan int, 1000000000), make(chan int, 1000000000)

	var rwm1, rwm2 sync.RWMutex
	var fin1, fin2 bool
	var c1, c2 int

	go run(0, ops, ch1to0, ch0to1, &c1, &rwm1, &fin1)
	go run(1, ops, ch0to1, ch1to0, &c2, &rwm2, &fin2)

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
