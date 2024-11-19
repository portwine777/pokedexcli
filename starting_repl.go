package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliComand struct {
	name        string
	description string
	callback    func() error
}

func getComands() map[string]cliComand {
	return map[string]cliComand{
		"help": {
			name:        "help",
			description: "Displays a have message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func cleanInput(text string) []string { // Добовляем слова ввода в срез!
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

func startREPL() {
	reading := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("pokedex >")
		reading.Scan()

		words := cleanInput(reading.Text())
		if len(words) == 0 {
			continue
		}
		commandName := words[0]

		command, exists := getComands()[commandName]
		if exists {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}

	}
}
