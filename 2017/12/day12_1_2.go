package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	visited bool
	conns   []int
}

func main() {
	nodes := readInput()
	fmt.Println(countNodes(nodes[0], nodes))

	count := 1
	start := prune(nodes)
	for start != nil {
		count++
		countNodes(start, nodes)
		start = prune(nodes)
	}
	fmt.Println(count)
}

func prune(graph map[int]*node) *node {
	var start *node
	for k, v := range graph {
		if v.visited {
			delete(graph, k)
		} else if start == nil {
			start = v
		}
	}

	return start
}

func countNodes(start *node, graph map[int]*node) int {
	start.visited = true
	count := 1
	for _, c := range start.conns {
		if !graph[c].visited {
			count += countNodes(graph[c], graph)
		}
	}

	return count
}

func readInput() map[int]*node {
	sc := bufio.NewScanner(os.Stdin)
	nodes := make(map[int]*node)
	for sc.Scan() {
		parts := strings.Split(sc.Text(), " <-> ")
		nid, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}

		var conns []int
		for _, cstr := range strings.Split(parts[1], ", ") {
			c, err := strconv.Atoi(cstr)
			if err != nil {
				panic(err)
			}

			conns = append(conns, c)
		}

		nodes[nid] = &node{visited: false, conns: conns}
	}

	return nodes
}
