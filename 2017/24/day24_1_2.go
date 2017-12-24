package main

import (
	"bufio"
	"fmt"
	"os"
)

type component struct {
	start int
	end   int
}

func readInput() []*component {
	sc := bufio.NewScanner(os.Stdin)
	var cs []*component
	for sc.Scan() {
		a, b := 0, 0
		fmt.Sscanf(sc.Text(), "%d/%d", &a, &b)
		cs = append(cs, &component{start: a, end: b})
	}

	return cs
}

func maxStrength(startWidth int, comps []*component) int {
	max := 0
	for i, c := range comps {
		var endWidth int
		if c.start == startWidth {
			endWidth = c.end
		} else if c.end == startWidth {
			endWidth = c.start
		} else {
			continue
		}

		copyComps := append([]*component(nil), comps...)
		remaining := append(copyComps[:i], copyComps[i+1:]...)
		w := maxStrength(endWidth, remaining)
		s := c.start + c.end + w
		if s > max {
			max = s
		}
	}
	return max
}

func maxLength(startWidth int, comps []*component) (int, int) {
	maxLen, strength := 0, 0
	for i, c := range comps {
		var endWidth int
		if c.start == startWidth {
			endWidth = c.end
		} else if c.end == startWidth {
			endWidth = c.start
		} else {
			continue
		}

		copyComps := append([]*component(nil), comps...)
		remaining := append(copyComps[:i], copyComps[i+1:]...)
		l, s := maxLength(endWidth, remaining)
		s += c.start + c.end
		if l+1 > maxLen || (l+1 == maxLen && s > strength) {
			maxLen = l + 1
			strength = s
		}
	}
	return maxLen, strength
}

func main() {
	comps := readInput()

	if len(os.Args) > 1 {
		maxL, s := maxLength(0, comps)
		fmt.Println(maxL, s)
	} else {
		maxS := maxStrength(0, comps)
		fmt.Println(maxS)
	}
}
