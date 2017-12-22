package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fws := make(map[int]int)

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		parts := strings.Split(sc.Text(), ": ")
		level, _ := strconv.Atoi(parts[0])
		depth, _ := strconv.Atoi(parts[1])

		fws[level] = depth
	}

	wait := 1
delayLoop:
	for {
		for l, fw := range fws {
			if (wait+l)%((fw-1)*2) == 0 {
				wait++
				continue delayLoop
			}
		}
		break
	}
	fmt.Println("Min delay", wait)

	severity := 0
	for l, fw := range fws {
		if l%(fw*2-2) == 0 {
			severity += l * fw
		}
	}

	fmt.Println("Severity", severity)
}
