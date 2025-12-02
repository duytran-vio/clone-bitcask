package main

import (
	"bufio"
	Bitcaskdb "clone-bitcask/bitcaskDB"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Bitcask CLI - Type 'QUIT' to exit")

	bitcaskdb := Bitcaskdb.NewBitcaskDB()
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break // EOF or error
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "QUIT" {
			fmt.Println("End Bitcask CLI!")
			break
		}

		// Handle commands
		output := bitcaskdb.HandleCommand(input)
		fmt.Println(output)
	}
}
