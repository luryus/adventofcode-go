package main

import (
	"container/ring"
	"fmt"
	"os"
)

type bitmap struct {
	width  uint
	height uint

	bits [][]byte
}

func (bm *bitmap) get(x, y int) byte {
	return bm.bits[y][x]
}

func (bm *bitmap) print() {
	for _, r := range bm.bits {
		for _, c := range r {
			if c == 0 {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}

func (bm *bitmap) set(x, y int, val byte) {
	bm.bits[y][x] = val
}

func main() {
	input := os.Args[1]
	var hashes [][]byte
	s := 0
	for i := 0; i < 128; i++ {
		h := hash(fmt.Sprintf("%s-%d", input, i))
		s += sum(h)
		hashes = append(hashes, h)
	}

	fmt.Println(s)

	areas := 0
	bm := toBitmap(hashes)
	for y, x := findEnabled(bm, 0, 0); y != -1; y, x = findEnabled(bm, y, x) {
		areas++
		clearArea(bm, x, y)
	}

	fmt.Println(areas)
}

func toBitmap(hashes [][]byte) *bitmap {

	var rows [][]byte

	for _, hr := range hashes {
		var row []byte
		for _, b := range hr {
			row = append(row,
				byte(b>>5&0x1),
				byte(b>>6&0x1),
				byte(b>>5&0x1),
				byte(b>>4&0x1),
				byte(b>>3&0x1),
				byte(b>>2&0x1),
				byte(b>>1&0x1),
				byte(b&0x1))
		}

		rows = append(rows, row)
	}

	return &bitmap{width: 128, height: 128, bits: rows}

}

func findEnabled(bm *bitmap, start_y, start_x int) (int, int) {
	for i := start_y; i < 128; i++ {
		j := 0
		if start_x > 0 {
			j = start_x
			start_x = -1
		}
		for ; j < 128; j++ {
			if bm.get(j, i) != 0 {
				return i, j
			}
		}
	}

	return -1, -1
}

func clearArea(bm *bitmap, x, y int) {

	if bm.get(x, y) != 0 {
		bm.set(x, y, 0)
		if y-1 >= 0 && bm.get(x, y-1) != 0 {
			clearArea(bm, x, y-1)
		}
		if y+1 < 128 && bm.get(x, y+1) != 0 {
			clearArea(bm, x, y+1)
		}

		if x-1 >= 0 && bm.get(x-1, y) != 0 {
			clearArea(bm, x-1, y)
		}
		if x+1 < 128 && bm.get(x+1, y) != 0 {
			clearArea(bm, x+1, y)
		}
	}
}

func sum(hash []byte) int {
	sum := byte(0)
	for _, v := range hash {
		h := v
		for i := 0; i < 128; i++ {
			sum += h & 0x1
			h = h >> 1
		}
	}
	return int(sum)
}

func hash(input string) []byte {
	const ringLen int = 256
	const rounds int = 64

	var inputSuffix []uint8 = []uint8{17, 31, 73, 47, 23}

	rStart := ring.New(ringLen)
	e := rStart
	for i := 0; i < ringLen; i++ {
		e.Value = uint8(i)
		e = e.Next()
	}

	inputLengths := inputToLens(input)
	inputLengths = append(inputLengths, inputSuffix...)

	skip := 0
	r := rStart
	for i := 0; i < rounds; i++ {
		for _, readLen := range inputLengths {
			l := int(readLen)
			s, e := r, r.Move(l-1)
			reverseValues(s, e, l, ringLen)
			r = r.Move(l + skip)
			skip++
		}
	}

	return buildDenseHash(rStart)

}

func inputToLens(input string) []uint8 {
	var lens []uint8
	for _, c := range input {
		lens = append(lens, uint8(c))
	}
	return lens
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

func reverseValues(s, e *ring.Ring, readLen, ringLen int) {
	for s != e && (e.Next() != s || readLen == ringLen) {
		s.Value, e.Value = e.Value, s.Value

		s = s.Next()
		e = e.Prev()
	}
}
