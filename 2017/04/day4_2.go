package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func sortStringChars(s string) string {
	c := strings.Split(s, "")
	sort.Strings(c)
	return strings.Join(c, "")
}

func main() {
	valid, invalid := 0, 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		words := strings.Fields(line)
		for i, s := range words {
			words[i] = sortStringChars(s)
		}

		v := true
	outer:
		for i, a := range words {
			for j, b := range words {
				if i != j && a == b {
					v = false
					break outer
				}
			}
		}
		if v {
			valid++
		} else {
			invalid++
		}
	}

	fmt.Printf("Valid: %d, Invalid: %d\n", valid, invalid)
}
