package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Pokedex > ")

		var input string
		for scanner.Scan() {
			input = scanner.Text()
			break
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
			return
		}

		cleanedInput := CleanInput(input)
		fmt.Printf("Your command was: %s\n", cleanedInput[0])

	}
}
