package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanRunes)

	scanner.Scan()

	if scanner.Text() != "{" {
		panic("Invalid input start")
	}

	score, glen := readGroup(scanner, 1)
	fmt.Println("Score", score)
	fmt.Println("Garbage", glen)
}

func readGroup(sc *bufio.Scanner, level int) (int, int) {
	// current position should have '{'
	score := level
	garblen := 0
	for sc.Scan() {
		next := sc.Text()
		if next == "}" {
			return score, garblen
		} else if next == "{" {
			csc, glen := readGroup(sc, level+1)
			score += csc
			garblen += glen
		} else if next == "<" {
			garblen += readGarbage(sc)
		} else if next == "!" {
			readCancel(sc)
		}
	}
	return score, garblen
}

func readGarbage(sc *bufio.Scanner) int {
	// the current text is <
	// read until garbage ends
	glen := 0
	for sc.Scan() {
		next := sc.Text()
		if next == ">" {
			return glen
		} else if next == "!" {
			readCancel(sc)
		} else {
			glen++
		}
	}
	return glen
}

func readCancel(sc *bufio.Scanner) {
	// the current position has !
	// read another character
	sc.Scan()
}
