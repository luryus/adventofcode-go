package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func max(s []int) (int, int) {
	m, mi := s[0], 0
	for i, e := range s {
		if e > m {
			mi = i
			m = e
		}
	}

	return mi, m
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	var mem []int

	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		mem = append(mem, i)
	}

	seen := make(map[string]bool)

	cycles := 0
	seccycle := false
	for {
		cycles++
		maxi, maxv := max(mem)
		mem[maxi] = 0
		for maxv > 0 {
			maxi = (maxi + 1) % len(mem)
			mem[maxi]++
			maxv--
		}

		// store the state
		var strmem []string
		for _, blks := range mem {
			strmem = append(strmem, strconv.Itoa(blks))
		}
		state := strings.Join(strmem, ",")

		if seen[state] && !seccycle {
			seen = make(map[string]bool)
			cycles = 0
			seccycle = true
			seen[state] = true
		} else if seen[state] {
			break
		} else {
			seen[state] = true
		}
	}

	fmt.Printf("%d\n", len(seen)+1)
	fmt.Printf("%d\n", cycles)
}
