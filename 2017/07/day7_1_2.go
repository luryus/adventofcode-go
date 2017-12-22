package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type node struct {
	name       string
	weight     int
	fullWeight int
	children   []*node
	parent     *node
}

var nodeRegexp *regexp.Regexp = regexp.MustCompile(`^(\w+) \((\d+)\)`)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	graph := make(map[string]*node)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " -> ")

		m := nodeRegexp.FindAllStringSubmatch(parts[0], -1)

		name := m[0][1]
		weight, err := strconv.Atoi(m[0][2])
		if err != nil {
			panic(err)
		}

		n, found := graph[name]
		if !found {
			graph[name] = &node{name: name}
			n = graph[name]
		}
		n.weight = weight

		if len(parts) > 1 {
			m := strings.Split(strings.TrimSpace(parts[1]), ", ")
			for _, cname := range m {
				c, cf := graph[cname]
				if !cf {
					graph[cname] = &node{name: cname}
					c = graph[cname]
				}
				c.parent = n
				n.children = append(n.children, c)
			}
		}
	}

	var root *node

	for _, v := range graph {
		if v.parent == nil {
			root = v
			fmt.Println("Root:", root.name)
			break
		}
	}

	calculateWeights(root)

	unbc := root

	for {
		unbc = unbalancedChild(unbc)
		if isBalanced(unbc) {
			break
		}
	}

	rp := unbc.parent
	diff := 0
	for _, c := range rp.children {
		if c.fullWeight != unbc.fullWeight {
			diff = (unbc.fullWeight - c.fullWeight)
			break
		}
	}
	fmt.Println(diff)
	fmt.Println(unbc.weight - diff)
}

func isBalanced(n *node) bool {
	if len(n.children) <= 1 {
		return true
	}
	fw := n.children[0].fullWeight
	for _, c := range n.children {
		if c.fullWeight != fw {
			return false
		}
	}
	return true
}

func unbalancedChild(n *node) *node {
	fw := n.children[0].fullWeight
	if fw != n.children[1].fullWeight && fw != n.children[2].fullWeight {
		return n.children[0]
	}

	for _, c := range n.children[1:] {
		if c.fullWeight != fw {
			return c
		}
	}

	return nil
}

func calculateWeights(root *node) int {
	cw := 0
	for _, c := range root.children {
		cw += calculateWeights(c)
	}
	root.fullWeight = cw + root.weight
	return root.fullWeight
}
