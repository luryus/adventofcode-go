package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const part2rounds = 1000000000

type op interface {
	run(buf []string) []string
}

type spinOp struct {
	n int
}

type swapiOp struct {
	a int
	b int
}

type swapOp struct {
	a string
	b string
}

func find(c string, buf []string) int {
	for i, v := range buf {
		if v == c {
			return i
		}
	}

	return -1
}

func (op spinOp) run(buf []string) []string {
	return append(buf[len(buf)-op.n:], buf[:len(buf)-op.n]...)
}

func (op swapiOp) run(buf []string) []string {
	buf[op.a], buf[op.b] = buf[op.b], buf[op.a]
	return buf
}

func (op swapOp) run(buf []string) []string {
	a_i, b_i := find(op.a, buf), find(op.b, buf)
	buf[a_i], buf[b_i] = buf[b_i], buf[a_i]
	return buf
}

func main() {
	re := bufio.NewReader(os.Stdin)

	buf := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p"}

	var ops []op

	for {
		tok, err := re.ReadString(',')
		if err != nil && err != io.EOF {
			break
		}
		tok = strings.TrimRight(tok, ", \n")

		switch tok[0] {
		case 's':
			n, _ := strconv.Atoi(tok[1:])
			op := spinOp{n: n}
			ops = append(ops, op)
			buf = op.run(buf)
			break
		case 'x':
			ps := strings.Split(tok[1:], "/")
			a, _ := strconv.Atoi(ps[0])
			b, _ := strconv.Atoi(ps[1])
			op := swapiOp{a: a, b: b}
			buf = op.run(buf)
			ops = append(ops, op)
			break
		case 'p':
			op := swapOp{a: string(tok[1]), b: string(tok[3])}
			buf = op.run(buf)
			ops = append(ops, op)
			break
		}

		if err == io.EOF {
			break
		}
	}

	res := strings.Join(buf, "")
	fmt.Println(res)

	roundRes := make(map[string]int)
	roundRes[res] = 0

	for i := 1; i < part2rounds; i++ {
		if i%100 == 0 {
			fmt.Printf("\r%d", i)
		}
		for _, op := range ops {
			buf = op.run(buf)
		}

		res := strings.Join(buf, "")
		r, found := roundRes[res]
		if found {
			diff := i - r
			rem := part2rounds - i

			i += (rem / diff) * diff
			roundRes = make(map[string]int)
		} else {
			roundRes[res] = i
		}
	}

	fmt.Println(strings.Join(buf, ""))
}
