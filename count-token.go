package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/tiktoken-go/tokenizer"
)

func main() {
	enc, err := tokenizer.Get(tokenizer.O200kBase)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error getting tokenizer:", err)
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)
	totalTokens := 0

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Fprintln(os.Stderr, "Error reading from stdin:", err)
			os.Exit(1)
		}

		ids, _, _ := enc.Encode(line)
		totalTokens += len(ids)
	}

	fmt.Println(totalTokens)
}
