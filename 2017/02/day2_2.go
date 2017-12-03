package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	sum := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		fs := strings.Fields(line)
		nums := make([]int, len(fs))

		for idx, num := range fs {
			i, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}

			nums[idx] = i
		}

		fmt.Println(nums)
		for i := 0; i < len(nums); i++ {
			for j := 0; j < len(nums); j++ {
				if i != j && nums[i]%nums[j] == 0 {
					sum += nums[i] / nums[j]
					break
				}
			}
		}
	}

	fmt.Println(sum)
}
