package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	commands := getCommands()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Welcome to the Pokedex!\n Usage:\n\n")
	commands["help"].callback()

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
		commandStr := cleanedInput[0]

		command, ok := commands[commandStr]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		if err := command.callback(); err != nil {
			fmt.Println(err)
			continue
		}
	}
}
