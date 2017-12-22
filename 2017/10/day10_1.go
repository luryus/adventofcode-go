package main

import (
	"bufio"
	"container/ring"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const ringLen int = 256

func main() {
	rStart := ring.New(ringLen)
	e := rStart
	for i := 0; i < ringLen; i++ {
		e.Value = i
		e = e.Next()
	}

	skip := 0
	reader := bufio.NewReader(os.Stdin)
	inStr, _ := reader.ReadString(0)
	strLens := strings.Split(strings.TrimSpace(inStr), ",")
	var readLens []int
	for _, v := range strLens {
		length, _ := strconv.Atoi(v)
		readLens = append(readLens, length)
	}

	r := rStart
	for _, readLen := range readLens {
		s, e := r, r.Move(readLen-1)
		for s != e && (e.Next() != s || readLen == ringLen) {
			tmp := s.Value
			s.Value = e.Value
			e.Value = tmp

			s = s.Next()
			e = e.Prev()
		}
		r = r.Move(readLen + skip)
		skip++
	}

	fmt.Println(rStart.Value.(int) * rStart.Next().Value.(int))
}
