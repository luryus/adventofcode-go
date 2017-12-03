package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	content, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	content = bytes.TrimSpace(content)

	sum := 0
	cont_len := len(content)
	half := cont_len / 2
	for i := 0; i < cont_len; i++ {
		if content[i] == content[(i+half)%cont_len] {
			log.Printf("Match: %d\n", content[i]-48)
			sum += int(content[i]) - 48
		}
	}

	log.Printf("Sum: %d\n", sum)
}
