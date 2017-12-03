package main

import (
	"crypto/md5"
	"encoding/hex"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		println("Invalid arguments. Give input word.")
		os.Exit(1)
	}

	inputWord := os.Args[1]
	var counter uint64 = 0
	password := ""

	for len(password) < 8 {
		data := []byte(inputWord + strconv.FormatUint(counter, 10))
		sum := md5.Sum(data)
		hexSum := hex.EncodeToString(sum[:])

		if strings.HasPrefix(hexSum, "00000") {
			password += string(hexSum[5])
			println("Found hash at", counter, ":", hexSum)
		}

		counter += 1
	}

	println("Password is", password)
}
