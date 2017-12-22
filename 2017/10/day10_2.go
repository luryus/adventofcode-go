package main

import (
	"bufio"
	"container/ring"
	"fmt"
	"os"
	"strings"
)

const ringLen int = 256
const rounds int = 64

var inputSuffix []uint8 = []uint8{17, 31, 73, 47, 23}

func main() {

	rStart := ring.New(ringLen)
	e := rStart
	for i := 0; i < ringLen; i++ {
		e.Value = uint8(i)
		e = e.Next()
	}

	inputLengths := readInput()

	skip := 0
	r := rStart
	for i := 0; i < rounds; i++ {
		for _, readLen := range inputLengths {
			l := int(readLen)
			s, e := r, r.Move(l-1)
			reverseValues(s, e, l)
			r = r.Move(l + skip)
			skip++
		}
	}

	denseHash := buildDenseHash(rStart)

	fmt.Printf("%x\n", denseHash)
}

func buildDenseHash(rs *ring.Ring) []byte {
	var denseHash []byte
	for i := 0; i < 16; i++ {
		r := rs.Move(i * 16)
		h := r.Value.(uint8)
		for j := 1; j < 16; j++ {
			r = r.Next()
			h = h ^ r.Value.(uint8)
		}
		denseHash = append(denseHash, h)
	}
	return denseHash
}

func reverseValues(s, e *ring.Ring, readLen int) {
	for s != e && (e.Next() != s || readLen == ringLen) {
		s.Value, e.Value = e.Value, s.Value

		s = s.Next()
		e = e.Prev()
	}
}

func readInput() []uint8 {
	reader := bufio.NewReader(os.Stdin)
	inStr, _ := reader.ReadString(0)
	inStr = strings.TrimSpace(inStr)
	var readLens []uint8
	for _, v := range inStr {
		readLens = append(readLens, uint8(v))
	}

	readLens = append(readLens, inputSuffix...)

	return readLens
}
