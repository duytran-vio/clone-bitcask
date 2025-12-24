package main

import (
	"bufio"
	Bitcaskdb "clone-bitcask/bitcaskDB"
	BitcaskFile "clone-bitcask/bitcaskDB/bitcaskFile"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Bitcask CLI - Type 'QUIT' to exit")

	fileFactory := func(fileID int, filePath string) Bitcaskdb.StorageFile {
		return BitcaskFile.NewBitcaskFile(fileID, filePath)
	}
	bitcaskdb := Bitcaskdb.NewBitcaskDB(fileFactory)
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
