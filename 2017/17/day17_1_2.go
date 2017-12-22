package main

import (
	"container/ring"
	"fmt"
	"os"
	"strconv"
	"sync"
)

func printRing(r *ring.Ring) {
	start := r
	for {
		fmt.Printf("%d,", r.Value.(int))
		r = r.Next()
		if r == start {
			fmt.Println()
			return
		}
	}
}

func main() {
	input, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go run1(input, &wg)
	go run2(input, &wg)

	wg.Wait()
}

func run2(input int, wg *sync.WaitGroup) {
	const rounds int = 50000000
	pos, next := 0, 0
	for i := 1; i < rounds; i++ {
		pos = (pos + input) % i
		if pos == 0 {
			next = i
		}
		pos = (pos + 1)
	}

	fmt.Println("After zero", next)
	wg.Done()
}

func run1(input int, wg *sync.WaitGroup) {
	const rounds int = 2018
	r := ring.New(1)
	r.Value = 0

	for i := 1; i < rounds; i++ {
		if i%100000 == 0 {
			fmt.Println(i)
		}
		r = r.Move(input)
		new := ring.New(1)
		new.Value = i

		r.Link(new)
		r = r.Next()
	}

	r = r.Next()
	fmt.Println("After r", r.Value)

	wg.Done()
}
