package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type config struct {
	Next     int
	Previous int
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    CommandExit,
		},
		"help": {
			name:        "help",
			description: "Display the help message",
			callback:    HelpCallback,
		},
		"map": {
			name:        "map",
			description: "Display 20 location area. Each subsequent call displays next 20",
			callback:    MapCallBack,
		},
		"mapb": {
			name:        "mapb",
			description: "Display previous page",
			callback:    MapBCallBack,
		},
	}
}

func CleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func CommandExit(config *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func HelpCallback(config *config) error {
	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}

func MapCallBack(config *config) error {
	locationID := config.Next
	for locationID <= config.Next+19 {
		res, err := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%d/", locationID))
		if err != nil {
			return err
		}

		defer res.Body.Close()

		loc := location{}
		if err := json.NewDecoder(res.Body).Decode(&loc); err != nil {
			return err
		}

		fmt.Println(loc.Name)
		locationID++
	}
	config.Next = locationID
	config.Previous = locationID - 40
	return nil
}

func MapBCallBack(config *config) error {
	locationID := config.Previous
	if locationID < 1 {
		return errors.New("can not fetch location before id 1")
	}
	for locationID <= config.Previous+19 {
		res, err := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%d/", locationID))
		if err != nil {
			return err
		}

		defer res.Body.Close()

		loc := location{}
		if err := json.NewDecoder(res.Body).Decode(&loc); err != nil {
			return err
		}

		fmt.Println(loc.Name)
		locationID++
	}
	config.Next = locationID
	config.Previous = locationID - 40
	return nil
}
