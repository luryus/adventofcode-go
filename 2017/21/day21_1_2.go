package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type bitmap struct {
	side   int
	pixels []byte
}

func newBitmap(side int, initial []byte) bitmap {
	if initial == nil {
		initial = make([]byte, side*side)
	}
	return bitmap{
		side:   side,
		pixels: initial,
	}
}

func (bm *bitmap) get(x, y int) byte {
	return bm.pixels[y*bm.side+x]
}

func (bm *bitmap) getSquare(x, y, side int) [][]byte {
	var sq [][]byte

	for i := 0; i < side; i++ {
		var row []byte
		for j := 0; j < side; j++ {
			row = append(row, bm.get(x+j, y+i))
		}
		sq = append(sq, row)
	}

	return sq
}

func (bm *bitmap) set(x, y int, val byte) {
	bm.pixels[y*bm.side+x] = val
}

func (bm *bitmap) setSquare(x, y int, vals [][]byte) {
	side := len(vals)
	for i := 0; i < side; i++ {
		for j := 0; j < side; j++ {
			bm.set(x+j, y+i, vals[i][j])
		}
	}
}

func main() {
	mappings := readInput()
	bm := newBitmap(3, []byte{'.', '#', '.', '.', '.', '#', '#', '#', '#'})
	iterations := 5
	if len(os.Args) > 1 && os.Args[1] == "2" {
		iterations = 18
	}

	for i := 0; i < iterations; i++ {
		if bm.side%2 == 0 {
			newBm := newBitmap(bm.side*3/2, nil)
			count := bm.side / 2

			for x := 0; x < count; x++ {
				for y := 0; y < count; y++ {
					oldSq := bm.getSquare(x*2, y*2, 2)
					key := string(oldSq[0]) + string(oldSq[1])
					newSq := mappings[key]
					newBm.setSquare(x*3, y*3, newSq)
				}
			}

			bm = newBm
		} else if bm.side%3 == 0 {
			newBm := newBitmap(bm.side*4/3, nil)
			count := bm.side / 3

			for x := 0; x < count; x++ {
				for y := 0; y < count; y++ {
					oldSq := bm.getSquare(x*3, y*3, 3)
					key := string(oldSq[0]) + string(oldSq[1]) + string(oldSq[2])
					newSq := mappings[key]
					newBm.setSquare(x*4, y*4, newSq)
				}
			}

			bm = newBm
		} else {
			panic("Illegal dimensions")
		}
	}

	// count
	count := 0
	for _, c := range bm.pixels {
		if c == '#' {
			count++
		}
	}
	fmt.Println(count)
}

func readInput() map[string][][]byte {
	mappings := make(map[string][][]byte)

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		parts := strings.Split(sc.Text(), " => ")

		var square [][]byte
		rows := strings.Split(parts[1], "/")
		for _, r := range rows {
			square = append(square, []byte(r))
		}

		keys := rotations(strings.Replace(parts[0], "/", "", -1))
		for _, k := range keys {
			mappings[k] = square
		}
	}

	return mappings
}

func rotations(sq string) []string {
	var rots []string = []string{sq}
	n := int(math.Sqrt(float64(len(sq))))

	rots = append(rots, flip(sq, n))

	for i := 0; i < 3; i++ {
		sq = rotate(sq, n)
		rots = append(rots, sq, flip(sq, n))
	}
	return rots
}

func rotate(sq string, n int) string {
	var out string
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			out += string(sq[(n-j-1)*n+i])
		}
	}
	return out
}

func flip(sq string, n int) string {
	row := func(i int) string {
		return sq[i*n : (i+1)*n]
	}

	flipped := row(n - 1)
	for r := n - 2; r >= 0; r-- {
		flipped += row(r)
	}
	return flipped
}
