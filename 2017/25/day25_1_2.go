package main

import "fmt"

const (
	stateA = iota
	stateB = iota
	stateC = iota
	stateD = iota
	stateE = iota
	stateF = iota
)

func run(newVal0, newVal1, move0, move1, ns0, ns1 int, tape map[int]int, cursor int) (int, int) {
	if tape[cursor] == 0 {
		tape[cursor] = newVal0
		return cursor + move0, ns0
	} else {
		tape[cursor] = newVal1
		return cursor + move1, ns1
	}
}

func a(tape map[int]int, cursor int) (int, int) {
	return run(1, 0, 1, -1, stateB, stateF, tape, cursor)
}
func b(tape map[int]int, cursor int) (int, int) {
	return run(0, 0, 1, 1, stateC, stateD, tape, cursor)
}
func c(tape map[int]int, cursor int) (int, int) {
	return run(1, 1, -1, 1, stateD, stateE, tape, cursor)
}
func d(tape map[int]int, cursor int) (int, int) {
	return run(0, 0, -1, -1, stateE, stateD, tape, cursor)
}
func e(tape map[int]int, cursor int) (int, int) {
	return run(0, 1, 1, 1, stateA, stateC, tape, cursor)
}
func f(tape map[int]int, cursor int) (int, int) {
	return run(1, 1, -1, 1, stateA, stateA, tape, cursor)
}

func main() {
	tape := make(map[int]int)
	cursor := 0
	ps := stateA

	for i := 0; i < 12794428; i++ {
		switch ps {
		case stateA:
			cursor, ps = a(tape, cursor)
			break
		case stateB:
			cursor, ps = b(tape, cursor)
			break
		case stateC:
			cursor, ps = c(tape, cursor)
			break
		case stateD:
			cursor, ps = d(tape, cursor)
			break
		case stateE:
			cursor, ps = e(tape, cursor)
			break
		case stateF:
			cursor, ps = f(tape, cursor)
			break
		}
	}

	checksum := 0
	for _, v := range tape {
		checksum += v
	}
	fmt.Println(checksum)
}
