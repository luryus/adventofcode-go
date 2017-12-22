package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	ps := readInput()

	part2 := len(os.Args) > 1

	const ticks int = 10000
	for i := 0; i < ticks; i++ {
		minDist, minID := 0x7fffffff, -1
		checkDist := i%100 == 0
		for _, p := range ps {
			p.advance()
			if checkDist {
				d := p.distance()
				if d < minDist {
					minID = p.id
					minDist = d
				}
			}
		}

		if part2 {
			var filtered []*particle
			type coord struct {
				x, y, z int
			}
			seen := make(map[coord]bool)
			for _, p := range ps {
				c := coord{x: p.x, y: p.y, z: p.z}
				if _, ok := seen[c]; !ok {
					seen[c] = true
					filtered = append(filtered, p)
				} else {
					for i, f := range filtered {
						fc := coord{x: f.x, y: f.y, z: f.z}
						if c == fc {
							filtered = append(filtered[:i], filtered[i+1:]...)
						}
					}
				}
			}
			ps = filtered
		}

		if checkDist {

			if minID == -1 {
				fmt.Println(ps)
				panic(-1)
			}
			fmt.Printf("\r%d, %d", minID, minDist)
		}
	}
	fmt.Println()
	fmt.Println("After filtering", len(ps), "remain")
}

type particle struct {
	id               int
	x, y, z          int
	velX, velY, velZ int
	accX, accY, accZ int
}

func (p *particle) advance() {
	p.velX += p.accX
	p.velY += p.accY
	p.velZ += p.accZ
	p.x += p.velX
	p.y += p.velY
	p.z += p.velZ
}

func (p *particle) distance() int {
	x, y, z := p.x, p.y, p.z
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	if z < 0 {
		z = -z
	}

	return x + y + z
}

func readInput() []*particle {
	const lineFormat string = `p=<%d,%d,%d>, v=<%d,%d,%d>, a=<%d,%d,%d>`

	var particles []*particle
	sc := bufio.NewScanner(os.Stdin)
	id := 0
	for sc.Scan() {
		var p particle
		p.id = id
		id++
		_, err := fmt.Sscanf(sc.Text(), lineFormat, &p.x, &p.y, &p.z, &p.velX, &p.velY, &p.velZ, &p.accX, &p.accY, &p.accZ)
		if err != nil {
			panic(err)
		}
		particles = append(particles, &p)
	}

	return particles
}
