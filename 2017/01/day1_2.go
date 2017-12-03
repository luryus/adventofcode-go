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
	for i := 0; i < len(content)-1; i++ {
		if content[i] == content[i+1] {
			log.Printf("Match: %d\n", content[i]-48)
			sum += int(content[i]) - 48
		}
	}

	if content[0] == content[len(content)-1] {
		sum += int(content[0]) - 48
	}

	log.Printf("Sum: %d\n", sum)
	log.Printf("Last: %U", content[len(content)-1])
}
